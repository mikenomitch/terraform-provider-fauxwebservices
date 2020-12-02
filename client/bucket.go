package client

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/svanharmelen/jsonapi"
)

// Bucket -
type Bucket struct {
	ID   string `jsonapi:"primary,fake-resources"`
	Name string `jsonapi:"attr,name,omitempty"`
}

// GetBucket - Returns a specifc bucket
func (c *Client) GetBucket(bucketID string) (*Bucket, error) {
	s := new(Bucket)
	req, err := c.NewRequest("GET", fmt.Sprintf("api/v2/fake-resources/bucket/%s", bucketID), s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	bucket := new(Bucket)

	err = jsonapi.UnmarshalPayload(r, bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// CreateBucket - Create new bucket
func (c *Client) CreateBucket(s *Bucket) (*Bucket, error) {
	req, err := c.NewRequest("POST", "api/v2/fake-resources/bucket", s)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	bucket := new(Bucket)

	err = jsonapi.UnmarshalPayload(r, bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// UpdateBucket - Updates a bucket
func (c *Client) UpdateBucket(s *Bucket) (*Bucket, error) {
	req, err := c.NewRequest("PUT", fmt.Sprintf("api/v2/fake-resources/bucket/%s", s.ID), s)

	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(body)
	bucket := new(Bucket)

	err = jsonapi.UnmarshalPayload(r, bucket)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

// DeleteBucket - Deletes a bucket
func (c *Client) DeleteBucket(bucketID string) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("api/v2/fake-resources/bucket/%s", bucketID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted bucket" {
		return errors.New(string(body))
	}

	return nil
}
