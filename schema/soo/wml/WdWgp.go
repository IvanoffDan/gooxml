// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"

	"github.com/IvanoffDan/gooxml"
	"github.com/IvanoffDan/gooxml/schema/soo/dml"
)

type WdWgp struct {
	WdCT_WordprocessingGroup
}

func NewWdWgp() *WdWgp {
	ret := &WdWgp{}
	ret.WdCT_WordprocessingGroup = *NewWdCT_WordprocessingGroup()
	return ret
}

func (m *WdWgp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.WdCT_WordprocessingGroup.MarshalXML(e, start)
}

func (m *WdWgp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.WdCT_WordprocessingGroup = *NewWdCT_WordprocessingGroup()
lWdWgp:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvPr"}:
				m.CNvPr = dml.NewCT_NonVisualDrawingProps()
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvGrpSpPr"}:
				if err := d.DecodeElement(m.CNvGrpSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "grpSpPr"}:
				if err := d.DecodeElement(m.GrpSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wsp"}:
				tmp := NewWdCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.Wsp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "grpSp"}:
				tmp := NewWdCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.GrpSp, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "graphicFrame"}:
				tmp := NewWdCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.GraphicFrame, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "pic"}:
				tmp := NewWdCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.Pic, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "contentPart"}:
				tmp := NewWdCT_WordprocessingGroupChoice()
				if err := d.DecodeElement(&tmp.ContentPart, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdWgp %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdWgp
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdWgp and its children
func (m *WdWgp) Validate() error {
	return m.ValidateWithPath("WdWgp")
}

// ValidateWithPath validates the WdWgp and its children, prefixing error messages with path
func (m *WdWgp) ValidateWithPath(path string) error {
	if err := m.WdCT_WordprocessingGroup.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
