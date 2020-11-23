package client

import (
	"errors"
	"net/http"

	"github.com/ciscoecosystem/dcnm-go-client/container"
	"github.com/ciscoecosystem/dcnm-go-client/models"
)

func (c *Client) GetviaURL(endpoint string) (*container.Container, error) {
	req, err := c.makeRequest("GET", endpoint, nil, true)
	if err != nil {
		return nil, err
	}

	obj, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, errors.New("Empty response body")
	}
	return obj, checkforerrors(resp)
}

func (c *Client) Save(endpoint string, obj models.Model) (*container.Container, error) {
	jsonPayload, err := c.prepareModel(obj)
	if err != nil {
		return nil, err
	}

	req, err := c.makeRequest("POST", endpoint, jsonPayload, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	return cont, checkforerrors(resp)
}

func (c *Client) Update(endpoint string, obj models.Model) (*container.Container, error) {
	jsonPayload, err := c.prepareModel(obj)
	if err != nil {
		return nil, err
	}

	req, err := c.makeRequest("PUT", endpoint, jsonPayload, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	return cont, checkforerrors(resp)
}

func (c *Client) Delete(endpoint string) (*container.Container, error) {
	req, err := c.makeRequest("DELETE", endpoint, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	return cont, checkforerrors(resp)
}

func checkforerrors(resp *http.Response) error {
	return nil
}

func (c *Client) prepareModel(obj models.Model) (*container.Container, error) {
	con, err := obj.ToMap()
	if err != nil {
		return nil, err
	}

	payload := &container.Container{}

	for key, value := range con {
		payload.Set(value, key)
	}
	return payload, nil
}
