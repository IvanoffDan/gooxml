// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/IvanoffDan/gooxml/schema/purl.org/dc/terms"
)

func TestElementsAndRefinementsGroupConstructor(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroup()
	if v == nil {
		t.Errorf("terms.NewElementsAndRefinementsGroup must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementsAndRefinementsGroup should validate: %s", err)
	}
}

func TestElementsAndRefinementsGroupMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroup()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementsAndRefinementsGroup()
	xml.Unmarshal(buf, v2)
}
