package entity

import (
	"strconv"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Subscription struct {
	ID        string
	UserID    int64
	FeedID    string
	CreatedAt time.Time
}

func ValidateURL(url string) error {
	return validation.Validate(url, validation.Required, is.URL)
}

func ValidateNumber(num string) (int, error) {
	n, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, err
	} else if n < 1 {
		return 0, ErrWrongSubscriptionNumber
	}
	return int(n), nil
}
