package v1

import (
	"encoding/json"
	"fmt"
)

type RunShellCommandResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		PhoneID string `json:"phone_id"`
		JobID   string `json:"job_id"`
	} `json:"jobs"`
}

func (c *CPHClient) RunShellCommand(content string, phoneIDS, serverIDS []string) (*RunShellCommandResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/commands", c.GetProjectID())

	res := RunShellCommandResponse{}

	body := map[string]interface{}{
		"command": "shell",
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
