package huaweicloud

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
)

type DescribeProjectsResponse struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID      string `json:"id"`
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

func (h *HWClient) GetProjectID() string {
	if len(h.projectID) > 0 {
		return h.projectID
	} else {
		projectID, err := h.DescribeProjects(h.Region)
		if err != nil {
			glog.Fatal("Failed to get project id: ", err)
		}
		h.projectID = projectID
	}
	return h.projectID
}

func (h *HWClient) DescribeProjects(region string) (string, error) {
	query := map[string]string{
		"name": region,
	}

	var global bool = h.Global
	var service string = h.Service
	h.Global = true
	h.Service = "iam"

	res, err := h.DoRequest("GET", "/v3/projects", query, nil)
	h.Global = global
	h.Service = service
	if err != nil {
		return "", err
	}

	response := DescribeProjectsResponse{}
	err = json.Unmarshal(res, &response)

	if err != nil {
		return "", err
	}

	// if project no exists
	if len(response.Projects) == 0 {
		return "", fmt.Errorf("Failed to get project_id")
	}

	return response.Projects[0].ID, nil
}

