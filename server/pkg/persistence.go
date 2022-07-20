package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type AuthSource string

const (
	GoogleAuthSource AuthSource = "google"
)

func NewDatabase() (*Database, error) {
	//TODO remove this secret
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to db")
	}
	return &Database{
		Pool: pool,
	}, nil
}

type Database struct {
	Pool *pgxpool.Pool
}

func (d Database) ReadUsers() ([]User, error) {
	rows, err := d.Pool.Query(context.Background(), "select * from users")
	if err != nil {
		logrus.WithError(err).Error("Query failed")
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var id string
		var name string
		var email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}

		users = append(users, User{id, name, email})
	}

	return users, rows.Err()
}

func (d Database) ReadUser(id string) (*User, error) {
	var name string
	var email string
	err := d.Pool.QueryRow(context.Background(), "select name, email from users where id = $1", id).Scan(&name, &email)
	if err != nil {
		if noRowsFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to read user")
	}

	return &User{
		Id:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (d Database) ReadUsersForEvent(eventId string) ([]User, error) {
	rows, err := d.Pool.Query(context.Background(), "SELECT users.id, users.name, users.email FROM users INNER JOIN memberships ON users.id = memberships.user_id AND memberships.event_id = $1", eventId)
	if err != nil {
		logrus.WithError(err).Error("Query failed")
		return nil, err
	}

	users := []User{}
	for rows.Next() {
		var id string
		var name string
		var email string

		err := rows.Scan(&id, &name, &email)
		if err != nil {
			return nil, err
		}

		users = append(users, User{id, name, email})
	}

	return users, rows.Err()
}

func (d Database) WriteUser(user User) error {
	_, err := d.Pool.Exec(context.Background(), "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.Id, user.Name, user.Email)
	return err
}

func (d Database) WriteExternalUser(user User, mapping UserIdMapping) error {
	txn, err := d.Pool.Begin(context.Background())
	if err != nil {
		return errors.Wrap(err, "unable to being txn for write external user")
	}
	defer txn.Rollback(context.Background())

	_, err = txn.Exec(context.Background(), "INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", user.Id, user.Name, user.Email)
	if err != nil {
		return errors.Wrap(err, "unable to write into users for new external user")
	}

	_, err = txn.Exec(context.Background(), "INSERT INTO external_user_ids (user_id, external_id, source) VALUES ($1, $2, $3)", user.Id, mapping.ExternalId, mapping.Source)
	if err != nil {
		return errors.Wrap(err, "unable to write userid mapping")
	}

	return errors.Wrap(txn.Commit(context.Background()), "unable to commit new external user txn")
}

func (d Database) WriteEventAndMembership(event Event, membership Membership) error {
	txn, err := d.Pool.Begin(context.Background())
	defer txn.Rollback(context.Background())

	if err != nil {
		return errors.Wrap(err, "unable to begin transaction")
	}

	_, err = txn.Exec(context.Background(), "INSERT INTO events (id, name, owner_user_id) VALUES ($1, $2, $3)", event.Id, event.Name, event.OwnerUserId)
	if err != nil {
		return errors.Wrap(err, "unable to write event")
	}

	_, err = txn.Exec(context.Background(), "INSERT INTO memberships (id, event_id, user_id) VALUES ($1, $2, $3)", membership.Id, membership.EventId, membership.UserId)
	if err != nil {
		return errors.Wrap(err, "unable to write membership")
	}

	return errors.Wrap(txn.Commit(context.Background()), "unable to commit txn")
}

func (d Database) ReadEventForUser(eventId string, userId string) (*Event, error) {
	row := d.Pool.QueryRow(context.Background(), "SELECT id, name, owner_user_id FROM events inner join memberships ON event.id = memberships.event_id inner join users ON membership.user_id = users.id WHERE events.id = $1 AND user.id = $2", eventId, userId)

	var id string
	var name string
	var ownerUserId string
	err := row.Scan(&id, &name, &ownerUserId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read event for user")
	}

	return &Event{
		Id:          id,
		Name:        name,
		OwnerUserId: ownerUserId,
	}, nil
}

func (d Database) ReadEventsForUser(userId string) ([]Event, error) {
	rows, err := d.Pool.Query(context.Background(), "SELECT events.id, events.name, events.owner_user_id FROM events INNER JOIN memberships ON events.id = memberships.event_id INNER JOIN users ON memberships.user_id = users.id WHERE users.id = $1", userId)
	if err != nil {
		logrus.WithError(err).Error("Query failed")
		return nil, err
	}

	events := []Event{}
	for rows.Next() {
		var id string
		var name string
		var ownerUserId string

		err := rows.Scan(&id, &name, &ownerUserId)
		if err != nil {
			return nil, err
		}

		events = append(events, Event{id, name, ownerUserId})
	}

	return events, rows.Err()
}

func (d Database) ReadGiftRequestsForEvent(eventId string) ([]GiftRequest, error) {
	rows, err := d.Pool.Query(context.Background(), "SELECT id, name, user_id, assigned_user_id FROM gift_requests where event_id = $1", eventId)
	if err != nil {
		logrus.WithError(err).Error("Query failed")
		return nil, err
	}

	giftRequests := []GiftRequest{}
	for rows.Next() {
		var id string
		var name string
		var userId string
		var assignedUserId *string

		err := rows.Scan(&id, &name, &userId, &assignedUserId)
		if err != nil {
			return nil, err
		}

		giftRequests = append(giftRequests, GiftRequest{
			Id:             id,
			UserId:         userId,
			EventId:        eventId,
			Name:           name,
			AssignedUserId: assignedUserId,
		})
	}

	return giftRequests, rows.Err()
}

func (d Database) WriteGiftRequestEnsureEventMembership(request GiftRequest) error {
	_, err := d.Pool.Exec(context.Background(),
		"INSERT INTO gift_requests (id, user_id, event_id, name, assigned_user_id) "+
			"SELECT * FROM (VALUES ($1::uuid, $2::uuid, $3::uuid, $4, $5::uuid)) i(id, user_id, event_id, name, assigned_user_id) "+
			"WHERE EXISTS ( SELECT FROM memberships WHERE i.user_id = memberships.user_id AND i.event_id = memberships.event_id )",
		request.Id, request.UserId, request.EventId, request.Name, request.AssignedUserId)
	return err
}

func (d Database) DeleteGiftRequest(requestId string, eventId string, userId string) error {
	tag, err := d.Pool.Exec(context.Background(), "DELETE FROM gift_requests WHERE id = $1 AND event_id = $2 AND user_id = $3", requestId, eventId, userId)
	if err != nil {
		return errors.Wrap(err, "unable to delete gift")
	}

	if tag.RowsAffected() != 1 {
		return errors.New("delete gift request deleted incorrect number of gifts: " + strconv.Itoa(int(tag.RowsAffected())))
	}

	return nil
}

func (d Database) ReleaseGiftRequest(requestId string, eventId string, userId string) error {
	tag, err := d.Pool.Exec(context.Background(), "UPDATE gift_requests SET assigned_user_id = NULL WHERE id = $1 AND event_id = $2 AND assigned_user_id = $3 AND user_id != $3", requestId, eventId, userId)
	if err != nil {
		return errors.Wrap(err, "unable to release gift")
	}

	if tag.RowsAffected() != 1 {
		return errors.New("release gift request released incorrect number of gifts: " + strconv.Itoa(int(tag.RowsAffected())))
	}

	return nil
}

func (d Database) ClaimGiftRequest(requestId string, eventId string, userId string) error {
	tag, err := d.Pool.Exec(context.Background(), "UPDATE gift_requests SET assigned_user_id = $1 WHERE id = $2 AND event_id = $3 AND user_id != $1 AND assigned_user_id IS NULL", userId, requestId, eventId)
	if err != nil {
		return errors.Wrap(err, "unable to claim gift")
	}

	if tag.RowsAffected() != 1 {
		return errors.New("claim gift request claimed incorrect number of gifts: " + strconv.Itoa(int(tag.RowsAffected())))
	}

	return nil
}

//UpdateEvent needs to diff existing event membership with expected new members. challenge here is that some of
// the new members might not even be users yet, so we have to delete them.
/**
option 1)

returns a row per input email.
	rows that don't have a users.id value aren't users: create user & add membership
	rows that don't have a membership.id value aren't members: add membership
	rows that have all 3 fields: leave alone
	missing: currents members we have to delete
SELECT i.email, users.id, memberships.id FROM ( VALUES (email1), (email2), ...(emailN in memberEmails) ) AS i(email)
LEFT JOIN users ON i.email = users.email
LEFT JOIN memberships ON users.id = memberships.user_id
WHERE memberships.event_id = $1

$1 = eventId


option 2) WINS, single query gives all 4 cases back

row with no i.email: need to delete that membership (membership.id)
row with only i.email: need to invite user and add membership (i.email)
row with i.email and mem_users.id: dont need to do anything
row with i.email, no mem_users.id, with users.id: add membership (users.id)

SELECT users.id, memberships.id users.email, i.email, users2.id FROM memberships
	INNER JOIN users ON memberships.user_id = users.id AND memberships.event_id = $1
	FULL JOIN ( VALUES (email1), (email2), ...(emailN in memberEmails) ) AS i(email) ON users.email = memberships.email
	LEFT JOIN users2 ON i.email = users2.email

$1 = eventId
*/
func (d Database) UpdateEvent(eventId string, userId string, name string, memberEmails []string) error {
	txn, err := d.Pool.Begin(context.Background())
	if err != nil {
		return errors.Wrap(err, "unable to being txn")
	}
	defer txn.Rollback(context.Background())

	// update event name
	tag, err := txn.Exec(context.Background(), "UPDATE events SET name = $1 WHERE id = $2 AND owner_user_id = $3", name, eventId, userId)
	if err != nil {
		return errors.Wrap(err, "unable to update event name")
	}

	if tag.RowsAffected() != 1 {
		return errors.New("expected to find an event to update")
	}

	// build a state map to determine what to do with existing members and new members
	//memberEmailsInClause := ", ('" + strings.Join(memberEmails, "'), ('") + "')"
	memberEmailsInClause, stateMapQueryParams := valuesClause([]string{eventId}, memberEmails)

	sqlSelect := fmt.Sprintf("SELECT mem_users.id, memberships.id, i.email, users.id FROM memberships "+
		"INNER JOIN users mem_users ON memberships.user_id = mem_users.id AND memberships.event_id = $1 "+
		"FULL JOIN ( VALUES %s ) AS i(email) ON mem_users.email = i.email "+
		"LEFT JOIN users ON i.email = users.email", memberEmailsInClause)
	logrus.Infof("update sql: %s    ... params: %+v", sqlSelect, stateMapQueryParams)

	rows, err := txn.Query(context.Background(), sqlSelect, stateMapQueryParams...)
	if err != nil {
		return errors.Wrap(err, "unable to query event membership state map")
	}

	var memberIdsToRemove []string  //memberId
	var memberUserIdsToAdd []string // userId, eventId
	var usersToInvite []User        // new id, email
	for rows.Next() {
		var memberUserId *string
		var membershipId *string
		var inputEmail *string
		var userId *string
		err = rows.Scan(&memberUserId, &membershipId, &inputEmail, &userId)
		if err != nil {
			return errors.Wrap(err, "unable to scan event membership state map")
		}

		logrus.WithFields(logrus.Fields{"memberUserId": stringPtr(memberUserId), "membershipId": stringPtr(membershipId), "inputEmail": stringPtr(inputEmail), "usersId": stringPtr(userId)}).Info("row")

		if inputEmail == nil {
			// member exists but they aren't supposed to be there anymore, remove
			memberIdsToRemove = append(memberIdsToRemove, *membershipId)
			continue
		}

		if memberUserId != nil {
			// already a member and should be. nothing to do
			continue
		}

		if userId == nil {
			newUser := User{
				Id:    uuid.New().String(),
				Name:  *inputEmail,
				Email: *inputEmail,
			}
			usersToInvite = append(usersToInvite, newUser)
			memberUserIdsToAdd = append(memberUserIdsToAdd, newUser.Id)
		} else {
			memberUserIdsToAdd = append(memberUserIdsToAdd, *userId)
		}
	}

	// remove old members
	if len(memberIdsToRemove) > 0 {
		logrus.WithField("oldMembers", len(memberIdsToRemove)).Info("removing old members")

		deleteMemberIdsInClause, params := inClause([]string{eventId}, memberIdsToRemove)
		_, err = txn.Exec(context.Background(), fmt.Sprintf("DELETE FROM memberships WHERE event_id = $1 AND id IN ( %s )", trimLeftChar(deleteMemberIdsInClause)), params...)
		if err != nil {
			return errors.Wrap(err, "unable to delete bad memberships")
		}
	}

	// invite new users to the platform
	if len(usersToInvite) > 0 {
		logrus.WithField("newUsers", len(usersToInvite)).Info("inviting new users")

		insertNewUsersInClause := ""
		i := 1
		var params []interface{}
		for _, user := range usersToInvite {
			insertNewUsersInClause = insertNewUsersInClause + fmt.Sprintf(", ( $%s, $%s, $%s )", strconv.Itoa(i), strconv.Itoa(i+1), strconv.Itoa(i+2))
			params = append(params, user.Id, user.Name, user.Email)
			i += 3
		}
		insertNewUsersInClause = trimLeftChar(insertNewUsersInClause)

		_, err := txn.Exec(context.Background(), fmt.Sprintf("INSERT INTO users (id, name, email) VALUES %s", insertNewUsersInClause), params...)
		if err != nil {
			return errors.Wrap(err, "unable to create new users")
		}
	}

	// add new memberships
	if len(memberUserIdsToAdd) > 0 {
		logrus.WithField("newMembers", len(memberUserIdsToAdd)).Info("adding new memberships")

		insertNewMembersInClause := ""
		i := 1
		var params []interface{}
		for _, memberUserId := range memberUserIdsToAdd {
			newUuid := uuid.New().String()
			insertNewMembersInClause = insertNewMembersInClause + fmt.Sprintf(", ( $%s, $%s, $%s )", strconv.Itoa(i), strconv.Itoa(i+1), strconv.Itoa(i+2))
			params = append(params, newUuid, eventId, memberUserId)
			i += 3
		}
		insertNewMembersInClause = trimLeftChar(insertNewMembersInClause)

		_, err := txn.Exec(context.Background(), fmt.Sprintf("INSERT INTO memberships (id, event_id, user_id) VALUES %s", insertNewMembersInClause), params...)
		if err != nil {
			return errors.Wrap(err, "unable to create new memberships")
		}
	}

	return errors.Wrap(txn.Commit(context.Background()), "unable to commit transaction")
}

func (d Database) MapExternalIdToUser(externalId string, source AuthSource) (*User, error) {
	row := d.Pool.QueryRow(context.Background(), "SELECT user_id FROM external_user_ids WHERE external_id = $1 AND source = $2", externalId, source)

	var userId string
	err := row.Scan(&userId)
	if err != nil {
		logrus.WithField("err", err.Error()).WithField("sqlerr", sql.ErrNoRows.Error()).Error("error scanning user")
		if noRowsFoundError(err) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "unable to scan userid when mapping external ids")
	}

	return d.ReadUser(userId)
}

type Session struct {
	Id        string
	UserId    string
	CsrfToken string
	Active    bool
}

func (d Database) ReadSession(id string) (*Session, error) {
	row := d.Pool.QueryRow(context.Background(), "SELECT user_id, csrf_token, active FROM sessions WHERE id = $1", id)

	var userId string
	var csrfToken string
	var active bool
	err := row.Scan(&userId, &csrfToken, &active)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read session")
	}

	return &Session{
		Id:        id,
		UserId:    userId,
		CsrfToken: csrfToken,
		Active:    active,
	}, nil
}

func (d Database) CreateSessionAndDeactivateOld(userId string) (*Session, error) {
	txn, err := d.Pool.Begin(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "unable to begin txn")
	}
	defer txn.Rollback(context.Background())

	session := Session{
		Id:        uuid.New().String(),
		UserId:    userId,
		CsrfToken: uuid.New().String(),
		Active:    true,
	}
	_, err = txn.Exec(context.Background(), "INSERT INTO sessions (id, user_id, csrf_token, active) VALUES ($1, $2, $3, $4)", session.Id, userId, session.CsrfToken, session.Active)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create session")
	}

	_, err = txn.Exec(context.Background(), "UPDATE sessions SET active = false WHERE user_id = $1 AND id != $2", userId, session.Id)
	if err != nil {
		return nil, errors.Wrap(err, "unable to update session to deactivate previously active sessions")
	}

	err = txn.Commit(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "unable to commit txn for creating a session")
	}
	return &session, nil
}

func (d Database) ReadSessionForUser(userId string) (*Session, error) {
	row := d.Pool.QueryRow(context.Background(), "SELECT id, csrf_token, active FROM sessions WHERE user_id = $1", userId)

	var id string
	var csrfToken string
	var active bool
	err := row.Scan(&id, &csrfToken, &active)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read session for userid")
	}

	return &Session{
		Id:        id,
		UserId:    userId,
		CsrfToken: csrfToken,
		Active:    active,
	}, nil
}

func noRowsFoundError(err error) bool {
	return strings.Contains(err.Error(), sql.ErrNoRows.Error()) || strings.Contains(sql.ErrNoRows.Error(), err.Error())
}

func valuesClause(initialParams []string, clauseParams []string) (string, []interface{}) {
	resultValuesClause := ""
	var queryParams []interface{}
	for _, param := range initialParams {
		queryParams = append(queryParams, param)
	}

	for i, param := range clauseParams {
		resultValuesClause = resultValuesClause + ", ($" + strconv.Itoa(i+len(initialParams)+1) + ")"
		queryParams = append(queryParams, param)
	}
	resultValuesClause = trimLeftChar(resultValuesClause)
	return resultValuesClause, queryParams
}

func inClause(initialParams []string, clauseParams []string) (string, []interface{}) {
	resultValuesClause := ""
	var queryParams []interface{}
	for _, param := range initialParams {
		queryParams = append(queryParams, param)
	}

	for i, email := range clauseParams {
		resultValuesClause = resultValuesClause + ", $" + strconv.Itoa(i+len(initialParams)+1) + ""
		queryParams = append(queryParams, email)
	}
	resultValuesClause = trimLeftChar(resultValuesClause)
	return resultValuesClause, queryParams
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func stringPtr(s *string) string {
	if s == nil {
		return "nil"
	}
	return *s
}
