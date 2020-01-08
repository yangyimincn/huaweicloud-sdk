package huaweicloud

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

func (h *HWClient) GetProjectID() string {
	if len(h.projectID) > 0 {
		return h.projectID
	} else {
		projectID, err := h.DescribeProjects(h.Region)
		if err != nil {
			log.Fatal("[fatal] Failed to get project id")
		}
		h.projectID = projectID
	}
	return h.projectID
}

func (h *HWClient) DescribeProjects(region string) (string, error) {
	query := map[string]string{
		"name": region,
	}

	var global bool = h.global
	var service string = h.Service
	h.global = true
	h.Service = "iam"

	res, err := h.DoRequest("GET", "/v3/projects", query, nil)
	h.global = global
	h.Service = service
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

