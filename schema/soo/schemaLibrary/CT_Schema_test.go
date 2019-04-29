// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package schemaLibrary_test

import (
	"encoding/xml"
	"testing"

	"github.com/baliance/gooxml/schema/soo/schemaLibrary"
)

func TestCT_SchemaConstructor(t *testing.T) {
	v := schemaLibrary.NewCT_Schema()
	if v == nil {
		t.Errorf("schemaLibrary.NewCT_Schema must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed schemaLibrary.CT_Schema should validate: %s", err)
	}
}

func TestCT_SchemaMarshalUnmarshal(t *testing.T) {
	v := schemaLibrary.NewCT_Schema()
	buf, _ := xml.Marshal(v)
	v2 := schemaLibrary.NewCT_Schema()
	xml.Unmarshal(buf, v2)
}
