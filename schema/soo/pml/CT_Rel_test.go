// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/IvanoffDan/gooxml/schema/soo/pml"
)

func TestCT_RelConstructor(t *testing.T) {
	v := pml.NewCT_Rel()
	if v == nil {
		t.Errorf("pml.NewCT_Rel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_Rel should validate: %s", err)
	}
}

func TestCT_RelMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_Rel()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_Rel()
	xml.Unmarshal(buf, v2)
}
