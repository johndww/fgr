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
