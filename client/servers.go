package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Server -
type Server struct {
	ID   int `json:"id,omitempty"`
	Name int `json:"name,omitempty"`
}

// GetServer - Returns a specifc server
func (c *Client) GetServer(serverID string) (*Server, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/servers/%s", c.HostURL, serverID), nil)
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
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/servers", c.HostURL), strings.NewReader(name))
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
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/servers/%s", c.HostURL, serverID), strings.NewReader(name))
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
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/servers/%s", c.HostURL, serverID), nil)
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
