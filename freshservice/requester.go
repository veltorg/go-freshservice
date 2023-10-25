package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const RequesterURL = "/api/v2/requesters"

// RequesterService is an interface for interacting with
// the Requester endpoints of the Freshservice API
type RequesterService interface {
	List(context.Context, QueryFilter) (*Requesters, string, error)
	Create(context.Context, *RequesterDetails) (*Requester, error)
	Get(context.Context, int) (*Requester, error)
	Update(context.Context, int, *RequesterDetails) (*Requester, error)
	Delete(context.Context, int) (*ErrorResponse, error)
	Deactivate(context.Context, int) (*Requester, error)
	Reactivate(context.Context, int) (*Requester, error)
	ConvertToAgent(context.Context, int) (*Requester, error)
	MergeRequesters(context.Context, int, []int) (*Requester, error)
}

// RequesterServiceClient facilitates requests with the RequesterService methods
type RequesterServiceClient struct {
	client *Client
}

// List all freshservice Requesters
func (rs *RequesterServiceClient) List(ctx context.Context, filter QueryFilter) (*Requesters, string, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   RequesterURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &Requesters{}
	resp, err := rs.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res, HasNextPage(resp), nil
}

// Get a specific Freshservice Requester
func (rs *RequesterServiceClient) Get(ctx context.Context, id int) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d", RequesterURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Create a new Freshservice Requester
func (rs *RequesterServiceClient) Create(ctx context.Context, ad *RequesterDetails) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   RequesterURL,
	}

	RequesterContent, err := json.Marshal(ad)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(RequesterContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Update a Freshservice Requester
func (rs *RequesterServiceClient) Update(ctx context.Context, id int, ad *RequesterDetails) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d", RequesterURL, id),
	}

	RequesterContent, err := json.Marshal(ad)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(RequesterContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil

}

// Delete a Freshservice Requester
func (rs *RequesterServiceClient) Delete(ctx context.Context, id int) (*ErrorResponse, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d/forget", RequesterURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ErrorResponse{}

	_, err = rs.client.makeRequest(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Deactivate a Frehservice Requester (does not delete)
func (rs *RequesterServiceClient) Deactivate(ctx context.Context, id int) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d", RequesterURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Reactivate a Freshservice Requester
func (rs *RequesterServiceClient) Reactivate(ctx context.Context, id int) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d/reactivate", RequesterURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ConvertToAgent will convert a Freshservice Requester to an Agent
func (rs *RequesterServiceClient) ConvertToAgent(ctx context.Context, id int) (*Requester, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d/convert_to_agent", RequesterURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Requester{}
	if _, err := rs.client.makeRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// MergeRequesters will merge secondary requesters into a primary requester.
func (rs *RequesterServiceClient) MergeRequesters(ctx context.Context, id int, secondaryRequesterIDs []int) (*Requester, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   rs.client.Domain,
		Path:   fmt.Sprintf("%s/%d/merge", RequesterURL, id),
	}

	for _, secondaryRequesterID := range secondaryRequesterIDs {
		url.Query().Add("secondary_requesters", strconv.Itoa(secondaryRequesterID))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), nil)

	if err != nil {
		return nil, err
	}

	res := &Requester{}

	_, err = rs.client.makeRequest(req, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
