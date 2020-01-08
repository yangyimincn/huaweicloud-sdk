package v3

import (
	"encoding/json"
	"log"
)

type DescribeProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

func (c *v3.IAMClient) DescribeProjects(region string) (string, error) {
	query := map[string]string{
		"name": region,
	}

	res, err := c.HWClient.DoRequest("GET", "/v3/projects", query, nil)
	if err != nil {
		log.Println("[warn] Failed to get project")
		return "", err
	}

	response := DescribeProjectsResponse{}
	err = json.Unmarshal(res, &response)

	if err != nil {
		return "", err
	}
	return response.Projects[0].ID, nil
}
