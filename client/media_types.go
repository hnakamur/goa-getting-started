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

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// ワインボトル (default view)
//
// Identifier: application/vnd.goa.myexample.bottle+json; view=default
type GoaMyexampleBottle struct {
	// このボトルにリクエストを送るためのAPIのhref
	Href string `form:"href" json:"href" xml:"href"`
	// 唯一なボトルID
	ID int `form:"id" json:"id" xml:"id"`
	// ワインの名前
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

// DecodeGoaMyexampleBottle decodes the GoaMyexampleBottle instance encoded in resp body.
func (c *Client) DecodeGoaMyexampleBottle(resp *http.Response) (*GoaMyexampleBottle, error) {
	var decoded GoaMyexampleBottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
