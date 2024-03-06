package entity

import "errors"

var (
	// ErrFeedNotExists error feed not exists
	ErrFeedNotExists = errors.New("feed not exists")
	// ErrUserNotExists error user not exists
	ErrUserNotExists = errors.New("user not exists")
	// ErrUnknownBotCommand error user not exists
	ErrUnknownBotCommand = errors.New("unknown bot command")
	// ErrSubscriptionAlreadyExists error subscription already exists
	ErrSubscriptionAlreadyExists = errors.New("subscription already exists")
	// ErrSubscriptionNotExists error subscription not exists
	ErrSubscriptionNotExists = errors.New("subscription not exists")
	// ErrWrongSubscriptionNumber failed to load URL
	ErrWrongSubscriptionNumber = errors.New("wrong subscription number")
	// ErrNotCommand error not a command
	ErrNotCommand = errors.New("not a command")
	// ErrWrongNumberOfArguments error wrong number of arguments
	ErrWrongNumberOfArguments = errors.New("wrong number of arguments")
	// ErrWrongURL error wrong URL
	ErrWrongURL = errors.New("wrong URL")
	// ErrLoadURL failed to load URL
	ErrLoadURL = errors.New("failed to load URL")
	// ErrInternal internal error
	ErrInternal = errors.New("internal error")
	// ErrNoFeedItems no feed items received
	ErrNoFeedItems = errors.New("no feed items received")
	// ErrNoFeed the site does not contain RSS feed
	ErrNoFeed = errors.New("the site does not contain RSS feed")
)

func GetErrorMessage(err error) string {
	switch err {
	case ErrUnknownBotCommand, ErrNotCommand, ErrWrongNumberOfArguments:
		return MessageUnknownCommand
	case ErrSubscriptionAlreadyExists:
		return MessageSubscriptionAlreadyExists
	case ErrWrongSubscriptionNumber:
		return MessageWrongSubscriptionNumber
	case ErrWrongURL:
		return MessageWrongURL
	case ErrNoFeed:
		return MessageNoFeed
	case ErrLoadURL:
		return MessageErrorLoadURL
	case ErrInternal:
		return MessageInternalError
	default:
		return ""
	}
}
