package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type ListAppsResponse struct {
	RequestID string `json:"request_id"`
	Apps      []struct {
		Name           string    `json:"name"`
		PackageName    string    `json:"package_name"`
		LaunchActivity string    `json:"launch_activity"`
		AppID          string    `json:"app_id"`
		Description    string    `json:"description"`
		CreateTime     time.Time `json:"create_time"`
		UpdateTime     time.Time `json:"update_time"`
	} `json:"apps"`
}

func (c *CPHClient) ListApps() (*ListAppsResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps", c.GetProjectID())

	res := ListAppsResponse{}


	result, err := c.DoRequest("GET", uri, nil, nil)
	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}