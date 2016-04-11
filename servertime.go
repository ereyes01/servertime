package servertime

import (
	"encoding/json"
	"time"
)

// allows us to mock out the call to time.Now() for a consistent test
type timeNow interface {
	GetCurrentTime() time.Time
}

type realTimeNow struct{}

func (r *realTimeNow) GetCurrentTime() time.Time {
	return time.Now()
}

var now timeNow

func init() {
	now = new(realTimeNow)
}

// ServerTime is mostly just a time.Time value, except that when it's zero, it
// JSON-unmarshals into the equivalent of time.Now() on the server doing the
// unmarshaling. This is intended for server-client applications, where clients
// wish to send objects with timestamps to the server, and have the server
// compute consistent timestamps using its own time, much like:
// https://www.firebase.com/docs/web/api/servervalue/timestamp.html
type ServerTime struct {
	time.Time
}

// UnmarshalJSON turns an RFC 3339 compliant JSON string into a time value. If
// the unmarshaled value is a time.Time zero value, the server's current time
// is unmarshaled instead.
func (s *ServerTime) UnmarshalJSON(b []byte) error {
	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	if t.IsZero() {
		t = now.GetCurrentTime().UTC()
	}

	*s = ServerTime{t}
	return nil
}
