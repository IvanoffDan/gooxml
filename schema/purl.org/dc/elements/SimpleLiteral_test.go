// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package elements_test

import (
	"encoding/xml"
	"testing"

	"github.com/baliance/gooxml/schema/purl.org/dc/elements"
)

func TestSimpleLiteralConstructor(t *testing.T) {
	v := elements.NewSimpleLiteral()
	if v == nil {
		t.Errorf("elements.NewSimpleLiteral must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed elements.SimpleLiteral should validate: %s", err)
	}
}

func TestSimpleLiteralMarshalUnmarshal(t *testing.T) {
	v := elements.NewSimpleLiteral()
	buf, _ := xml.Marshal(v)
	v2 := elements.NewSimpleLiteral()
	xml.Unmarshal(buf, v2)
}
