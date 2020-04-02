package v1

import (
	"encoding/json"
	"fmt"
)

type DeleteAppResponse struct {
	RequestID string `json:"request_id"`
}

func (c *CPHClient) DeleteApp(appID string) (*DeleteAppResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps/%s", c.GetProjectID(), appID)

	res := DeleteAppResponse{}

	result, err := c.DoRequest("DELETE", uri, nil, nil)
	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
