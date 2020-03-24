package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type ListCloudPhonesResponse struct {
	RequestID string   `json:"request_id"`
	Phones    []Phones `json:"phones"`
}

type Phones struct {
	PhoneName      string    `json:"phone_name"`
	ServerID       string    `json:"server_id"`
	PhoneID        string    `json:"phone_id"`
	PhoneModelName string    `json:"phone_model_name"`
	ImageID        string    `json:"image_id"`
	VncEnable      string    `json:"vnc_enable"`
	Status         int       `json:"status"`
	Type           int       `json:"type"`
	CreateTime     time.Time `json:"create_time"`
	UpdateTime     time.Time `json:"update_time"`
}

func (c *CPHClient) ListCloudPhones(params map[string]string) (*ListCloudPhonesResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones", c.GetProjectID())
	res := ListCloudPhonesResponse{}

	result, err := c.DoRequest("GET", uri, params, nil)

	if err != nil {
		return &res, err
	}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}