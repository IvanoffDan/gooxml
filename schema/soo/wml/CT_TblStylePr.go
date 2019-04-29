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
	"fmt"

	"github.com/IvanoffDan/gooxml"
)

type CT_TblStylePr struct {
	// Table Style Conditional Formatting Type
	TypeAttr ST_TblStyleOverrideType
	// Table Style Conditional Formatting Paragraph Properties
	PPr *CT_PPrGeneral
	// Table Style Conditional Formatting Run Properties
	RPr *CT_RPr
	// Table Style Conditional Formatting Table Properties
	TblPr *CT_TblPrBase
	// Table Style Conditional Formatting Table Row Properties
	TrPr *CT_TrPr
	// Table Style Conditional Formatting Table Cell Properties
	TcPr *CT_TcPr
}

func NewCT_TblStylePr() *CT_TblStylePr {
	ret := &CT_TblStylePr{}
	ret.TypeAttr = ST_TblStyleOverrideType(1)
	return ret
}

func (m *CT_TblStylePr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.PPr != nil {
		sepPr := xml.StartElement{Name: xml.Name{Local: "w:pPr"}}
		e.EncodeElement(m.PPr, sepPr)
	}
	if m.RPr != nil {
		serPr := xml.StartElement{Name: xml.Name{Local: "w:rPr"}}
		e.EncodeElement(m.RPr, serPr)
	}
	if m.TblPr != nil {
		setblPr := xml.StartElement{Name: xml.Name{Local: "w:tblPr"}}
		e.EncodeElement(m.TblPr, setblPr)
	}
	if m.TrPr != nil {
		setrPr := xml.StartElement{Name: xml.Name{Local: "w:trPr"}}
		e.EncodeElement(m.TrPr, setrPr)
	}
	if m.TcPr != nil {
		setcPr := xml.StartElement{Name: xml.Name{Local: "w:tcPr"}}
		e.EncodeElement(m.TcPr, setcPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblStylePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_TblStyleOverrideType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_TblStylePr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "pPr"}:
				m.PPr = NewCT_PPrGeneral()
				if err := d.DecodeElement(m.PPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPr"}:
				m.RPr = NewCT_RPr()
				if err := d.DecodeElement(m.RPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblPr"}:
				m.TblPr = NewCT_TblPrBase()
				if err := d.DecodeElement(m.TblPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "trPr"}:
				m.TrPr = NewCT_TrPr()
				if err := d.DecodeElement(m.TrPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tcPr"}:
				m.TcPr = NewCT_TcPr()
				if err := d.DecodeElement(m.TcPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TblStylePr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblStylePr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblStylePr and its children
func (m *CT_TblStylePr) Validate() error {
	return m.ValidateWithPath("CT_TblStylePr")
}

// ValidateWithPath validates the CT_TblStylePr and its children, prefixing error messages with path
func (m *CT_TblStylePr) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_TblStyleOverrideTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.PPr != nil {
		if err := m.PPr.ValidateWithPath(path + "/PPr"); err != nil {
			return err
		}
	}
	if m.RPr != nil {
		if err := m.RPr.ValidateWithPath(path + "/RPr"); err != nil {
			return err
		}
	}
	if m.TblPr != nil {
		if err := m.TblPr.ValidateWithPath(path + "/TblPr"); err != nil {
			return err
		}
	}
	if m.TrPr != nil {
		if err := m.TrPr.ValidateWithPath(path + "/TrPr"); err != nil {
			return err
		}
	}
	if m.TcPr != nil {
		if err := m.TcPr.ValidateWithPath(path + "/TcPr"); err != nil {
			return err
		}
	}
	return nil
}
