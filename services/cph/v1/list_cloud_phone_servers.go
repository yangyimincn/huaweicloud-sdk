package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type ListCloudPhoneServersResponse struct {
	RequestID string        `json:"request_id"`
	Servers   []PhoneServer `json:"servers"`
}

type PhoneServerAddresses struct {
	IntranetIP string `json:"intranet_ip"`
	AccessIP   string `json:"access_ip"`
	ServerIP   string `json:"server_ip"`
	PublicIP   string `json:"public_ip"`
}

type PhoneServer struct {
	ServerName        string      `json:"server_name"`
	ServerID          string      `json:"server_id"`
	ServerModelName   string      `json:"server_model_name"`
	PhoneModelName    string      `json:"phone_model_name"`
	KeypairName       string      `json:"keypair_name"`
	Status            int         `json:"status"`
	VpcID             string      `json:"vpc_id"`
	ResourceProjectID string      `json:"resource_project_id"`
	Cidr              string      `json:"cidr"`
	Addresses         []PhoneServerAddresses `json:"addresses"`
	CreateTime        time.Time   `json:"create_time"`
	UpdateTime        time.Time   `json:"update_time"`
}

func (c *CPHClient) ListCloudPhoneServers() (*ListCloudPhoneServersResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/servers", c.GetProjectID())

	res := ListCloudPhoneServersResponse{}

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