// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/IvanoffDan/gooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestToConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewTo()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewTo must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.To should validate: %s", err)
	}
}

func TestToMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewTo()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewTo()
	xml.Unmarshal(buf, v2)
}
