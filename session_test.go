package transmission

import (
	"context"
	"testing"

	"github.com/kr/pretty"
)

func TestSession(t *testing.T) {
	// Let's create a simple client
	conf := Config{
		Address:  "http://192.168.11.7:9091/transmission/rpc",
		User:     "admin",
		Password: "admin",
	}
	trans, err := New(conf)
	if err != nil {
		t.Fatal(err)
	}
	trans.Context = context.TODO()

	// Update and print the current session
	err = trans.Session.Update()
	if err != nil {
		t.Fatal(err)
	}

	pretty.Println(trans.Session)
}
