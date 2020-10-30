package v2

import (
	"encoding/json"
	"fmt"
	"time"
)

type DescribeImageTagsResponse []struct {
	ID           int         `json:"id"`
	RepoID       int         `json:"repo_id"`
	Tag          string      `json:"Tag"`
	ImageID      string      `json:"image_id"`
	Manifest     string      `json:"manifest"`
	Digest       string      `json:"digest"`
	Schema       int         `json:"schema"`
	Path         string      `json:"path"`
	InternalPath string      `json:"internal_path"`
	Size         int         `json:"size"`
	IsTrusted    bool        `json:"is_trusted"`
	Created      time.Time   `json:"created"`
	Updated      time.Time   `json:"updated"`
	Deleted      interface{} `json:"deleted"`
}

func (v *SWRClient) DescribeVolumes(namespace, repository string) (*DescribeImageTagsResponse, error) {
	res := DescribeImageTagsResponse{}
	uri := fmt.Sprintf("/v2/manage/namespaces/%s/repos/%s/tags", namespace, repository)
	result, err := v.HWClient.DoRequest("GET", uri, nil, nil)

	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
