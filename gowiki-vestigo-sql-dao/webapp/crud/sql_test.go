package crud

import "testing"

func TestGetConnection(t *testing.T) {
	con, err := GetConnection()
	if(err != nil){
		t.Error("the connection has failed", err)
	}

	con.Close()
}
