package v1

import (
	"encoding/json"
	"fmt"
)

type RestartEncodeServerResponse struct {
	Jobs []struct {
		JobID          string `json:"job_id,omitempty"`
		ErrorCode      string `json:"error_code,omitempty"`
		ErrorMsg       string `json:"error_msg,omitempty"`
		EncodeServerID string `json:"encode_server_id"`
	} `json:"jobs"`
	RequestID string `json:"request_id"`
}

func (c *CPHClient) RestartEncodeServer(encodeServerIDS []string) (*RestartEncodeServerResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/encode-servers/batch-restart", c.GetProjectID())
	res := RestartEncodeServerResponse{}

	body := map[string]interface{}{
		"encode_server_ids": encodeServerIDS,
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
