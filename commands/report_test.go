package commands

import (
	"bytes"
	"log"
	"testing"
)

func TestReportParseGood(t *testing.T) {
	c := Report{}
	if c.Parse("REPORT") != true {
		t.Errorf("Should have parsed REPORT.")
	}
}
func TestReportParseBad(t *testing.T) {
	c := Report{}
	if c.Parse("REPORTS") == true {
		t.Errorf("Shouldn't have parsed REPORTS.")
	}
}

func TestReportInitialize(t *testing.T) {
	state := testInitialModel()
	c := Report{}

	var str bytes.Buffer
	log.SetOutput(&str)

	c.Execute(state)

	if str.String() != "" {
		t.Errorf("Reported while in an invalid model")
	}
}

func TestReportValid(t *testing.T) {
	state := testModel()
	c := Report{}

	var str bytes.Buffer
	log.SetOutput(&str)

	c.Execute(state)

	if str.String() != "Robot at 2 2 NORTH\n" {
		t.Errorf("Wrong position reported %s", str.String())
	}
}
