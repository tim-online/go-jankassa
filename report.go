package jankassa

import (
	"encoding/json"
	"encoding/xml"
	"time"

	"github.com/aodin/date"
)

type Report struct {
	// XMLName xml.Name `xml:"Report"`

	TimeCreated DateTime `xml:"TimeCreated,attr"`
	DagStart    Date     `xml:"DAG_START"`
	DagStop     Date     `xml:"DAG_STOP"`
	TijdStart   Time     `xml:"TIJD_START"`
	TijdStop    Time     `xml:"TIJD_STOP"`

	// fixed items
	AantalBonnen ReportItem `xml:"Aantal_bonnen"`
	BTWHoog21    ReportItem `xml:"BTW_hoog_21"`
	BTWLaag6     ReportItem `xml:"BTW_laag_6"`
	BTWLaag9     ReportItem `xml:"BTW_laag_9"`
	ExclBTWHoog  ReportItem `xml:"EXCL_BTW_hoog"`
	ExclBTWLaag  ReportItem `xml:"EXCL_BTW_laag"`
	ExclBTWLaag9 ReportItem `xml:"EXCL_BTW_laag_9"`
	TotaalBTW    ReportItem `xml:"Totaal_BTW"`
	BTWVrij      ReportItem `xml:"BTW_vrij"`
	ExclBTWVrij  ReportItem `xml:"EXCL_BTW_vrij"`
	Totaal       ReportItem `xml:"Totaal"`
	Bruto        ReportItem `xml:"Bruto"`
	Gemiddeld    ReportItem `xml:"Gemiddeld"`
	Netto        ReportItem `xml:"Netto"`

	// Custom per sync
	CustomItems ReportItems `xml:",any"`
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

type ReportItems []ReportItem

func (items ReportItems) GetByName(name string) (ReportItem, bool) {
	for _, item := range items {
		if item.Name == name {
			return item, true
		}
	}

	return ReportItem{}, false
}

func (items ReportItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias ReportItems

	for _, item := range items {
		start.Name.Local = item.Name
		err := e.EncodeElement(item, start)
		if err != nil {
			return err
		}
	}

	return nil
}

type ReportItem struct {
	Name  string  `xml:"-"`
	Code  string  `xml:"code,attr"`
	Value float64 `xml:",chardata"`
}

func (ri *ReportItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias ReportItem
	alias := (*Alias)(ri)
	ri.Name = start.Name.Local

	err := d.DecodeElement(alias, &start)
	if err != nil {
		return err
	}
	return nil
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
