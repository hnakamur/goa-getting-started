//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/hnakamur/goa-getting-started/design
// --out=$(GOPATH)/src/github.com/hnakamur/goa-getting-started
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.myexample.bottle+json; view=default
type GoaMyexampleBottle struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GoaMyexampleBottle media type instance.
func (mt *GoaMyexampleBottle) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}
