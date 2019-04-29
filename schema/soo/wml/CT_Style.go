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
	"github.com/IvanoffDan/gooxml/schema/soo/ofc/sharedTypes"
)

type CT_Style struct {
	// Style Type
	TypeAttr ST_StyleType
	// Style ID
	StyleIdAttr *string
	// Default Style
	DefaultAttr *sharedTypes.ST_OnOff
	// User-Defined Style
	CustomStyleAttr *sharedTypes.ST_OnOff
	// Primary Style Name
	Name *CT_String
	// Alternate Style Names
	Aliases *CT_String
	// Parent Style ID
	BasedOn *CT_String
	// Style For Next Paragraph
	Next *CT_String
	// Linked Style Reference
	Link *CT_String
	// Automatically Merge User Formatting Into Style Definition
	AutoRedefine *CT_OnOff
	// Hide Style From User Interface
	Hidden *CT_OnOff
	// Optional User Interface Sorting Order
	UiPriority *CT_DecimalNumber
	// Hide Style From Main User Interface
	SemiHidden *CT_OnOff
	// Remove Semi-Hidden Property When Style Is Used
	UnhideWhenUsed *CT_OnOff
	// Primary Style
	QFormat *CT_OnOff
	// Style Cannot Be Applied
	Locked *CT_OnOff
	// E-Mail Message Text Style
	Personal *CT_OnOff
	// E-Mail Message Composition Style
	PersonalCompose *CT_OnOff
	// E-Mail Message Reply Style
	PersonalReply *CT_OnOff
	// Revision Identifier for Style Definition
	Rsid *CT_LongHexNumber
	// Style Paragraph Properties
	PPr *CT_PPrGeneral
	// Run Properties
	RPr *CT_RPr
	// Style Table Properties
	TblPr *CT_TblPrBase
	// Style Table Row Properties
	TrPr *CT_TrPr
	// Style Table Cell Properties
	TcPr *CT_TcPr
	// Style Conditional Table Formatting Properties
	TblStylePr []*CT_TblStylePr
}

func NewCT_Style() *CT_Style {
	ret := &CT_Style{}
	return ret
}

func (m *CT_Style) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_StyleTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "w:type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StyleIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:styleId"},
			Value: fmt.Sprintf("%v", *m.StyleIdAttr)})
	}
	if m.DefaultAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:default"},
			Value: fmt.Sprintf("%v", *m.DefaultAttr)})
	}
	if m.CustomStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:customStyle"},
			Value: fmt.Sprintf("%v", *m.CustomStyleAttr)})
	}
	e.EncodeToken(start)
	if m.Name != nil {
		sename := xml.StartElement{Name: xml.Name{Local: "w:name"}}
		e.EncodeElement(m.Name, sename)
	}
	if m.Aliases != nil {
		sealiases := xml.StartElement{Name: xml.Name{Local: "w:aliases"}}
		e.EncodeElement(m.Aliases, sealiases)
	}
	if m.BasedOn != nil {
		sebasedOn := xml.StartElement{Name: xml.Name{Local: "w:basedOn"}}
		e.EncodeElement(m.BasedOn, sebasedOn)
	}
	if m.Next != nil {
		senext := xml.StartElement{Name: xml.Name{Local: "w:next"}}
		e.EncodeElement(m.Next, senext)
	}
	if m.Link != nil {
		selink := xml.StartElement{Name: xml.Name{Local: "w:link"}}
		e.EncodeElement(m.Link, selink)
	}
	if m.AutoRedefine != nil {
		seautoRedefine := xml.StartElement{Name: xml.Name{Local: "w:autoRedefine"}}
		e.EncodeElement(m.AutoRedefine, seautoRedefine)
	}
	if m.Hidden != nil {
		sehidden := xml.StartElement{Name: xml.Name{Local: "w:hidden"}}
		e.EncodeElement(m.Hidden, sehidden)
	}
	if m.UiPriority != nil {
		seuiPriority := xml.StartElement{Name: xml.Name{Local: "w:uiPriority"}}
		e.EncodeElement(m.UiPriority, seuiPriority)
	}
	if m.SemiHidden != nil {
		sesemiHidden := xml.StartElement{Name: xml.Name{Local: "w:semiHidden"}}
		e.EncodeElement(m.SemiHidden, sesemiHidden)
	}
	if m.UnhideWhenUsed != nil {
		seunhideWhenUsed := xml.StartElement{Name: xml.Name{Local: "w:unhideWhenUsed"}}
		e.EncodeElement(m.UnhideWhenUsed, seunhideWhenUsed)
	}
	if m.QFormat != nil {
		seqFormat := xml.StartElement{Name: xml.Name{Local: "w:qFormat"}}
		e.EncodeElement(m.QFormat, seqFormat)
	}
	if m.Locked != nil {
		selocked := xml.StartElement{Name: xml.Name{Local: "w:locked"}}
		e.EncodeElement(m.Locked, selocked)
	}
	if m.Personal != nil {
		sepersonal := xml.StartElement{Name: xml.Name{Local: "w:personal"}}
		e.EncodeElement(m.Personal, sepersonal)
	}
	if m.PersonalCompose != nil {
		sepersonalCompose := xml.StartElement{Name: xml.Name{Local: "w:personalCompose"}}
		e.EncodeElement(m.PersonalCompose, sepersonalCompose)
	}
	if m.PersonalReply != nil {
		sepersonalReply := xml.StartElement{Name: xml.Name{Local: "w:personalReply"}}
		e.EncodeElement(m.PersonalReply, sepersonalReply)
	}
	if m.Rsid != nil {
		sersid := xml.StartElement{Name: xml.Name{Local: "w:rsid"}}
		e.EncodeElement(m.Rsid, sersid)
	}
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
	if m.TblStylePr != nil {
		setblStylePr := xml.StartElement{Name: xml.Name{Local: "w:tblStylePr"}}
		for _, c := range m.TblStylePr {
			e.EncodeElement(c, setblStylePr)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Style) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "default" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DefaultAttr = &parsed
			continue
		}
		if attr.Name.Local == "customStyle" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.CustomStyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "styleId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleIdAttr = &parsed
			continue
		}
	}
lCT_Style:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "name"}:
				m.Name = NewCT_String()
				if err := d.DecodeElement(m.Name, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "aliases"}:
				m.Aliases = NewCT_String()
				if err := d.DecodeElement(m.Aliases, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "basedOn"}:
				m.BasedOn = NewCT_String()
				if err := d.DecodeElement(m.BasedOn, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "next"}:
				m.Next = NewCT_String()
				if err := d.DecodeElement(m.Next, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "link"}:
				m.Link = NewCT_String()
				if err := d.DecodeElement(m.Link, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "autoRedefine"}:
				m.AutoRedefine = NewCT_OnOff()
				if err := d.DecodeElement(m.AutoRedefine, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hidden"}:
				m.Hidden = NewCT_OnOff()
				if err := d.DecodeElement(m.Hidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "uiPriority"}:
				m.UiPriority = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.UiPriority, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "semiHidden"}:
				m.SemiHidden = NewCT_OnOff()
				if err := d.DecodeElement(m.SemiHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "unhideWhenUsed"}:
				m.UnhideWhenUsed = NewCT_OnOff()
				if err := d.DecodeElement(m.UnhideWhenUsed, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "qFormat"}:
				m.QFormat = NewCT_OnOff()
				if err := d.DecodeElement(m.QFormat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "locked"}:
				m.Locked = NewCT_OnOff()
				if err := d.DecodeElement(m.Locked, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "personal"}:
				m.Personal = NewCT_OnOff()
				if err := d.DecodeElement(m.Personal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "personalCompose"}:
				m.PersonalCompose = NewCT_OnOff()
				if err := d.DecodeElement(m.PersonalCompose, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "personalReply"}:
				m.PersonalReply = NewCT_OnOff()
				if err := d.DecodeElement(m.PersonalReply, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rsid"}:
				m.Rsid = NewCT_LongHexNumber()
				if err := d.DecodeElement(m.Rsid, &el); err != nil {
					return err
				}
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblStylePr"}:
				tmp := NewCT_TblStylePr()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblStylePr = append(m.TblStylePr, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_Style %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Style
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Style and its children
func (m *CT_Style) Validate() error {
	return m.ValidateWithPath("CT_Style")
}

// ValidateWithPath validates the CT_Style and its children, prefixing error messages with path
func (m *CT_Style) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if m.DefaultAttr != nil {
		if err := m.DefaultAttr.ValidateWithPath(path + "/DefaultAttr"); err != nil {
			return err
		}
	}
	if m.CustomStyleAttr != nil {
		if err := m.CustomStyleAttr.ValidateWithPath(path + "/CustomStyleAttr"); err != nil {
			return err
		}
	}
	if m.Name != nil {
		if err := m.Name.ValidateWithPath(path + "/Name"); err != nil {
			return err
		}
	}
	if m.Aliases != nil {
		if err := m.Aliases.ValidateWithPath(path + "/Aliases"); err != nil {
			return err
		}
	}
	if m.BasedOn != nil {
		if err := m.BasedOn.ValidateWithPath(path + "/BasedOn"); err != nil {
			return err
		}
	}
	if m.Next != nil {
		if err := m.Next.ValidateWithPath(path + "/Next"); err != nil {
			return err
		}
	}
	if m.Link != nil {
		if err := m.Link.ValidateWithPath(path + "/Link"); err != nil {
			return err
		}
	}
	if m.AutoRedefine != nil {
		if err := m.AutoRedefine.ValidateWithPath(path + "/AutoRedefine"); err != nil {
			return err
		}
	}
	if m.Hidden != nil {
		if err := m.Hidden.ValidateWithPath(path + "/Hidden"); err != nil {
			return err
		}
	}
	if m.UiPriority != nil {
		if err := m.UiPriority.ValidateWithPath(path + "/UiPriority"); err != nil {
			return err
		}
	}
	if m.SemiHidden != nil {
		if err := m.SemiHidden.ValidateWithPath(path + "/SemiHidden"); err != nil {
			return err
		}
	}
	if m.UnhideWhenUsed != nil {
		if err := m.UnhideWhenUsed.ValidateWithPath(path + "/UnhideWhenUsed"); err != nil {
			return err
		}
	}
	if m.QFormat != nil {
		if err := m.QFormat.ValidateWithPath(path + "/QFormat"); err != nil {
			return err
		}
	}
	if m.Locked != nil {
		if err := m.Locked.ValidateWithPath(path + "/Locked"); err != nil {
			return err
		}
	}
	if m.Personal != nil {
		if err := m.Personal.ValidateWithPath(path + "/Personal"); err != nil {
			return err
		}
	}
	if m.PersonalCompose != nil {
		if err := m.PersonalCompose.ValidateWithPath(path + "/PersonalCompose"); err != nil {
			return err
		}
	}
	if m.PersonalReply != nil {
		if err := m.PersonalReply.ValidateWithPath(path + "/PersonalReply"); err != nil {
			return err
		}
	}
	if m.Rsid != nil {
		if err := m.Rsid.ValidateWithPath(path + "/Rsid"); err != nil {
			return err
		}
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
	for i, v := range m.TblStylePr {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblStylePr[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
