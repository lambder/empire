package saml

import (
	"crypto/rand"
	"time"
)

// TimeNow is a function that returns the current time. The default
// value is time.Now, but it can be replaced for testing.
var TimeNow = func() time.Time { return time.Now().UTC() }

// RandReader is the io.Reader that produces cryptographically random
// bytes when they are need by the library. The default value is
// rand.Reader, but it can be replaced for testing.
var RandReader = rand.Reader

func randomBytes(n int) []byte {
	rv := make([]byte, n)
	if _, err := RandReader.Read(rv); err != nil {
		panic(err)
	}
	return rv
}
