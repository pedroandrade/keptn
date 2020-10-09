// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetEventsByTypeParams creates a new GetEventsByTypeParams object
// with the default values initialized.
func NewGetEventsByTypeParams() GetEventsByTypeParams {

	var (
		// initialize parameters with default values

		pageSizeDefault = int64(20)
	)

	return GetEventsByTypeParams{
		PageSize: &pageSizeDefault,
	}
}

// GetEventsByTypeParams contains all the bound params for the get events by type operation
// typically these are obtained from a http.Request
//
// swagger:parameters getEventsByType
type GetEventsByTypeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	EventType string
	/*
	  In: query
	*/
	ExcludeInvalidated *bool
	/*
	  Required: true
	  In: query
	*/
	Filter string
	/*
	  In: query
	*/
	FromTime *string
	/*Key of the page to be returned
	  In: query
	*/
	NextPageKey *string
	/*Page size to be returned
	  Maximum: 100
	  Minimum: 1
	  In: query
	  Default: 20
	*/
	PageSize *int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEventsByTypeParams() beforehand.
func (o *GetEventsByTypeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rEventType, rhkEventType, _ := route.Params.GetOK("eventType")
	if err := o.bindEventType(rEventType, rhkEventType, route.Formats); err != nil {
		res = append(res, err)
	}

	qExcludeInvalidated, qhkExcludeInvalidated, _ := qs.GetOK("excludeInvalidated")
	if err := o.bindExcludeInvalidated(qExcludeInvalidated, qhkExcludeInvalidated, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
		res = append(res, err)
	}

	qFromTime, qhkFromTime, _ := qs.GetOK("fromTime")
	if err := o.bindFromTime(qFromTime, qhkFromTime, route.Formats); err != nil {
		res = append(res, err)
	}

	qNextPageKey, qhkNextPageKey, _ := qs.GetOK("nextPageKey")
	if err := o.bindNextPageKey(qNextPageKey, qhkNextPageKey, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEventType binds and validates parameter EventType from path.
func (o *GetEventsByTypeParams) bindEventType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.EventType = raw

	return nil
}

// bindExcludeInvalidated binds and validates parameter ExcludeInvalidated from query.
func (o *GetEventsByTypeParams) bindExcludeInvalidated(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("excludeInvalidated", "query", "bool", raw)
	}
	o.ExcludeInvalidated = &value

	return nil
}

// bindFilter binds and validates parameter Filter from query.
func (o *GetEventsByTypeParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("filter", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("filter", "query", raw); err != nil {
		return err
	}

	o.Filter = raw

	return nil
}

// bindFromTime binds and validates parameter FromTime from query.
func (o *GetEventsByTypeParams) bindFromTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FromTime = &raw

	return nil
}

// bindNextPageKey binds and validates parameter NextPageKey from query.
func (o *GetEventsByTypeParams) bindNextPageKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.NextPageKey = &raw

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetEventsByTypeParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetEventsByTypeParams()
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	if err := o.validatePageSize(formats); err != nil {
		return err
	}

	return nil
}

// validatePageSize carries on validations for parameter PageSize
func (o *GetEventsByTypeParams) validatePageSize(formats strfmt.Registry) error {

	if err := validate.MinimumInt("pageSize", "query", int64(*o.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "query", int64(*o.PageSize), 100, false); err != nil {
		return err
	}

	return nil
}
