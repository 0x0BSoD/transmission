package transmission

import (
	"testing"

	"github.com/kr/pretty"
)

func TestSession(t *testing.T) {
	// Let's create a simple client
	conf := Config{
		Address:  "http://localhost:9091/transmission/rpc",
		User:     "admin",
		Password: "admin",
	}
	trans, err := New(conf)
	if err != nil {
		t.Fatal(err)
	}
	// Update and print the current session
	trans.Session.Update()
	pretty.Println(trans.Session)
}
