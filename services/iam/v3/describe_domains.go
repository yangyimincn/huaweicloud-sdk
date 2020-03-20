package v3

import (
	"encoding/json"
	"log"
)

type DescribeDomainsResponse struct {
	Domains []struct {
		Description string `json:"description"`
		Enabled     bool   `json:"enabled"`
		ID          string `json:"id"`
		Links       struct {
			Self string `json:"self"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"domains"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

func (c *IAMClient) DescribeDomains() (*DescribeDomainsResponse, error) {
	res, err := c.HWClient.DoRequest("GET", "/v3/auth/domains", nil, nil)
	if err != nil {
		log.Println("[warn] Failed to get project")
		return nil, err
	}

	response := DescribeDomainsResponse{}
	err = json.Unmarshal(res, &response)

	if err != nil {
		return &response, err
	}
	return &response, nil
}
