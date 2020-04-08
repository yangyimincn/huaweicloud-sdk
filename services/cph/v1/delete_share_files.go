package v1

import (
	"encoding/json"
	"fmt"
)

type DeleteShareFilesResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		ServerID  string `json:"server_id"`
		JobID     string `json:"job_id,omitempty"`
		ErrorCode string `json:"error_code,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
	} `json:"jobs"`
}

func (c *CPHClient) DeleteShareFiles(filePath, serverIDS []string) (*DeleteShareFilesResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/share-files", c.GetProjectID())
	res := DeleteShareFilesResponse{}

	body := map[string]interface{}{
		"file_paths": filePath,
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
