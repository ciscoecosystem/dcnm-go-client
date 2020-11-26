package client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ciscoecosystem/dcnm-go-client/container"
	"github.com/ciscoecosystem/dcnm-go-client/models"
)

func (c *Client) GetviaURL(endpoint string) (*container.Container, error) {
	req, err := c.makeRequest("GET", endpoint, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if cont == nil {
		return nil, errors.New("Empty response body")
	}
	return cont, checkforerrors(cont, resp)
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
	return cont, checkforerrors(cont, resp)
}

func (c *Client) GetSegID(endpoint string) (*container.Container, error) {
	req, err := c.makeRequest("POST", endpoint, nil, true)
	if err != nil {
		return nil, err
	}

	cont, resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	return cont, checkforerrors(cont, resp)
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
	return cont, checkforerrors(cont, resp)
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
	return cont, checkforerrors(cont, resp)
}

func checkforerrors(cont *container.Container, resp *http.Response) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	} else if resp.StatusCode == 400 {
		return fmt.Errorf("%s Error : %s", resp.Status, cont.S("message").String())
	}
	return fmt.Errorf("%d Error : %s", resp.StatusCode, resp.Status)
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
