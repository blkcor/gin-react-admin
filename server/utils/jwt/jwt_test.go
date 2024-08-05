package jwt

import (
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(1, "blkcor", "blkcor.dev@gmail.com")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	claim, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImJsa2NvciIsImVtYWlsIjoiYmxrY29yLmRldkBnbWFpbC5jb20iLCJpc3MiOiJzb21lYm9keSIsImV4cCI6MTcyMjg2OTYzNH0.cj2sBErUQijxUYK8FS2NAtF5zFzKLJjRetFa5VxdSwA")
	if err != nil {
		t.Error(err)
	}
	t.Log(claim)
}
