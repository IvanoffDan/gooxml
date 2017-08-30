// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"log"
)

type CT_NonVisualConnectorProperties struct {
	CxnSpLocks *CT_ConnectorLocking
	StCxn      *CT_Connection
	EndCxn     *CT_Connection
	ExtLst     *CT_OfficeArtExtensionList
}

func NewCT_NonVisualConnectorProperties() *CT_NonVisualConnectorProperties {
	ret := &CT_NonVisualConnectorProperties{}
	return ret
}
func (m *CT_NonVisualConnectorProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.CxnSpLocks != nil {
		secxnSpLocks := xml.StartElement{Name: xml.Name{Local: "a:cxnSpLocks"}}
		e.EncodeElement(m.CxnSpLocks, secxnSpLocks)
	}
	if m.StCxn != nil {
		sestCxn := xml.StartElement{Name: xml.Name{Local: "a:stCxn"}}
		e.EncodeElement(m.StCxn, sestCxn)
	}
	if m.EndCxn != nil {
		seendCxn := xml.StartElement{Name: xml.Name{Local: "a:endCxn"}}
		e.EncodeElement(m.EndCxn, seendCxn)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NonVisualConnectorProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NonVisualConnectorProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cxnSpLocks":
				m.CxnSpLocks = NewCT_ConnectorLocking()
				if err := d.DecodeElement(m.CxnSpLocks, &el); err != nil {
					return err
				}
			case "stCxn":
				m.StCxn = NewCT_Connection()
				if err := d.DecodeElement(m.StCxn, &el); err != nil {
					return err
				}
			case "endCxn":
				m.EndCxn = NewCT_Connection()
				if err := d.DecodeElement(m.EndCxn, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualConnectorProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NonVisualConnectorProperties) Validate() error {
	return m.ValidateWithPath("CT_NonVisualConnectorProperties")
}
func (m *CT_NonVisualConnectorProperties) ValidateWithPath(path string) error {
	if m.CxnSpLocks != nil {
		if err := m.CxnSpLocks.ValidateWithPath(path + "/CxnSpLocks"); err != nil {
			return err
		}
	}
	if m.StCxn != nil {
		if err := m.StCxn.ValidateWithPath(path + "/StCxn"); err != nil {
			return err
		}
	}
	if m.EndCxn != nil {
		if err := m.EndCxn.ValidateWithPath(path + "/EndCxn"); err != nil {
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