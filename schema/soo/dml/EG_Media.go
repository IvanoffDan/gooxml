// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/baliance/gooxml"
)

type EG_Media struct {
	AudioCd       *CT_AudioCD
	WavAudioFile  *CT_EmbeddedWAVAudioFile
	AudioFile     *CT_AudioFile
	VideoFile     *CT_VideoFile
	QuickTimeFile *CT_QuickTimeFile
}

func NewEG_Media() *EG_Media {
	ret := &EG_Media{}
	return ret
}

func (m *EG_Media) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "a:EG_Media"
	if m.AudioCd != nil {
		seaudioCd := xml.StartElement{Name: xml.Name{Local: "a:audioCd"}}
		e.EncodeElement(m.AudioCd, seaudioCd)
	}
	if m.WavAudioFile != nil {
		sewavAudioFile := xml.StartElement{Name: xml.Name{Local: "a:wavAudioFile"}}
		e.EncodeElement(m.WavAudioFile, sewavAudioFile)
	}
	if m.AudioFile != nil {
		seaudioFile := xml.StartElement{Name: xml.Name{Local: "a:audioFile"}}
		e.EncodeElement(m.AudioFile, seaudioFile)
	}
	if m.VideoFile != nil {
		sevideoFile := xml.StartElement{Name: xml.Name{Local: "a:videoFile"}}
		e.EncodeElement(m.VideoFile, sevideoFile)
	}
	if m.QuickTimeFile != nil {
		sequickTimeFile := xml.StartElement{Name: xml.Name{Local: "a:quickTimeFile"}}
		e.EncodeElement(m.QuickTimeFile, sequickTimeFile)
	}
	return nil
}

func (m *EG_Media) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Media:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "audioCd"}:
				m.AudioCd = NewCT_AudioCD()
				if err := d.DecodeElement(m.AudioCd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "wavAudioFile"}:
				m.WavAudioFile = NewCT_EmbeddedWAVAudioFile()
				if err := d.DecodeElement(m.WavAudioFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "audioFile"}:
				m.AudioFile = NewCT_AudioFile()
				if err := d.DecodeElement(m.AudioFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "videoFile"}:
				m.VideoFile = NewCT_VideoFile()
				if err := d.DecodeElement(m.VideoFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "quickTimeFile"}:
				m.QuickTimeFile = NewCT_QuickTimeFile()
				if err := d.DecodeElement(m.QuickTimeFile, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_Media %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Media
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Media and its children
func (m *EG_Media) Validate() error {
	return m.ValidateWithPath("EG_Media")
}

// ValidateWithPath validates the EG_Media and its children, prefixing error messages with path
func (m *EG_Media) ValidateWithPath(path string) error {
	if m.AudioCd != nil {
		if err := m.AudioCd.ValidateWithPath(path + "/AudioCd"); err != nil {
			return err
		}
	}
	if m.WavAudioFile != nil {
		if err := m.WavAudioFile.ValidateWithPath(path + "/WavAudioFile"); err != nil {
			return err
		}
	}
	if m.AudioFile != nil {
		if err := m.AudioFile.ValidateWithPath(path + "/AudioFile"); err != nil {
			return err
		}
	}
	if m.VideoFile != nil {
		if err := m.VideoFile.ValidateWithPath(path + "/VideoFile"); err != nil {
			return err
		}
	}
	if m.QuickTimeFile != nil {
		if err := m.QuickTimeFile.ValidateWithPath(path + "/QuickTimeFile"); err != nil {
			return err
		}
	}
	return nil
}
