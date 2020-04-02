package v1

import (
	"encoding/json"
	"fmt"
)

type UpdateAppVersionOfPhoneResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		JobType      int    `json:"job_type"`
		JobID        string `json:"job_id"`
		PhoneID      string `json:"phone_id"`
		AppVersionID string `json:"app_version_id"`
	} `json:"jobs"`
}

func (c *CPHClient) UpdateAppVersionOfPhone(appVersionIDS []string, phoneID string) (*UpdateAppVersionOfPhoneResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/%s/action", c.GetProjectID(), phoneID)
	res := UpdateAppVersionOfPhoneResponse{}

	body := map[string]interface{}{
		"update_app_version": map[string]interface{}{
			"app_version_ids": appVersionIDS,
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