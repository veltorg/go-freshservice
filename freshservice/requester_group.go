package freshservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const requesterGroupURL = "/api/v2/requester_groups"

type RequesterGroupService interface {
	List(context.Context, QueryFilter) ([]RequesterGroupDetails, string, error)
	Create(context.Context, *RequesterGroupDetails) (*RequesterGroupDetails, error)
	Get(context.Context, int) (*RequesterGroupDetails, error)
	Update(context.Context, int, *RequesterGroupDetails) (*RequesterGroupDetails, error)
	Delete(context.Context, int) error
	AddRequesterToGroup(context.Context, int, int) error
	DeleteRequesterFromGroup(context.Context, int, int) error
	// TODO: Implement this method when the Requester type exists
	// ListRequesterGroupMembers(context.Context, int, QueryFilter) ([]Requester, string, error)
}

type RequesterGroupServiceClient struct {
	client *Client
}

// List all freshservice requester groups
func (as *RequesterGroupServiceClient) List(ctx context.Context, filter QueryFilter) ([]RequesterGroupDetails, string, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   requesterGroupURL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &RequesterGroups{}
	resp, err := as.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}

func (as *RequesterGroupServiceClient) Create(ctx context.Context, rg *RequesterGroupDetails) (*RequesterGroupDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   requesterGroupURL,
	}

	requesterGroupContent, err := json.Marshal(rg)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(requesterGroupContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &RequesterGroupDetails{}

	_, err = as.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *RequesterGroupServiceClient) Get(ctx context.Context, id int) (*RequesterGroupDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", requesterGroupURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &RequesterGroup{}

	_, err = as.client.makeRequest(req, res)

	if err != nil {
		return nil, err
	}

	return &res.Details, nil
}

func (as *RequesterGroupServiceClient) Update(ctx context.Context, id int, rg *RequesterGroupDetails) (*RequesterGroupDetails, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", requesterGroupURL, id),
	}

	requesterGroupContent, err := json.Marshal(rg)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(requesterGroupContent)

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url.String(), body)
	if err != nil {
		return nil, err
	}

	res := &RequesterGroup{}

	_, err = as.client.makeRequest(req, res)

	if err != nil {
		return nil, err
	}

	return &res.Details, nil
}

func (as *RequesterGroupServiceClient) Delete(ctx context.Context, id int) error {

	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/%d", requesterGroupURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)
	if err != nil {
		return err
	}

	_, err = as.client.makeRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}

func (as *RequesterGroupServiceClient) AddRequesterToGroup(ctx context.Context, groupID int, requesterID int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/members/%d", requesterGroupURL, groupID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url.String(), nil)

	if err != nil {
		return err
	}

	_, err = as.client.makeRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}

func (as *RequesterGroupServiceClient) DeleteRequesterFromGroup(ctx context.Context, groupID int, requesterID int) error {
	url := &url.URL{
		Scheme: "https",
		Host:   as.client.Domain,
		Path:   fmt.Sprintf("%s/members/%d", requesterGroupURL, groupID),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url.String(), nil)

	if err != nil {
		return err
	}

	_, err = as.client.makeRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}
