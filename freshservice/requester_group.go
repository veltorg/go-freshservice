package freshservice

import (
	"context"
	"net/http"
	"net/url"
)

const requesterGroupURL = "/api/v2/requester_groups"

type RequesterGroupService interface {
	List(context.Context, QueryFilter) ([]RequesterGroupDetails, string, error)
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
