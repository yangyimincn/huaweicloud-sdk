package v1

import (
	"encoding/json"
	"fmt"
)

type CreateCloudPhoneServerResponse struct {
	RequestID string `json:"request_id"`
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id"`
}

func (c *CPHClient) CreateCloudPhoneServer(body map[string]interface{}) (*CreateCloudPhoneServerResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phones", c.GetProjectID())
	res := CreateCloudPhoneServerResponse{}

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