// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_OfficeArtExtensionList struct {
	Ext []*CT_OfficeArtExtension
}

func NewCT_OfficeArtExtensionList() *CT_OfficeArtExtensionList {
	ret := &CT_OfficeArtExtensionList{}
	return ret
}
func (m *CT_OfficeArtExtensionList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "a:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_OfficeArtExtensionList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_OfficeArtExtensionList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "ext":
				tmp := NewCT_OfficeArtExtension()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ext = append(m.Ext, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OfficeArtExtensionList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_OfficeArtExtensionList) Validate() error {
	return m.ValidateWithPath("CT_OfficeArtExtensionList")
}
func (m *CT_OfficeArtExtensionList) ValidateWithPath(path string) error {
	for i, v := range m.Ext {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ext[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}