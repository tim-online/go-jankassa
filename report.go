package jankassa

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/aodin/date"
)

type Report struct {
	// XMLName xml.Name `xml:"Report"`

	TimeCreated  DateTime     `xml:"TimeCreated,attr"`
	AantalBonnen AantalBonnen `xml:"Aantal_bonnen"`
	BTWHoog21    BTWHoog21    `xml:"BTW_hoog_21"`
	BTWLaag6     BTWLaag6     `xml:"BTW_laag_6"`
	Bruto        Bruto        `xml:"Bruto"`
	DagStart     Date         `xml:"DAG_START"`
	DagStop      Date         `xml:"DAG_STOP"`
	ExclBTWHoog  ExclBTWHoog  `xml:"EXCL_BTW_hoog"`
	ExclBTWLaag  ExclBTWLaag  `xml:"EXCL_BTW_laag"`
	GeldInLade   GeldInLade   `xml:"Geld_in_lade"`
	Gemiddeld    Gemiddeld    `xml:"Gemiddeld"`
	Keuken       Keuken       `xml:"Keuken"`
	Netto        Netto        `xml:"Netto"`
	TijdStart    Time         `xml:"TIJD_START"`
	TijdStop     Time         `xml:"TIJD_STOP"`
}

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	var err error
	layout := "02/01/2006 15:04:05"
	dt.Time, err = time.Parse(layout, attr.Value)
	return err
}

func (dt DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "02/01/2006 15:04:05"
	return e.EncodeElement(dt.Time.Format(layout), start)
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	layout := "02/01/2006 15:04:05"
	dt.Time, err = time.Parse(layout, value)
	return err
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	layout := "02/01/2006 15:04:05"
	return json.Marshal(dt.Time.Format(layout))
}

type AantalBonnen struct {
	Code  string `xml:"code,attr"`
	Value string `xml:",chardata"`
}

type Gemiddeld struct {
	Code  string `xml:"code,attr"`
	Value string `xml:",chardata"`
}

type Bruto struct {
	Code  string `xml:"code,attr"`
	Value string `xml:",chardata"`
}

type Date struct {
	date.Date
}

func (dt *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	layout := "02.01.2006"
	t, err := time.Parse(layout, value)
	if err != nil {
		return err
	}

	dt.Date = date.FromTime(t)
	return nil
}

func (dt Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "02.01.2006"
	return e.EncodeElement(dt.Time.Format(layout), start)
}

func (dt *Date) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &dt.Time)
}

func (dt Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Time)
}

type Netto struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type GeldInLade struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type ExclBTWLaag struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type BTWLaag6 struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type ExclBTWHoog struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type BTWHoog21 struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type Keuken struct {
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

type Time struct {
	time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var value string
	err := d.DecodeElement(&value, &start)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	layout := "15:04"
	t.Time, err = time.Parse(layout, value)
	return err
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	layout := "15:04"
	return e.EncodeElement(t.Time.Format(layout), start)
}
