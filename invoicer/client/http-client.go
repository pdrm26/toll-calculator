package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pdrm26/toll-calculator/types"
)

type Client struct {
	Endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

func (c *Client) AggregateDistance(distance types.Distance) error {
	distBytes, err := json.Marshal(distance)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(distBytes))
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error status %d", res.StatusCode)
	}

	return nil
}
