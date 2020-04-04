package v1

import (
	"encoding/json"
	"fmt"
)

type PowerOffCloudPhoneResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		PhoneID   string `json:"phone_id"`
		JobID     string `json:"job_id,omitempty"`
		ErrorCode string `json:"error_code,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
	} `json:"jobs"`
}

func (c *CPHClient) PowerOffCloudPhone(phoneIDS []string) (*ResetCloudPhoneResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/batch-stop", c.GetProjectID())
	res := ResetCloudPhoneResponse{}

	body := map[string]interface{}{
		"phone_ids": phoneIDS,
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