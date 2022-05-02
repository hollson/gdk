package try

import (
	"errors"
)

var errMaxRetriesReached = errors.New("exceeded retry limit")

// MaxRetries is the maximum number of retries before bailing.
var MaxRetries = 3

// Do keeps trying the function until the second argument
// returns false, or no error is returned.
func Do(fn func(attempt int) (retry bool, err error)) error {
	var (
		err  error
		cont bool
	)

	attempt := 1
	for {
		cont, err = fn(attempt)
		if !cont || err == nil {
			break
		}
		attempt++
		if attempt > MaxRetries {
			return errMaxRetriesReached
		}
	}

	return err
}

// IsMaxRetries checks whether the error is due to hitting the
// maximum number of retries or not.
func IsMaxRetries(err error) bool {
	return err == errMaxRetriesReached
}
