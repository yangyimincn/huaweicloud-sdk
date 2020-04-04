package v1

import (
	"encoding/json"
	"fmt"
)

type QueryPhoneAppsDetailResponse struct {
	RequestID string `json:"request_id"`
	Apps      []struct {
		AppID          string `json:"app_id"`
		PackageName    string `json:"package_name"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		AppVerisonID   string `json:"app_verison_id"`
		VersionCode    string `json:"version_code"`
		VersionName    string `json:"version_name"`
		LaunchActivity string `json:"launch_activity"`
	} `json:"apps"`
}

func (c *CPHClient) QueryPhoneAppsDetail(phoneID string) (*QueryPhoneAppsDetailResponse, error) {
	uri := fmt.Sprintf("/v1/{project_id}/cloud-phone/phones/{phone_id}/apps", c.GetProjectID(), phoneID)
	res := QueryPhoneAppsDetailResponse{}

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
