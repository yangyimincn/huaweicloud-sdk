package v1

import (
	"encoding/json"
	"fmt"
)

type DescribePortsListResponse struct {
	Ports	[]Port	`json:"ports"`
}

type Port struct {
	ID	string	`json:"id"`
	Name	string	`json:"name"`
	Status 	string	`json:"status"`
	DeviceID	string	`json:"device_id"`
	DeviceOwner	string	`json:"device_owner"`
}

func (v *VPCClient) DescribePortsList() (*DescribePortsListResponse, error) {
	uri := fmt.Sprintf("/v1/%s/ports", v.GetProjectID())
	res := DescribePortsListResponse{}
	result, err := v.DoRequest("GET", uri, nil, nil)

	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
