// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"time"
)

type CT_TblPrExChange struct {
	AuthorAttr string
	DateAttr   *time.Time
	// Annotation Identifier
	IdAttr  int32
	TblPrEx *CT_TblPrExBase
}

func NewCT_TblPrExChange() *CT_TblPrExChange {
	ret := &CT_TblPrExChange{}
	ret.TblPrEx = NewCT_TblPrExBase()
	return ret
}
func (m *CT_TblPrExChange) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:author"},
		Value: fmt.Sprintf("%v", m.AuthorAttr)})
	if m.DateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:date"},
			Value: fmt.Sprintf("%v", *m.DateAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	setblPrEx := xml.StartElement{Name: xml.Name{Local: "w:tblPrEx"}}
	e.EncodeElement(m.TblPrEx, setblPrEx)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TblPrExChange) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TblPrEx = NewCT_TblPrExBase()
	for _, attr := range start.Attr {
		if attr.Name.Local == "author" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AuthorAttr = parsed
		}
		if attr.Name.Local == "date" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.DateAttr = &parsed
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = int32(parsed)
		}
	}
lCT_TblPrExChange:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tblPrEx":
				if err := d.DecodeElement(m.TblPrEx, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblPrExChange
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TblPrExChange) Validate() error {
	return m.ValidateWithPath("CT_TblPrExChange")
}
func (m *CT_TblPrExChange) ValidateWithPath(path string) error {
	if err := m.TblPrEx.ValidateWithPath(path + "/TblPrEx"); err != nil {
		return err
	}
	return nil
}