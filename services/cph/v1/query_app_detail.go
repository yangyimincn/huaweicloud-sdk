package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryAppDetailResponse struct {
	RequestID   string `json:"request_id"`
	PackageName string `json:"package_name"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AppVersions []struct {
		AppVerisonID   string    `json:"app_version_id"`
		VersionCode    string    `json:"version_code"`
		VersionName    string    `json:"version_name"`
		BucketName     string    `json:"bucket_name"`
		ObjectPath     string    `json:"object_path"`
		LaunchActivity string    `json:"launch_activity"`
		CreateTime     time.Time `json:"create_time"`
		UpdateTime     time.Time `json:"update_time"`
	} `json:"app_versions"`
}

func (c *CPHClient) QueryAppDetail(appID string) (*QueryAppDetailResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps/%s", c.GetProjectID(), appID)
	res := QueryAppDetailResponse{}

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
