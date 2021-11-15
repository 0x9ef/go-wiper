// Copyright (c) 2021 0x9ef. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.
package wipe

import "fmt"

// Policy describing general information about wiping rule with name,
// description and a raw Rule representation.
type Policy struct {
	Name        string
	Description string
	Rule        *Rule
}

func (p *Policy) String() string {
	if p != nil {
		return fmt.Sprintf("%s, \"%s\"\n%s", p.Name, p.Description, p.Rule.String())
	}
	return "<nil>"
}
