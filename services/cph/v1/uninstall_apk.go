package v1

import (
	"encoding/json"
	"fmt"
)

type UninstallApkResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		PhoneID   string `json:"phone_id"`
		JobID     string `json:"job_id,omitempty"`
		ErrorCode string `json:"error_code,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
	} `json:"jobs"`
}

func (c *CPHClient) UninstallApk(content string, phoneIDS, serverIDS []string) (*UninstallApkResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/commands", c.GetProjectID())

	res := UninstallApkResponse{}

	body := map[string]interface{}{
		"command": "uninstall",
		"content": content,
	}

	if len(phoneIDS) > 0 {
		body["phone_ids"] = phoneIDS
	} else {
		body["server_ids"] = serverIDS
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
