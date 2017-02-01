// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "ferrariserver": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

// resourceLink user type.
type resourceLink struct {
	// Represents a link href
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// Represents the link http Method
	Method *string `form:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Represents a link rel
	Rel *string `form:"rel,omitempty" json:"rel,omitempty" xml:"rel,omitempty"`
}

// Publicize creates ResourceLink from resourceLink
func (ut *resourceLink) Publicize() *ResourceLink {
	var pub ResourceLink
	if ut.Href != nil {
		pub.Href = ut.Href
	}
	if ut.Method != nil {
		pub.Method = ut.Method
	}
	if ut.Rel != nil {
		pub.Rel = ut.Rel
	}
	return &pub
}

// ResourceLink user type.
type ResourceLink struct {
	// Represents a link href
	Href *string `form:"href,omitempty" json:"href,omitempty" xml:"href,omitempty"`
	// Represents the link http Method
	Method *string `form:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	// Represents a link rel
	Rel *string `form:"rel,omitempty" json:"rel,omitempty" xml:"rel,omitempty"`
}
