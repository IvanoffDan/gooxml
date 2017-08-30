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
)

type CT_FunctionGroups struct {
	// Built-in Function Group Count
	BuiltInGroupCountAttr *uint32
	// Function Group
	FunctionGroup []*CT_FunctionGroup
}

func NewCT_FunctionGroups() *CT_FunctionGroups {
	ret := &CT_FunctionGroups{}
	return ret
}
func (m *CT_FunctionGroups) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.BuiltInGroupCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "builtInGroupCount"},
			Value: fmt.Sprintf("%v", *m.BuiltInGroupCountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.FunctionGroup != nil {
		sefunctionGroup := xml.StartElement{Name: xml.Name{Local: "x:functionGroup"}}
		e.EncodeElement(m.FunctionGroup, sefunctionGroup)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_FunctionGroups) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "builtInGroupCount" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := uint32(parsed)
			m.BuiltInGroupCountAttr = &pt
		}
	}
lCT_FunctionGroups:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "functionGroup":
				tmp := NewCT_FunctionGroup()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FunctionGroup = append(m.FunctionGroup, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_FunctionGroups
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_FunctionGroups) Validate() error {
	return m.ValidateWithPath("CT_FunctionGroups")
}
func (m *CT_FunctionGroups) ValidateWithPath(path string) error {
	for i, v := range m.FunctionGroup {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FunctionGroup[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}