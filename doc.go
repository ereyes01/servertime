package servertime

// Package servertime is a small wrapper around the time.Time type that behaves
// just like time.Time, except that if a servertime.Time value is 0, it is
// initialized when JSON unmarshalled to the caller's time.Now(). This is
// useful for applications that communicate with a server and want to have
// consistent timestamps, arbitrated from a single server.
//
// This functionality is very similar (if not the same, logically) to that offered by
// Firebase's ServerValue.TIMESTAMP feature
// (https://www.firebase.com/docs/web/api/servervalue/timestamp.html). Unlike
// Firebase's implementation, this library is 100% conformant to how Go
// marhals / unmarshals timestamps.
