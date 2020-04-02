package v1

import (
	"encoding/json"
	"fmt"
)

type CreateAppResponse struct {
	RequestID string `json:"request_id"`
	AppID     string `json:"app_id"`
}

func (c *CPHClient) CreateApp(name, packageName string) (*CreateAppResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps", c.GetProjectID())

	res := CreateAppResponse{}

	body := map[string]string{
		"name": name,
		"package_name": packageName,
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
