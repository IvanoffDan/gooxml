// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/baliance/gooxml"
)

type CT_IndexedColors struct {
	// RGB Color
	RgbColor []*CT_RgbColor
}

func NewCT_IndexedColors() *CT_IndexedColors {
	ret := &CT_IndexedColors{}
	return ret
}

func (m *CT_IndexedColors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sergbColor := xml.StartElement{Name: xml.Name{Local: "ma:rgbColor"}}
	for _, c := range m.RgbColor {
		e.EncodeElement(c, sergbColor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IndexedColors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_IndexedColors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rgbColor"}:
				tmp := NewCT_RgbColor()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RgbColor = append(m.RgbColor, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_IndexedColors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_IndexedColors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_IndexedColors and its children
func (m *CT_IndexedColors) Validate() error {
	return m.ValidateWithPath("CT_IndexedColors")
}

// ValidateWithPath validates the CT_IndexedColors and its children, prefixing error messages with path
func (m *CT_IndexedColors) ValidateWithPath(path string) error {
	for i, v := range m.RgbColor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RgbColor[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
