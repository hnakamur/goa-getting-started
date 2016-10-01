//************************************************************************//
// API "cellar": Application Contexts
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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	BottleID int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *GoaMyexampleBottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.myexample.bottle+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
