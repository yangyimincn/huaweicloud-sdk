package v2

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type DescribeVolumesResponse struct {
	Count   int      `json:"count"`
	Volumes []Volume `json:"volumes"`
}

type Volume struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Status           string       `json:"status"`
	Attachments      []Attachment `json:"attachments"`
	AvailabilityZone string       `json:"availability_zone"`
	SnapshotID       string       `json:"snapshot_id"`
	VolumeType       string       `json:"volume_type"`
	Size             int          `json:"size"`
	Bootable         string       `json:"bootable"`
	CreatedAt        string       `json:"created_at"`
	UpdatedAt        string       `json:"updated_at"`
}

type Attachment struct {
	ServerID     string `json:"server_id"`
	AttachmentID string `json:"attachment_id"`
	AttachedAt   string `json:"attached_at"`
	HostName     string `json:"host_name"`
	VolumeID     string `json:"volume_id"`
	Device       string `json:"device"`
}

func (v *EVSClient) DescribeVolumes(limit int, offset int, id string) (*DescribeVolumesResponse, error) {
	if limit == 0 {
		limit = 1000
	}
	query := map[string]string{}
	query["limit"] = strconv.Itoa(limit)
	query["offset"] = strconv.Itoa(offset)

	if len(id) > 0 {
		query["id"] = id
	}
	res := DescribeVolumesResponse{}
	uri := fmt.Sprintf("/v2/%s/cloudvolumes/detail", v.GetProjectID())
	result, err := v.HWClient.DoRequest("GET", uri, query, nil)

	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
