package v1

import (
	"encoding/json"
	"fmt"
)

type DescribePublicIPListResponse struct {
	PublicIPs []PublicIP `json:"publicips"`
}

type PublicIP struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	PublicIPAddress    string `json:"public_ip_address"`
	PrivateIPAddress   string `json:"private_ip_address"`
	PortID             string `json:"port_id"`
	BandwidthID        string `json:"bandwidth_id"`
	BandwidthShareType string `json:"bandwidth_share_type"`
	BandwidthSize      int    `json:"bandwidth_size"`
	BandwidthName      string `json:"bandwidth_name"`
}

func (v *VPCClient) DescribePublicIPList() (*DescribePublicIPListResponse, error) {
	uri := fmt.Sprintf("/v1/%s/publicips", v.GetProjectID())

	res := DescribePublicIPListResponse{}

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
