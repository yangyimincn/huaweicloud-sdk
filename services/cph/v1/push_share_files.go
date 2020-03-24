package v1

import (
	"encoding/json"
	"fmt"
)

type PushShareFilesResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		ServerID string `json:"server_id"`
		JobID    string `json:"job_id"`
	} `json:"jobs"`
}

func (c *CPHClient) PushShareFiles(bucketName, objectPath string, serverIDS []string) (*PushShareFilesResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/share-files", c.GetProjectID())
	res := PushShareFilesResponse{}

	body := map[string]interface{}{
		"bucket_name": bucketName,
		"object_path": objectPath,
		"server_ids": serverIDS,
	}

	bodyB, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	result, err := c.DoRequest("POST", uri, nil, bodyB)

	if err != nil {
		return &res, err
	}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}