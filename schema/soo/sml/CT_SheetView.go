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

	"github.com/IvanoffDan/gooxml"
)

type CT_SheetView struct {
	// Window Protection
	WindowProtectionAttr *bool
	// Show Formulas
	ShowFormulasAttr *bool
	// Show Grid Lines
	ShowGridLinesAttr *bool
	// Show Headers
	ShowRowColHeadersAttr *bool
	// Show Zero Values
	ShowZerosAttr *bool
	// Right To Left
	RightToLeftAttr *bool
	// Sheet Tab Selected
	TabSelectedAttr *bool
	// Show Ruler
	ShowRulerAttr *bool
	// Show Outline Symbols
	ShowOutlineSymbolsAttr *bool
	// Default Grid Color
	DefaultGridColorAttr *bool
	// Show White Space
	ShowWhiteSpaceAttr *bool
	// View Type
	ViewAttr ST_SheetViewType
	// Top Left Visible Cell
	TopLeftCellAttr *string
	// Color Id
	ColorIdAttr *uint32
	// Zoom Scale
	ZoomScaleAttr *uint32
	// Zoom Scale Normal View
	ZoomScaleNormalAttr *uint32
	// Zoom Scale Page Break Preview
	ZoomScaleSheetLayoutViewAttr *uint32
	// Zoom Scale Page Layout View
	ZoomScalePageLayoutViewAttr *uint32
	// Workbook View Index
	WorkbookViewIdAttr uint32
	// View Pane
	Pane *CT_Pane
	// Selection
	Selection []*CT_Selection
	// PivotTable Selection
	PivotSelection []*CT_PivotSelection
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_SheetView() *CT_SheetView {
	ret := &CT_SheetView{}
	return ret
}

func (m *CT_SheetView) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WindowProtectionAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "windowProtection"},
			Value: fmt.Sprintf("%d", b2i(*m.WindowProtectionAttr))})
	}
	if m.ShowFormulasAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showFormulas"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowFormulasAttr))})
	}
	if m.ShowGridLinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showGridLines"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowGridLinesAttr))})
	}
	if m.ShowRowColHeadersAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRowColHeaders"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowRowColHeadersAttr))})
	}
	if m.ShowZerosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showZeros"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowZerosAttr))})
	}
	if m.RightToLeftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rightToLeft"},
			Value: fmt.Sprintf("%d", b2i(*m.RightToLeftAttr))})
	}
	if m.TabSelectedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tabSelected"},
			Value: fmt.Sprintf("%d", b2i(*m.TabSelectedAttr))})
	}
	if m.ShowRulerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showRuler"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowRulerAttr))})
	}
	if m.ShowOutlineSymbolsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showOutlineSymbols"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowOutlineSymbolsAttr))})
	}
	if m.DefaultGridColorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultGridColor"},
			Value: fmt.Sprintf("%d", b2i(*m.DefaultGridColorAttr))})
	}
	if m.ShowWhiteSpaceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showWhiteSpace"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowWhiteSpaceAttr))})
	}
	if m.ViewAttr != ST_SheetViewTypeUnset {
		attr, err := m.ViewAttr.MarshalXMLAttr(xml.Name{Local: "view"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TopLeftCellAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "topLeftCell"},
			Value: fmt.Sprintf("%v", *m.TopLeftCellAttr)})
	}
	if m.ColorIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "colorId"},
			Value: fmt.Sprintf("%v", *m.ColorIdAttr)})
	}
	if m.ZoomScaleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScale"},
			Value: fmt.Sprintf("%v", *m.ZoomScaleAttr)})
	}
	if m.ZoomScaleNormalAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScaleNormal"},
			Value: fmt.Sprintf("%v", *m.ZoomScaleNormalAttr)})
	}
	if m.ZoomScaleSheetLayoutViewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScaleSheetLayoutView"},
			Value: fmt.Sprintf("%v", *m.ZoomScaleSheetLayoutViewAttr)})
	}
	if m.ZoomScalePageLayoutViewAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomScalePageLayoutView"},
			Value: fmt.Sprintf("%v", *m.ZoomScalePageLayoutViewAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "workbookViewId"},
		Value: fmt.Sprintf("%v", m.WorkbookViewIdAttr)})
	e.EncodeToken(start)
	if m.Pane != nil {
		sepane := xml.StartElement{Name: xml.Name{Local: "ma:pane"}}
		e.EncodeElement(m.Pane, sepane)
	}
	if m.Selection != nil {
		seselection := xml.StartElement{Name: xml.Name{Local: "ma:selection"}}
		for _, c := range m.Selection {
			e.EncodeElement(c, seselection)
		}
	}
	if m.PivotSelection != nil {
		sepivotSelection := xml.StartElement{Name: xml.Name{Local: "ma:pivotSelection"}}
		for _, c := range m.PivotSelection {
			e.EncodeElement(c, sepivotSelection)
		}
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetView) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "view" {
			m.ViewAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "topLeftCell" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TopLeftCellAttr = &parsed
			continue
		}
		if attr.Name.Local == "showFormulas" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowFormulasAttr = &parsed
			continue
		}
		if attr.Name.Local == "colorId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ColorIdAttr = &pt
			continue
		}
		if attr.Name.Local == "showRowColHeaders" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRowColHeadersAttr = &parsed
			continue
		}
		if attr.Name.Local == "zoomScale" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScaleAttr = &pt
			continue
		}
		if attr.Name.Local == "rightToLeft" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RightToLeftAttr = &parsed
			continue
		}
		if attr.Name.Local == "zoomScaleNormal" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScaleNormalAttr = &pt
			continue
		}
		if attr.Name.Local == "showRuler" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowRulerAttr = &parsed
			continue
		}
		if attr.Name.Local == "zoomScaleSheetLayoutView" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScaleSheetLayoutViewAttr = &pt
			continue
		}
		if attr.Name.Local == "workbookViewId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.WorkbookViewIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "tabSelected" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TabSelectedAttr = &parsed
			continue
		}
		if attr.Name.Local == "zoomScalePageLayoutView" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.ZoomScalePageLayoutViewAttr = &pt
			continue
		}
		if attr.Name.Local == "showZeros" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowZerosAttr = &parsed
			continue
		}
		if attr.Name.Local == "windowProtection" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.WindowProtectionAttr = &parsed
			continue
		}
		if attr.Name.Local == "showOutlineSymbols" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowOutlineSymbolsAttr = &parsed
			continue
		}
		if attr.Name.Local == "showWhiteSpace" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowWhiteSpaceAttr = &parsed
			continue
		}
		if attr.Name.Local == "showGridLines" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowGridLinesAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultGridColor" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultGridColorAttr = &parsed
			continue
		}
	}
lCT_SheetView:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pane"}:
				m.Pane = NewCT_Pane()
				if err := d.DecodeElement(m.Pane, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "selection"}:
				tmp := NewCT_Selection()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Selection = append(m.Selection, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "pivotSelection"}:
				tmp := NewCT_PivotSelection()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.PivotSelection = append(m.PivotSelection, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_SheetView %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SheetView
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SheetView and its children
func (m *CT_SheetView) Validate() error {
	return m.ValidateWithPath("CT_SheetView")
}

// ValidateWithPath validates the CT_SheetView and its children, prefixing error messages with path
func (m *CT_SheetView) ValidateWithPath(path string) error {
	if err := m.ViewAttr.ValidateWithPath(path + "/ViewAttr"); err != nil {
		return err
	}
	if m.Pane != nil {
		if err := m.Pane.ValidateWithPath(path + "/Pane"); err != nil {
			return err
		}
	}
	for i, v := range m.Selection {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Selection[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.PivotSelection {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/PivotSelection[%d]", path, i)); err != nil {
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
