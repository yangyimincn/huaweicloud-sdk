package v1

import (
	"encoding/json"
	"fmt"
)

type RemoveAppVersionFromServerResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		JobType      int    `json:"job_type"`
		JobID        string `json:"job_id,omitempty"`
		ErrorCode    string `json:"error_code,omitempty"`
		ErrorMsg     string `json:"error_msg,omitempty"`
		ServerID     string `json:"server_id"`
		AppVersionID string `json:"app_version_id"`
	} `json:"jobs"`
}

func (c *CPHClient) RemoveAppVersionFromServer(appVersionIDS, serverIDS []string) (*RemoveAppVersionFromServerResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/servers/action", c.GetProjectID())
	res := RemoveAppVersionFromServerResponse{}

	body := map[string]interface{}{
		"batch_remove_app_version": map[string]interface{}{
			"app_version_ids": appVersionIDS,
			"server_ids":      serverIDS,
		},
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
