// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_TablePartStyle struct {
	TcTxStyle *CT_TableStyleTextStyle
	TcStyle   *CT_TableStyleCellStyle
}

func NewCT_TablePartStyle() *CT_TablePartStyle {
	ret := &CT_TablePartStyle{}
	return ret
}
func (m *CT_TablePartStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.TcTxStyle != nil {
		setcTxStyle := xml.StartElement{Name: xml.Name{Local: "a:tcTxStyle"}}
		e.EncodeElement(m.TcTxStyle, setcTxStyle)
	}
	if m.TcStyle != nil {
		setcStyle := xml.StartElement{Name: xml.Name{Local: "a:tcStyle"}}
		e.EncodeElement(m.TcStyle, setcStyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TablePartStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TablePartStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tcTxStyle":
				m.TcTxStyle = NewCT_TableStyleTextStyle()
				if err := d.DecodeElement(m.TcTxStyle, &el); err != nil {
					return err
				}
			case "tcStyle":
				m.TcStyle = NewCT_TableStyleCellStyle()
				if err := d.DecodeElement(m.TcStyle, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TablePartStyle
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TablePartStyle) Validate() error {
	return m.ValidateWithPath("CT_TablePartStyle")
}
func (m *CT_TablePartStyle) ValidateWithPath(path string) error {
	if m.TcTxStyle != nil {
		if err := m.TcTxStyle.ValidateWithPath(path + "/TcTxStyle"); err != nil {
			return err
		}
	}
	if m.TcStyle != nil {
		if err := m.TcStyle.ValidateWithPath(path + "/TcStyle"); err != nil {
			return err
		}
	}
	return nil
}