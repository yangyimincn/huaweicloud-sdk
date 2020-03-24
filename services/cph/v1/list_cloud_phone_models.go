package v1

import (
	"encoding/json"
	"fmt"
)

type ListCloudPhoneModelsResponse struct {
	PhoneModels []struct {
		ServerModelName string  `json:"server_model_name"`
		PhoneModelName  string  `json:"phone_model_name"`
		CPU             float64 `json:"cpu"`
		Memory          int     `json:"memory"`
		Disk            int     `json:"disk"`
		SdDisk          int     `json:"sd_disk"`
		Resolution      string  `json:"resolution"`
		ExtendSpec      string  `json:"extend_spec"`
		SpecCode        string  `json:"spec_code"`
		OsType          string  `json:"os_type"`
		BusinessType    int     `json:"business_type"`
		PhoneCapacity   int     `json:"phone_capacity"`
		Status          int     `json:"status"`
	} `json:"phone_models"`
	RequestID string `json:"request_id"`
}

func (c *CPHClient) ListCloudPhoneModels() (*ListCloudPhoneModelsResponse, error)  {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phone-models", c.GetProjectID())

	res := ListCloudPhoneModelsResponse{}

	query := map[string]string{
		"status": "1",
	}

	result, err := c.DoRequest("GET", uri, query, nil)

	if err != nil {
		return &res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}