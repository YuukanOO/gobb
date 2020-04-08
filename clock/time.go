// Package clock provides a simple mockable clock to be used anywhere where the
// time is needed.
package clock

import "time"

// Time struct which should be used to retrieve time in components. It makes
// it easy to mock the time by setting the NowFn member.
//
// Testing time based stuff is really painful and I like this approach more than
// any other so here it is :)
type Time struct {
	NowFn func() time.Time
}

// Now returns the current time, UTC based.
func (t *Time) Now() time.Time {
	if t.NowFn == nil {
		return time.Now().UTC()
	}

	return t.NowFn()
}
