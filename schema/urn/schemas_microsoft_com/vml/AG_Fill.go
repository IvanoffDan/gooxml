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
	"fmt"

	"github.com/IvanoffDan/gooxml/schema/soo/ofc/sharedTypes"
)

type AG_Fill struct {
	FilledAttr    sharedTypes.ST_TrueFalse
	FillcolorAttr *string
}

func NewAG_Fill() *AG_Fill {
	ret := &AG_Fill{}
	return ret
}

func (m *AG_Fill) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FilledAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FilledAttr.MarshalXMLAttr(xml.Name{Local: "filled"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	return nil
}

func (m *AG_Fill) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "filled" {
			m.FilledAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Fill: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Fill and its children
func (m *AG_Fill) Validate() error {
	return m.ValidateWithPath("AG_Fill")
}

// ValidateWithPath validates the AG_Fill and its children, prefixing error messages with path
func (m *AG_Fill) ValidateWithPath(path string) error {
	if err := m.FilledAttr.ValidateWithPath(path + "/FilledAttr"); err != nil {
		return err
	}
	return nil
}
