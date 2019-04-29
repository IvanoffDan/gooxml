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
	"strconv"

	"github.com/baliance/gooxml"
)

type CT_RevisionMove struct {
	// Sheet Id
	SheetIdAttr uint32
	// Source
	SourceAttr string
	// Destination
	DestinationAttr string
	// Source Sheet Id
	SourceSheetIdAttr *uint32
	// Undo
	Undo []*CT_UndoInfo
	// Revision Cell Change
	Rcc []*CT_RevisionCellChange
	// Revision Format
	Rfmt    []*CT_RevisionFormatting
	RIdAttr *uint32
	UaAttr  *bool
	RaAttr  *bool
}

func NewCT_RevisionMove() *CT_RevisionMove {
	ret := &CT_RevisionMove{}
	return ret
}

func (m *CT_RevisionMove) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "source"},
		Value: fmt.Sprintf("%v", m.SourceAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "destination"},
		Value: fmt.Sprintf("%v", m.DestinationAttr)})
	if m.SourceSheetIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sourceSheetId"},
			Value: fmt.Sprintf("%v", *m.SourceSheetIdAttr)})
	}
	if m.RIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rId"},
			Value: fmt.Sprintf("%v", *m.RIdAttr)})
	}
	if m.UaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ua"},
			Value: fmt.Sprintf("%d", b2i(*m.UaAttr))})
	}
	if m.RaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ra"},
			Value: fmt.Sprintf("%d", b2i(*m.RaAttr))})
	}
	e.EncodeToken(start)
	if m.Undo != nil {
		seundo := xml.StartElement{Name: xml.Name{Local: "ma:undo"}}
		for _, c := range m.Undo {
			e.EncodeElement(c, seundo)
		}
	}
	if m.Rcc != nil {
		sercc := xml.StartElement{Name: xml.Name{Local: "ma:rcc"}}
		for _, c := range m.Rcc {
			e.EncodeElement(c, sercc)
		}
	}
	if m.Rfmt != nil {
		serfmt := xml.StartElement{Name: xml.Name{Local: "ma:rfmt"}}
		for _, c := range m.Rfmt {
			e.EncodeElement(c, serfmt)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_RevisionMove) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "source" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SourceAttr = parsed
			continue
		}
		if attr.Name.Local == "destination" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DestinationAttr = parsed
			continue
		}
		if attr.Name.Local == "sourceSheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SourceSheetIdAttr = &pt
			continue
		}
		if attr.Name.Local == "rId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RIdAttr = &pt
			continue
		}
		if attr.Name.Local == "ua" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UaAttr = &parsed
			continue
		}
		if attr.Name.Local == "ra" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RaAttr = &parsed
			continue
		}
	}
lCT_RevisionMove:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "undo"}:
				tmp := NewCT_UndoInfo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Undo = append(m.Undo, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rcc"}:
				tmp := NewCT_RevisionCellChange()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rcc = append(m.Rcc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "rfmt"}:
				tmp := NewCT_RevisionFormatting()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Rfmt = append(m.Rfmt, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_RevisionMove %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_RevisionMove
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_RevisionMove and its children
func (m *CT_RevisionMove) Validate() error {
	return m.ValidateWithPath("CT_RevisionMove")
}

// ValidateWithPath validates the CT_RevisionMove and its children, prefixing error messages with path
func (m *CT_RevisionMove) ValidateWithPath(path string) error {
	for i, v := range m.Undo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Undo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rcc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rcc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Rfmt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Rfmt[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
