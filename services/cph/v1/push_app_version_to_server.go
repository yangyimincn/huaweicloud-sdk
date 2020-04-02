package v1

import (
	"encoding/json"
	"fmt"
)

type PushAppVersionToServerResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		JobType      int    `json:"job_type"`
		JobID        string `json:"job_id"`
		ServerID     string `json:"server_id"`
		AppVersionID string `json:"app_version_id"`
	} `json:"jobs"`
}

func (c *CPHClient) PushAppVersionToServer(appVersionIDS, serverIDS []string) (*PushAppVersionToServerResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/servers/action ", c.GetProjectID())
	res := PushAppVersionToServerResponse{}

	body := map[string]interface{}{
		"batch_push_app_version": map[string]interface{}{
			"app_version_ids": appVersionIDS,
			"server_ids": serverIDS,
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
