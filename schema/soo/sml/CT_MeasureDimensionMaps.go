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

type CT_MeasureDimensionMaps struct {
	// Measure Group Count
	CountAttr *uint32
	// OLAP Measure Group
	Map []*CT_MeasureDimensionMap
}

func NewCT_MeasureDimensionMaps() *CT_MeasureDimensionMaps {
	ret := &CT_MeasureDimensionMaps{}
	return ret
}

func (m *CT_MeasureDimensionMaps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.Map != nil {
		semap := xml.StartElement{Name: xml.Name{Local: "ma:map"}}
		for _, c := range m.Map {
			e.EncodeElement(c, semap)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MeasureDimensionMaps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_MeasureDimensionMaps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "map"}:
				tmp := NewCT_MeasureDimensionMap()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Map = append(m.Map, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_MeasureDimensionMaps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MeasureDimensionMaps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MeasureDimensionMaps and its children
func (m *CT_MeasureDimensionMaps) Validate() error {
	return m.ValidateWithPath("CT_MeasureDimensionMaps")
}

// ValidateWithPath validates the CT_MeasureDimensionMaps and its children, prefixing error messages with path
func (m *CT_MeasureDimensionMaps) ValidateWithPath(path string) error {
	for i, v := range m.Map {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Map[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
