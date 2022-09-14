// Copyright 2019 James Cote
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by gen-esign; DO NOT EDIT.

// Package powerforms implements the DocuSign SDK
// category PowerForms.
//
// The PowerForms category enables you to create and manage PowerForms that you can use for self service and email forms.
//
// Service Api documentation may be found at:
// https://developers.docusign.com/esign-rest-api/reference/PowerForms
// Usage example:
//
//	import (
//	    "github.com/jfcote87/esign"
//	    "github.com/jfcote87/esign/v2.1/powerforms"
//	    "github.com/jfcote87/esign/v2.1/model"
//	)
//	...
//	powerformsService := powerforms.New(esignCredential)
package powerforms // import "github.com/jfcote87/esign/v2.1/powerforms"

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jfcote87/esign"
	"github.com/jfcote87/esign/v2.1/model"
)

// Service implements DocuSign PowerForms Category API operations
type Service struct {
	credential esign.Credential
}

// New initializes a powerforms service using cred to authorize ops.
func New(cred esign.Credential) *Service {
	return &Service{credential: cred}
}

// DataList returns the data that users entered in a PowerForm.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerformdata/list
//
// SDK Method PowerForms::getPowerFormData
func (s *Service) DataList(powerFormID string) *DataListOp {
	return &DataListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"powerforms", powerFormID, "form_data"}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// DataListOp implements DocuSign API SDK PowerForms::getPowerFormData
type DataListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DataListOp) Do(ctx context.Context) (*model.PowerFormsFormDataResponse, *esign.ResponseContext, error) {
	var res *model.PowerFormsFormDataResponse
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// DataLayout is the layout in which to return the PowerForm data. Valid values are:
//
// - `Native`
// - `Csv_Classic`
// - `Csv_One_Envelope_Per_Line`
// - `Xml_Classic`
func (op *DataListOp) DataLayout(val string) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("data_layout", val)
	}
	return op
}

// FromDate is the start date for a date range in UTC DateTime format.
//
// **Note**: If this property is null, no date filtering is applied.
func (op *DataListOp) FromDate(val time.Time) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// ToDate is the end date of a date range in UTC DateTime format. The default value is `UtcNow`.
func (op *DataListOp) ToDate(val time.Time) *DataListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// Create creates a new PowerForm.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/create
//
// SDK Method PowerForms::createPowerForm
func (s *Service) Create(powerForms *model.PowerForm) *CreateOp {
	return &CreateOp{
		Credential: s.credential,
		Method:     "POST",
		Path:       "powerforms",
		Payload:    powerForms,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// CreateOp implements DocuSign API SDK PowerForms::createPowerForm
type CreateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *CreateOp) Do(ctx context.Context) (*model.PowerForm, *esign.ResponseContext, error) {
	var res *model.PowerForm
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// Delete deletes a PowerForm.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/delete
//
// SDK Method PowerForms::deletePowerForm
func (s *Service) Delete(powerFormID string) *DeleteOp {
	return &DeleteOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// DeleteOp implements DocuSign API SDK PowerForms::deletePowerForm
type DeleteOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteOp) Do(ctx context.Context) (*esign.ResponseContext, error) {

	rspCtx, err := ((*esign.Op)(op)).Do(ctx, nil)
	return rspCtx, err
}

// DeleteList deletes one or more PowerForms.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/deletelist
//
// SDK Method PowerForms::deletePowerForms
func (s *Service) DeleteList(powerFormsRequest *model.PowerFormsRequest) *DeleteListOp {
	return &DeleteListOp{
		Credential: s.credential,
		Method:     "DELETE",
		Path:       "powerforms",
		Payload:    powerFormsRequest,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// DeleteListOp implements DocuSign API SDK PowerForms::deletePowerForms
type DeleteListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *DeleteListOp) Do(ctx context.Context) (*model.PowerFormsResponse, *esign.ResponseContext, error) {
	var res *model.PowerFormsResponse
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// Get returns a single PowerForm.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/get
//
// SDK Method PowerForms::getPowerForm
func (s *Service) Get(powerFormID string) *GetOp {
	return &GetOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// GetOp implements DocuSign API SDK PowerForms::getPowerForm
type GetOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *GetOp) Do(ctx context.Context) (*model.PowerForm, *esign.ResponseContext, error) {
	var res *model.PowerForm
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// List returns a list of PowerForms.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/list
//
// SDK Method PowerForms::listPowerForms
func (s *Service) List() *ListOp {
	return &ListOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "powerforms",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ListOp implements DocuSign API SDK PowerForms::listPowerForms
type ListOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListOp) Do(ctx context.Context) (*model.PowerFormsResponse, *esign.ResponseContext, error) {
	var res *model.PowerFormsResponse
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// FromDate (Optional) The start date for a date range.
//
// **Note**: If no value is provided, no date filtering is applied.
func (op *ListOp) FromDate(val time.Time) *ListOp {
	if op != nil {
		op.QueryOpts.Set("from_date", val.Format(time.RFC3339))
	}
	return op
}

// Order (Optional) The order in which to sort the results.
//
// Valid values are:
//
// * `asc`: Ascending order.
// * `desc`: Descending order.
func (op *ListOp) Order(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order", val)
	}
	return op
}

// OrderBy (Optional) The file attribute to use to sort the results.
//
// Valid values are:
//
// * `modified`
// * `name`
func (op *ListOp) OrderBy(val string) *ListOp {
	if op != nil {
		op.QueryOpts.Set("order_by", val)
	}
	return op
}

// ToDate (Optional) The end date for a date range.
//
// **Note**: If no value is provided, this property defaults to the current date.
func (op *ListOp) ToDate(val time.Time) *ListOp {
	if op != nil {
		op.QueryOpts.Set("to_date", val.Format(time.RFC3339))
	}
	return op
}

// ListSenders gets PowerForm senders.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/listsenders
//
// SDK Method PowerForms::listPowerFormSenders
func (s *Service) ListSenders() *ListSendersOp {
	return &ListSendersOp{
		Credential: s.credential,
		Method:     "GET",
		Path:       "powerforms/senders",
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// ListSendersOp implements DocuSign API SDK PowerForms::listPowerFormSenders
type ListSendersOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *ListSendersOp) Do(ctx context.Context) (*model.PowerFormSendersResponse, *esign.ResponseContext, error) {
	var res *model.PowerFormSendersResponse
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}

// StartPosition is the position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (op *ListSendersOp) StartPosition(val int) *ListSendersOp {
	if op != nil {
		op.QueryOpts.Set("start_position", fmt.Sprintf("%d", val))
	}
	return op
}

// Update updates an existing PowerForm.
//
// https://developers.docusign.com/esign-rest-api/reference/powerforms/powerforms/update
//
// SDK Method PowerForms::updatePowerForm
func (s *Service) Update(powerFormID string, powerForms *model.PowerForm) *UpdateOp {
	return &UpdateOp{
		Credential: s.credential,
		Method:     "PUT",
		Path:       strings.Join([]string{"powerforms", powerFormID}, "/"),
		Payload:    powerForms,
		QueryOpts:  make(url.Values),
		Version:    esign.VersionV21,
	}
}

// UpdateOp implements DocuSign API SDK PowerForms::updatePowerForm
type UpdateOp esign.Op

// Do executes the op.  A nil context will return error.
func (op *UpdateOp) Do(ctx context.Context) (*model.PowerForm, *esign.ResponseContext, error) {
	var res *model.PowerForm
	rspCtx, err := ((*esign.Op)(op)).Do(ctx, &res)
	return res, rspCtx, err
}
