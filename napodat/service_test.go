package napodate

import (
	"database/sql/driver"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	//testing status
	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T){
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil{
		t.Errorf("Error: %s", err)
	}
	
	time := time.Now()
	today := time.Format("02/01/2006")

	//testing today's date
	ok := today == driver.Bool.String()
	if != ok{
		t.Errorf("expected dates to be equal")
	}
}
