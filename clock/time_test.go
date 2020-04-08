package clock

import (
	"testing"
	"time"

	"github.com/yuukanoo/gobb/assert"
)

func TestClockWithoutFunc(t *testing.T) {
	var ti Time

	assert.Equals(t, false, ti.Now().IsZero(), "builtin time should have been called")
	assert.Equals(t, time.UTC, ti.Now().Location(), "builtin time should have been called using UTC")
}

func TestClockWithFunc(t *testing.T) {
	before := time.Now().UTC().Add(-5 * time.Hour)
	ti := Time{
		NowFn: func() time.Time { return before },
	}

	assert.Equals(t, before, ti.Now(), "should use the given NowFn")
}
