package v1

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type QueryAppDeployPhonesResponse struct {
	RequestID string `json:"request_id"`
	Phones    []struct {
		PhoneID        string `json:"phone_id"`
		AppVerisonID   string `json:"app_verison_id"`
		VersionCode    string `json:"version_code"`
		VersionName    string `json:"version_name"`
		LaunchActivity string `json:"launch_activity"`
	} `json:"phones"`
}

func (c *CPHClient) QueryAppDeployPhones(appID string, offset int) (*QueryAppDeployPhonesResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps/%s/phones", c.GetProjectID(), appID)
	res := QueryAppDeployPhonesResponse{}

	query := map[string]string{
		"limit": "100",
		"offset": strconv.Itoa(offset),
	}

	result, err := c.DoRequest("GET", uri, query, nil)
	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}