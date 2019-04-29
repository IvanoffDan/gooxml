// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/IvanoffDan/gooxml/schema/soo/dml/chart"
)

func TestCT_TrendlineTypeConstructor(t *testing.T) {
	v := chart.NewCT_TrendlineType()
	if v == nil {
		t.Errorf("chart.NewCT_TrendlineType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_TrendlineType should validate: %s", err)
	}
}

func TestCT_TrendlineTypeMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_TrendlineType()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_TrendlineType()
	xml.Unmarshal(buf, v2)
}
