// Package clock provides a simple mockable clock to be used anywhere where the
// time is needed.
package clock

import "time"

type (
	// Time struct which should be used to retrieve time in components. It makes
	// it easy to mock the time by setting the NowFn member.
	//
	// Testing time based stuff is really painful and I like this approach more than
	// any other so here it is :)
	Time struct {
		NowFn func() time.Time
	}

	// UTCTime is a value object representing an UTC time. This makes it clear that
	// a given date is expected to be UTC based. It embed the original time.Time to
	// comply with interfaces.
	//
	// You should always instantiate it with the associated method NewUTCTime which
	// will validate the input.
	UTCTime struct {
		time.Time
	}
)

// Now returns the current time, UTC based.
func (t *Time) Now() UTCTime {
	if t.NowFn == nil {
		// Forget the ctor here since we already know its an UTC time.
		return UTCTime{time.Now().UTC()}
	}

	return NewUTCTime(t.NowFn())
}

// NewUTCTime returns an UTCTime value object making sure the given time is
// UTC based.
func NewUTCTime(value time.Time) UTCTime {
	return UTCTime{value.UTC()}
}

// After reports whether the time instant t is after u.
func (t UTCTime) After(u UTCTime) bool {
	return t.Time.After(u.Time)
}

// Before reports whether the time instant t is before u.
func (t UTCTime) Before(u UTCTime) bool {
	return t.Time.Before(u.Time)
}

// Add returns the time t+d.
func (t UTCTime) Add(d time.Duration) UTCTime {
	return UTCTime{t.Time.Add(d)}
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration will
// be returned. To compute t-d for a duration d, use t.Add(-d).
func (t UTCTime) Sub(u UTCTime) time.Duration {
	return t.Time.Sub(u.Time)
}

// Truncate returns the result of rounding t down to a multiple of d (since the
// zero time).
func (t UTCTime) Truncate(d time.Duration) UTCTime {
	return UTCTime{t.Time.Truncate(d)}
}
