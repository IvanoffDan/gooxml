// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"
)

type CT_ObjectAnchor struct {
	// Move With Cells
	MoveWithCellsAttr *bool
	// Size With Cells
	SizeWithCellsAttr *bool
	From              *spreadsheetDrawing.From
	To                *spreadsheetDrawing.To
}

func NewCT_ObjectAnchor() *CT_ObjectAnchor {
	ret := &CT_ObjectAnchor{}
	ret.From = spreadsheetDrawing.NewFrom()
	ret.To = spreadsheetDrawing.NewTo()
	return ret
}
func (m *CT_ObjectAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.MoveWithCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "moveWithCells"},
			Value: fmt.Sprintf("%v", *m.MoveWithCellsAttr)})
	}
	if m.SizeWithCellsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sizeWithCells"},
			Value: fmt.Sprintf("%v", *m.SizeWithCellsAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	sefrom := xml.StartElement{Name: xml.Name{Local: "xdr:from"}}
	e.EncodeElement(m.From, sefrom)
	seto := xml.StartElement{Name: xml.Name{Local: "xdr:to"}}
	e.EncodeElement(m.To, seto)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_ObjectAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = spreadsheetDrawing.NewFrom()
	m.To = spreadsheetDrawing.NewTo()
	for _, attr := range start.Attr {
		if attr.Name.Local == "moveWithCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.MoveWithCellsAttr = &parsed
		}
		if attr.Name.Local == "sizeWithCells" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SizeWithCellsAttr = &parsed
		}
	}
lCT_ObjectAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "from":
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case "to":
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ObjectAnchor
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_ObjectAnchor) Validate() error {
	return m.ValidateWithPath("CT_ObjectAnchor")
}
func (m *CT_ObjectAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.To.ValidateWithPath(path + "/To"); err != nil {
		return err
	}
	return nil
}