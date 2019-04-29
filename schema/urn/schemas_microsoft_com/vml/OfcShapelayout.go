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

	"github.com/IvanoffDan/gooxml"
)

type OfcShapelayout struct {
	OfcCT_ShapeLayout
}

func NewOfcShapelayout() *OfcShapelayout {
	ret := &OfcShapelayout{}
	ret.OfcCT_ShapeLayout = *NewOfcCT_ShapeLayout()
	return ret
}

func (m *OfcShapelayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:shapelayout"
	return m.OfcCT_ShapeLayout.MarshalXML(e, start)
}

func (m *OfcShapelayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_ShapeLayout = *NewOfcCT_ShapeLayout()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcShapelayout:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "idmap"}:
				m.Idmap = NewOfcCT_IdMap()
				if err := d.DecodeElement(m.Idmap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "regrouptable"}:
				m.Regrouptable = NewOfcCT_RegroupTable()
				if err := d.DecodeElement(m.Regrouptable, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "rules"}:
				m.Rules = NewOfcCT_Rules()
				if err := d.DecodeElement(m.Rules, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on OfcShapelayout %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcShapelayout
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcShapelayout and its children
func (m *OfcShapelayout) Validate() error {
	return m.ValidateWithPath("OfcShapelayout")
}

// ValidateWithPath validates the OfcShapelayout and its children, prefixing error messages with path
func (m *OfcShapelayout) ValidateWithPath(path string) error {
	if err := m.OfcCT_ShapeLayout.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
