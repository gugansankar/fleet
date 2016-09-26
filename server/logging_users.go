package server

import (
	"time"

	"github.com/kolide/kolide-ose/kolide"
	"github.com/kolide/kolide-ose/server/contexts/viewer"
	"golang.org/x/net/context"
)

func (mw loggingMiddleware) NewUser(ctx context.Context, p kolide.UserPayload) (*kolide.User, error) {
	var (
		user     *kolide.User
		err      error
		username = "none"
	)

	vc, ok := viewer.FromContext(ctx)
	if !ok {
		return nil, errNoContext
	}

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "NewUser",
			"user", username,
			"created_by", vc.Username(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	user, err = mw.Service.NewUser(ctx, p)

	if user != nil {
		username = user.Username
	}
	return user, err
}

func (mw loggingMiddleware) ModifyUser(ctx context.Context, userID uint, p kolide.UserPayload) (*kolide.User, error) {
	var (
		user     *kolide.User
		err      error
		username = "none"
	)

	vc, ok := viewer.FromContext(ctx)
	if !ok {
		return nil, errNoContext
	}

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ModifyUser",
			"user", username,
			"modified_by", vc.Username(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	user, err = mw.Service.ModifyUser(ctx, userID, p)

	if user != nil {
		username = user.Username
	}

	return user, err
}

func (mw loggingMiddleware) User(ctx context.Context, id uint) (*kolide.User, error) {
	var (
		user     *kolide.User
		err      error
		username = "none"
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "User",
			"user", username,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	user, err = mw.Service.User(ctx, id)

	if user != nil {
		username = user.Username
	}
	return user, err
}

func (mw loggingMiddleware) AuthenticatedUser(ctx context.Context) (*kolide.User, error) {
	var (
		user     *kolide.User
		err      error
		username = "none"
	)

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "User",
			"user", username,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	user, err = mw.Service.AuthenticatedUser(ctx)

	if user != nil {
		username = user.Username
	}
	return user, err
}

func (mw loggingMiddleware) ResetPassword(ctx context.Context, token, password string) error {
	var err error

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "ChangePassword",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.ResetPassword(ctx, token, password)
	return err
}

func (mw loggingMiddleware) RequestPasswordReset(ctx context.Context, email string) error {
	var (
		requestedBy = "unauthenticated"
		err         error
	)
	vc, ok := viewer.FromContext(ctx)
	if ok {
		requestedBy = vc.Username()
	}

	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "RequestPasswordReset",
			"email", email,
			"err", err,
			"requested_by", requestedBy,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.RequestPasswordReset(ctx, email)
	return err
}