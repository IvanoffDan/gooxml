// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/IvanoffDan/gooxml/schema/urn/schemas_microsoft_com/vml"
)

func TestTextboxConstructor(t *testing.T) {
	v := vml.NewTextbox()
	if v == nil {
		t.Errorf("vml.NewTextbox must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Textbox should validate: %s", err)
	}
}

func TestTextboxMarshalUnmarshal(t *testing.T) {
	v := vml.NewTextbox()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewTextbox()
	xml.Unmarshal(buf, v2)
}
