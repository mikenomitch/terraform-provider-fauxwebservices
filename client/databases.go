package client

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/svanharmelen/jsonapi"
)

// Database -
type Database struct {
	ID   string `jsonapi:"primary,fake-resources"`
	Name string `jsonapi:"attr,name,omitempty"`
}

// GetDatabase - Returns a specifc database
func (c *Client) GetDatabase(databaseID string) (*Database, error) {
	s := new(Database)
	req, err := c.NewRequest("GET", fmt.Sprintf("api/v2/fake-resources/database/%s", databaseID), s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	database := new(Database)

	err = jsonapi.UnmarshalPayload(r, database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

// CreateDatabase - Create new database
func (c *Client) CreateDatabase(s *Database) (*Database, error) {
	req, err := c.NewRequest("POST", "api/v2/fake-resources/database", s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	database := new(Database)

	err = jsonapi.UnmarshalPayload(r, database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

// UpdateDatabase - Updates a database
func (c *Client) UpdateDatabase(s *Database) (*Database, error) {
	req, err := c.NewRequest("PUT", fmt.Sprintf("api/v2/fake-resources/database/%s", s.ID), s)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	database := new(Database)

	err = jsonapi.UnmarshalPayload(r, database)
	if err != nil {
		return nil, err
	}

	return database, nil
}

// DeleteDatabase - Deletes a database
func (c *Client) DeleteDatabase(databaseID string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("api/v2/fake-resources/database/%s", databaseID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted database" {
		return errors.New(string(body))
	}

	return nil
}
