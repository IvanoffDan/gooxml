// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package powerpoint

import "github.com/baliance/gooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	gooxml.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Empty", NewCT_Empty)
	gooxml.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Rel", NewCT_Rel)
	gooxml.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "iscomment", NewIscomment)
	gooxml.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "textdata", NewTextdata)
}
