package v1

import (
	"encoding/json"
	"fmt"
)

type DeleteAppVersionResponse struct {
	RequestID string `json:"request_id"`
}

func (c *CPHClient) DeleteAppVersion(appVersionID string) (*DeleteAppVersionResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/app-versions/%s", c.GetProjectID(), appVersionID)

	res := DeleteAppVersionResponse{}

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

