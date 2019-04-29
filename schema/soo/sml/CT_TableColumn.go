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

type CT_TableColumn struct {
	// Table Field Id
	IdAttr uint32
	// Unique Name
	UniqueNameAttr *string
	// Column name
	NameAttr string
	// Totals Row Function
	TotalsRowFunctionAttr ST_TotalsRowFunction
	// Totals Row Label
	TotalsRowLabelAttr *string
	// Query Table Field Id
	QueryTableFieldIdAttr *uint32
	// Header Row Cell Format Id
	HeaderRowDxfIdAttr *uint32
	// Data & Insert Row Format Id
	DataDxfIdAttr *uint32
	// Totals Row Format Id
	TotalsRowDxfIdAttr *uint32
	// Header Row Cell Style
	HeaderRowCellStyleAttr *string
	// Data Area Style Name
	DataCellStyleAttr *string
	// Totals Row Style Name
	TotalsRowCellStyleAttr *string
	// Calculated Column Formula
	CalculatedColumnFormula *CT_TableFormula
	// Totals Row Formula
	TotalsRowFormula *CT_TableFormula
	// XML Column Properties
	XmlColumnPr *CT_XmlColumnPr
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_TableColumn() *CT_TableColumn {
	ret := &CT_TableColumn{}
	return ret
}

func (m *CT_TableColumn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.UniqueNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
			Value: fmt.Sprintf("%v", *m.UniqueNameAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.TotalsRowFunctionAttr != ST_TotalsRowFunctionUnset {
		attr, err := m.TotalsRowFunctionAttr.MarshalXMLAttr(xml.Name{Local: "totalsRowFunction"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TotalsRowLabelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowLabel"},
			Value: fmt.Sprintf("%v", *m.TotalsRowLabelAttr)})
	}
	if m.QueryTableFieldIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "queryTableFieldId"},
			Value: fmt.Sprintf("%v", *m.QueryTableFieldIdAttr)})
	}
	if m.HeaderRowDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowDxfId"},
			Value: fmt.Sprintf("%v", *m.HeaderRowDxfIdAttr)})
	}
	if m.DataDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataDxfId"},
			Value: fmt.Sprintf("%v", *m.DataDxfIdAttr)})
	}
	if m.TotalsRowDxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowDxfId"},
			Value: fmt.Sprintf("%v", *m.TotalsRowDxfIdAttr)})
	}
	if m.HeaderRowCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headerRowCellStyle"},
			Value: fmt.Sprintf("%v", *m.HeaderRowCellStyleAttr)})
	}
	if m.DataCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dataCellStyle"},
			Value: fmt.Sprintf("%v", *m.DataCellStyleAttr)})
	}
	if m.TotalsRowCellStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "totalsRowCellStyle"},
			Value: fmt.Sprintf("%v", *m.TotalsRowCellStyleAttr)})
	}
	e.EncodeToken(start)
	if m.CalculatedColumnFormula != nil {
		secalculatedColumnFormula := xml.StartElement{Name: xml.Name{Local: "ma:calculatedColumnFormula"}}
		e.EncodeElement(m.CalculatedColumnFormula, secalculatedColumnFormula)
	}
	if m.TotalsRowFormula != nil {
		setotalsRowFormula := xml.StartElement{Name: xml.Name{Local: "ma:totalsRowFormula"}}
		e.EncodeElement(m.TotalsRowFormula, setotalsRowFormula)
	}
	if m.XmlColumnPr != nil {
		sexmlColumnPr := xml.StartElement{Name: xml.Name{Local: "ma:xmlColumnPr"}}
		e.EncodeElement(m.XmlColumnPr, sexmlColumnPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableColumn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "totalsRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.TotalsRowDxfIdAttr = &pt
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "headerRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HeaderRowCellStyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "totalsRowLabel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TotalsRowLabelAttr = &parsed
			continue
		}
		if attr.Name.Local == "queryTableFieldId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.QueryTableFieldIdAttr = &pt
			continue
		}
		if attr.Name.Local == "headerRowDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HeaderRowDxfIdAttr = &pt
			continue
		}
		if attr.Name.Local == "dataDxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DataDxfIdAttr = &pt
			continue
		}
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "totalsRowFunction" {
			m.TotalsRowFunctionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dataCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DataCellStyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "totalsRowCellStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TotalsRowCellStyleAttr = &parsed
			continue
		}
	}
lCT_TableColumn:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "calculatedColumnFormula"}:
				m.CalculatedColumnFormula = NewCT_TableFormula()
				if err := d.DecodeElement(m.CalculatedColumnFormula, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "totalsRowFormula"}:
				m.TotalsRowFormula = NewCT_TableFormula()
				if err := d.DecodeElement(m.TotalsRowFormula, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "xmlColumnPr"}:
				m.XmlColumnPr = NewCT_XmlColumnPr()
				if err := d.DecodeElement(m.XmlColumnPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_TableColumn %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableColumn
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableColumn and its children
func (m *CT_TableColumn) Validate() error {
	return m.ValidateWithPath("CT_TableColumn")
}

// ValidateWithPath validates the CT_TableColumn and its children, prefixing error messages with path
func (m *CT_TableColumn) ValidateWithPath(path string) error {
	if err := m.TotalsRowFunctionAttr.ValidateWithPath(path + "/TotalsRowFunctionAttr"); err != nil {
		return err
	}
	if m.CalculatedColumnFormula != nil {
		if err := m.CalculatedColumnFormula.ValidateWithPath(path + "/CalculatedColumnFormula"); err != nil {
			return err
		}
	}
	if m.TotalsRowFormula != nil {
		if err := m.TotalsRowFormula.ValidateWithPath(path + "/TotalsRowFormula"); err != nil {
			return err
		}
	}
	if m.XmlColumnPr != nil {
		if err := m.XmlColumnPr.ValidateWithPath(path + "/XmlColumnPr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
