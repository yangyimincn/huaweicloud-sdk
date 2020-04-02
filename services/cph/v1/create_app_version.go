package v1

import (
	"encoding/json"
	"fmt"
)

type CreateAppVersionResponse struct {
	RequestID    string `json:"request_id"`
	AppVersionID string `json:"app_version_id"`
}

func (c *CPHClient) CreateAppVersion(appID, versionCode, versionName, launchActivity, bucketName, objectPath string) (*CreateAppVersionResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/app-versions", c.GetProjectID())

	res := CreateAppVersionResponse{}

	body := map[string]string{
		"app_id": appID,
		"version_code": versionCode,
		"version_name": versionName,
		"launch_activity": launchActivity,
		"bucket_name": bucketName,
		"object_path": objectPath,
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
