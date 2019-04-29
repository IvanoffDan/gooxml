// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/baliance/gooxml"
)

type OfcCT_Rules struct {
	R       []*OfcCT_R
	ExtAttr ST_Ext
}

func NewOfcCT_Rules() *OfcCT_Rules {
	ret := &OfcCT_Rules{}
	return ret
}

func (m *OfcCT_Rules) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.R != nil {
		ser := xml.StartElement{Name: xml.Name{Local: "o:r"}}
		for _, c := range m.R {
			e.EncodeElement(c, ser)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Rules) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_Rules:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "r"}:
				tmp := NewOfcCT_R()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.R = append(m.R, tmp)
			default:
				gooxml.Log("skipping unsupported element on OfcCT_Rules %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_Rules
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_Rules and its children
func (m *OfcCT_Rules) Validate() error {
	return m.ValidateWithPath("OfcCT_Rules")
}

// ValidateWithPath validates the OfcCT_Rules and its children, prefixing error messages with path
func (m *OfcCT_Rules) ValidateWithPath(path string) error {
	for i, v := range m.R {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/R[%d]", path, i)); err != nil {
			return err
		}
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
