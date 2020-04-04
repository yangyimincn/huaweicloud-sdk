package v1

import (
	"encoding/json"
	"fmt"
)

type RestartCloudPhoneServerResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		PhoneID   string `json:"phone_id"`
		JobID     string `json:"job_id,omitempty"`
		ErrorCode string `json:"error_code,omitempty"`
		ErrorMsg  string `json:"error_msg,omitempty"`
	} `json:"jobs"`
}

func (c *CPHClient) RestartCloudPhonesServer(phones []map[string]string, imageID string) (*RestartCloudPhoneServerResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/servers/batch-restart", c.GetProjectID())
	res := RestartCloudPhoneServerResponse{}

	body := map[string]interface{}{
		"phones": phones,
	}

	if len(imageID) > 0 {
		body["image_id"] = imageID
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