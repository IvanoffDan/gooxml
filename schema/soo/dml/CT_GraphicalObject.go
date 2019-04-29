// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/IvanoffDan/gooxml"
)

type CT_GraphicalObject struct {
	GraphicData *CT_GraphicalObjectData
}

func NewCT_GraphicalObject() *CT_GraphicalObject {
	ret := &CT_GraphicalObject{}
	ret.GraphicData = NewCT_GraphicalObjectData()
	return ret
}

func (m *CT_GraphicalObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	segraphicData := xml.StartElement{Name: xml.Name{Local: "a:graphicData"}}
	e.EncodeElement(m.GraphicData, segraphicData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObject) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.GraphicData = NewCT_GraphicalObjectData()
lCT_GraphicalObject:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphicData"}:
				if err := d.DecodeElement(m.GraphicData, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_GraphicalObject %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObject
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObject and its children
func (m *CT_GraphicalObject) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObject")
}

// ValidateWithPath validates the CT_GraphicalObject and its children, prefixing error messages with path
func (m *CT_GraphicalObject) ValidateWithPath(path string) error {
	if err := m.GraphicData.ValidateWithPath(path + "/GraphicData"); err != nil {
		return err
	}
	return nil
}
