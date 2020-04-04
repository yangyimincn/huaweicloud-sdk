package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryAppVersionDetailResponse struct {
	RequestID    string    `json:"request_id"`
	AppID        string    `json:"app_id"`
	AppVersionID string    `json:"app_version_id"`
	VersionCode  string    `json:"version_code"`
	VersionName  string    `json:"version_name"`
	BucketName   string    `json:"bucket_name"`
	ObjectPath   string    `json:"object_path"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
}

func (c *CPHClient) QueryAppVersionDetail(appVersionID string) (*QueryAppVersionDetailResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/app-versions/%s", c.GetProjectID(), appVersionID)
	res := QueryAppVersionDetailResponse{}

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
