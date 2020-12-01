package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Server -
type Server struct {
	ID   int `jsonapi:"id,omitempty"`
	Name int `jsonapi:"name,omitempty"`
}

// GetServer - Returns a specifc server
func (c *Client) GetServer(serverID string) (*Server, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("api/v2/fake-resources/server/%s", serverID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := Server{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

// CreateServer - Create new server
func (c *Client) CreateServer(name string) (*Server, error) {
	req, err := c.NewRequest("POST", "api/v2/fake-resources/server", strings.NewReader(name))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := Server{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

// UpdateServer - Updates a server
func (c *Client) UpdateServer(serverID string, name string) (*Server, error) {
	req, err := c.NewRequest("PUT", fmt.Sprintf("api/v2/fake-resources/server/%s", serverID), strings.NewReader(name))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	server := Server{}
	err = json.Unmarshal(body, &server)
	if err != nil {
		return nil, err
	}

	return &server, nil
}

// DeleteServer - Deletes a server
func (c *Client) DeleteServer(serverID string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("api/v2/fake-resources/server/%s", serverID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted server" {
		return errors.New(string(body))
	}

	return nil
}
