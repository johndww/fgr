package conversion

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"server/pkg"
)

type Runner struct {
	Database *pkg.Database
}

func (r Runner) RunConversions() error {
	err := r.giftRequestsDescription()
	if err != nil {
		return err
	}

	err = r.eventsOwnerUserIdUUID()
	if err != nil {
		return err
	}

	err = r.demoUser()
	if err != nil {
		return err
	}

	return nil
}

func (r Runner) giftRequestsDescription() error {
	_, err := r.Database.Pool.Exec(context.Background(), "ALTER TABLE gift_requests ADD COLUMN IF NOT EXISTS description TEXT")
	if err != nil {
		return errors.Wrap(err, "added gift request description column")
	}

	logrus.Info("ran gift requests description conversion")
	return nil
}

func (r Runner) eventsOwnerUserIdUUID() error {
	_, err := r.Database.Pool.Exec(context.Background(), "ALTER TABLE events "+
		"ALTER COLUMN owner_user_id SET DATA TYPE uuid USING owner_user_id::uuid;")
	if err != nil {
		return errors.Wrap(err, "added gift request description column")
	}

	logrus.Info("ran events owner user id uuid conversion")
	return nil
}

func (r Runner) demoUser() error {
	_, err := r.Database.Pool.Exec(context.Background(), "ALTER TABLE users ADD COLUMN IF NOT EXISTS demo BOOLEAN DEFAULT false")
	if err != nil {
		return errors.Wrap(err, "unable to add demo column")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO users (id, name, email, admin, demo) VALUES ($1, $2, $3, $4, $5) on conflict (id) do nothing", pkg.DemoUserId1, "Joe", "demo-user-joe@simplegift.app", false, true)
	if err != nil {
		return errors.Wrap(err, "unable to add demo user 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO users (id, name, email, admin, demo) VALUES ($1, $2, $3, $4, $5) on conflict (id) do nothing", pkg.DemoUserId2, "Henry", "demo-user-henry@simplegift.app", false, true)
	if err != nil {
		return errors.Wrap(err, "unable to add demo user 2")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO users (id, name, email, admin, demo) VALUES ($1, $2, $3, $4, $5) on conflict (id) do nothing", pkg.DemoUserId3, "Emily", "demo-user-emily@simplegift.app", false, true)
	if err != nil {
		return errors.Wrap(err, "unable to add demo user 3")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO events (id, name, owner_user_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoEventId1, "Christmas", pkg.DemoUserId1)
	if err != nil {
		return errors.Wrap(err, "unable to add demo event 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO memberships (id, user_id, event_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoMembershipEvent1User1, pkg.DemoUserId1, pkg.DemoEventId1)
	if err != nil {
		return errors.Wrap(err, "unable to add demo membership event 1 user 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO memberships (id, user_id, event_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoMembershipEvent1User2, pkg.DemoUserId2, pkg.DemoEventId1)
	if err != nil {
		return errors.Wrap(err, "unable to add demo membership event 1 user 2")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO events (id, name, owner_user_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoEventId2, "Mike's Birthday", pkg.DemoUserId2)
	if err != nil {
		return errors.Wrap(err, "unable to add demo event 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO memberships (id, user_id, event_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoMembershipEvent2User1, pkg.DemoUserId1, pkg.DemoEventId2)
	if err != nil {
		return errors.Wrap(err, "unable to add demo membership event 2 user 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO memberships (id, user_id, event_id) VALUES ($1, $2, $3) on conflict (id) do nothing", pkg.DemoMembershipEvent2User3, pkg.DemoUserId3, pkg.DemoEventId2)
	if err != nil {
		return errors.Wrap(err, "unable to add demo membership event 2 user 3")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO gift_requests (id, name, event_id, user_id, assigned_user_id, description) VALUES ($1, $2, $3, $4, $5, $6) on conflict (id) do nothing", pkg.DemoGiftRequestEvent1Id1, "Xbox Series X", pkg.DemoEventId1, pkg.DemoUserId1, pkg.DemoUserId2, "Easiest to buy on Amazon. No gamepass needed - I already have one!")
	if err != nil {
		return errors.Wrap(err, "unable to add demo gift request 1 event 1")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO gift_requests (id, name, event_id, user_id, assigned_user_id, description) VALUES ($1, $2, $3, $4, $5, $6) on conflict (id) do nothing", pkg.DemoGiftRequestEvent1Id2, "PlayStation 5", pkg.DemoEventId1, pkg.DemoUserId2, pkg.DemoUserId1, "With the CD player please! I don't have a separate blue-ray player")
	if err != nil {
		return errors.Wrap(err, "unable to add demo gift request 2 (event 1)")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO gift_requests (id, name, event_id, user_id, assigned_user_id, description) VALUES ($1, $2, $3, $4, $5, $6) on conflict (id) do nothing", pkg.DemoGiftRequestEvent1Id3, "Nintendo Switch", pkg.DemoEventId1, pkg.DemoUserId2, nil, "I already have controllers - no need for extras")
	if err != nil {
		return errors.Wrap(err, "unable to add demo gift request 3 (event 1)")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO gift_requests (id, name, event_id, user_id, assigned_user_id, description) VALUES ($1, $2, $3, $4, $5, $6) on conflict (id) do nothing", pkg.DemoGiftRequestEvent2Id1, "Reedin Super E Kiteboard", pkg.DemoEventId2, pkg.DemoUserId1, nil, "Size 140x42cm")
	if err != nil {
		return errors.Wrap(err, "unable to add demo gift request 1 event 2")
	}

	_, err = r.Database.Pool.Exec(context.Background(), "INSERT INTO gift_requests (id, name, event_id, user_id, assigned_user_id, description) VALUES ($1, $2, $3, $4, $5, $6) on conflict (id) do nothing", pkg.DemoGiftRequestEvent2Id2, "MacBook Pro - M2", pkg.DemoEventId2, pkg.DemoUserId3, nil, "14 inch, 12-core, Space Grey, 32GB RAM, 1TB SSD")
	if err != nil {
		return errors.Wrap(err, "unable to add demo gift request 2 event 2")
	}

	logrus.Info("ran demo user conversion")
	return nil
}
