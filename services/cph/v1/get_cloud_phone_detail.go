package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetCloudPhoneDetailResponse struct {
	RequestID      string `json:"request_id"`
	PhoneName      string `json:"phone_name"`
	ServerID       string `json:"server_id"`
	PhoneID        string `json:"phone_id"`
	ImageID        string `json:"image_id"`
	VncEnable      string `json:"vnc_enable"`
	PhoneModelName string `json:"phone_model_name"`
	Status         int    `json:"status"`
	AccessInfos    []struct {
		Type       string `json:"type"`
		DeviceIP   string `json:"device_ip"`
		PhoneIP    string `json:"phone_ip"`
		ListenPort int    `json:"listen_port"`
		AccessIP   string `json:"access_ip"`
		PublicIP   string `json:"public_ip"`
		IntranetIP string `json:"intranet_ip"`
		ServerIP   string `json:"server_ip"`
		AccessPort int    `json:"access_port"`
	} `json:"access_infos"`
	Property string `json:"property"`
	Metadata struct {
		OrderID   string `json:"order_id"`
		ProductID string `json:"product_id"`
	} `json:"metadata"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (c *CPHClient) GetCloudPhoneDetail(phoneID string) (*GetCloudPhoneDetailResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones/%s", c.GetProjectID(), phoneID)
	res := GetCloudPhoneDetailResponse{}

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