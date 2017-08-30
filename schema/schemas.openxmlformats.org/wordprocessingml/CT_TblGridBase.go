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
)

type CT_TblGridBase struct {
	// Grid Column Definition
	GridCol []*CT_TblGridCol
}

func NewCT_TblGridBase() *CT_TblGridBase {
	ret := &CT_TblGridBase{}
	return ret
}
func (m *CT_TblGridBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.GridCol != nil {
		segridCol := xml.StartElement{Name: xml.Name{Local: "w:gridCol"}}
		e.EncodeElement(m.GridCol, segridCol)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TblGridBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblGridBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "gridCol":
				tmp := NewCT_TblGridCol()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridCol = append(m.GridCol, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblGridBase
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TblGridBase) Validate() error {
	return m.ValidateWithPath("CT_TblGridBase")
}
func (m *CT_TblGridBase) ValidateWithPath(path string) error {
	for i, v := range m.GridCol {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridCol[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}