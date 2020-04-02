package v1

import (
	"encoding/json"
	"fmt"
)


type QeuryServerAppsDetailResponse struct {
	RequestID string `json:"request_id"`
	Apps      []struct {
		AppID          string `json:"app_id"`
		AppVerisonID   string `json:"app_verison_id"`
		PackageName    string `json:"package_name"`
		Name           string `json:"name"`
		Description    string `json:"description"`
		VersionCode    string `json:"version_code"`
		VersionName    string `json:"version_name"`
		LaunchActivity string `json:"launch_activity"`
	} `json:"apps"`
}

func (c *CPHClient) QeuryServerAppsDetail(serverID string) (*QeuryServerAppsDetailResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/servers/%s/apps", c.GetProjectID(), serverID)
	res := QeuryServerAppsDetailResponse{}

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
