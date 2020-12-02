package client

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/svanharmelen/jsonapi"
)

// Server -
type Server struct {
	ID   string `jsonapi:"primary,fake-resources"`
	Name string `jsonapi:"attr,name,omitempty"`
}

// GetServer - Returns a specifc server
func (c *Client) GetServer(serverID string) (*Server, error) {
	s := new(Server)
	req, err := c.NewRequest("GET", fmt.Sprintf("api/v2/fake-resources/server/%s", serverID), s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	server := new(Server)

	err = jsonapi.UnmarshalPayload(r, server)
	if err != nil {
		return nil, err
	}

	return server, nil
}

// CreateServer - Create new server
func (c *Client) CreateServer(s *Server) (*Server, error) {
	req, err := c.NewRequest("POST", "api/v2/fake-resources/server", s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	server := new(Server)

	err = jsonapi.UnmarshalPayload(r, server)
	if err != nil {
		return nil, err
	}

	return server, nil
}

// UpdateServer - Updates a server
func (c *Client) UpdateServer(s *Server) (*Server, error) {
	req, err := c.NewRequest("PUT", fmt.Sprintf("api/v2/fake-resources/server/%s", s.ID), s)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	server := new(Server)

	err = jsonapi.UnmarshalPayload(r, server)
	if err != nil {
		return nil, err
	}

	return server, nil
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
