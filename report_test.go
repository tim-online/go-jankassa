package jankassa_test

import (
	"encoding/xml"
	"log"
	"os"
	"testing"

	jankassa "github.com/tim-online/go-jankassa"
)

func TestDing(t *testing.T) {
	file, err := os.Open("report.xml")
	if err != nil {
		t.Error(err)
	}

	report := jankassa.Report{}
	dec := xml.NewDecoder(file)
	err = dec.Decode(&report)
	if err != nil {
		t.Error(err)
	}

	b, _ := xml.MarshalIndent(report, "", "    ")
	log.Println(string(b))
}
