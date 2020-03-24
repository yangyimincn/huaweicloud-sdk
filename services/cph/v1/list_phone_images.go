package v1

import (
	"encoding/json"
	"fmt"
)

type ListPhoneImagesResonese struct {
	PhoneImages []struct {
		ImageName      string `json:"image_name"`
		OsType         string `json:"os_type"`
		PhoneModelName string `json:"phone_model_name"`
		OsName         string `json:"os_name"`
		ImageID        string `json:"image_id"`
		ImageType      int    `json:"image_type"`
	} `json:"phone_images"`
	RequestID string `json:"request_id"`
}

func (c *CPHClient) ListPhoneImages() (*ListPhoneImagesResonese, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phone-images", c.GetProjectID())

	res := ListPhoneImagesResonese{}

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
