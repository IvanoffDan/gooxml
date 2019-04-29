// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chartDrawing

import (
	"encoding/xml"

	"github.com/IvanoffDan/gooxml"
)

type EG_Anchor struct {
	RelSizeAnchor *CT_RelSizeAnchor
	AbsSizeAnchor *CT_AbsSizeAnchor
}

func NewEG_Anchor() *EG_Anchor {
	ret := &EG_Anchor{}
	return ret
}

func (m *EG_Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RelSizeAnchor != nil {
		serelSizeAnchor := xml.StartElement{Name: xml.Name{Local: "relSizeAnchor"}}
		e.EncodeElement(m.RelSizeAnchor, serelSizeAnchor)
	}
	if m.AbsSizeAnchor != nil {
		seabsSizeAnchor := xml.StartElement{Name: xml.Name{Local: "absSizeAnchor"}}
		e.EncodeElement(m.AbsSizeAnchor, seabsSizeAnchor)
	}
	return nil
}

func (m *EG_Anchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Anchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "relSizeAnchor"}:
				m.RelSizeAnchor = NewCT_RelSizeAnchor()
				if err := d.DecodeElement(m.RelSizeAnchor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "absSizeAnchor"}:
				m.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if err := d.DecodeElement(m.AbsSizeAnchor, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_Anchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Anchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Anchor and its children
func (m *EG_Anchor) Validate() error {
	return m.ValidateWithPath("EG_Anchor")
}

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (m *EG_Anchor) ValidateWithPath(path string) error {
	if m.RelSizeAnchor != nil {
		if err := m.RelSizeAnchor.ValidateWithPath(path + "/RelSizeAnchor"); err != nil {
			return err
		}
	}
	if m.AbsSizeAnchor != nil {
		if err := m.AbsSizeAnchor.ValidateWithPath(path + "/AbsSizeAnchor"); err != nil {
			return err
		}
	}
	return nil
}
