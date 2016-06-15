package crud

import "testing"

func TestGetConnection(t *testing.T) {
	con, err := GetConnection()
	if(err != nill){
		t.Error("the connection has failed", err)
	}
}
