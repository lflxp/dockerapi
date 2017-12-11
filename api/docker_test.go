package api 

import (
	"testing"
)

func Test_Ps(t *testing.T) {
	data,err := Ps()
	if err != nil {
		t.Error(err.Error())
	}
	for _,con := range data {
		t.Log(con.Command,con.ID,con.Image,con.Ports,con.NetworkSettings)
	}
}