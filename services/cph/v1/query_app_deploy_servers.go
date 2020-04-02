package v1

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type QueryAppDeployServersResponse struct {
	RequestID string `json:"request_id"`
	Servers   []struct {
		ServerID       string `json:"server_id"`
		AppVerisonID   string `json:"app_verison_id"`
		VersionCode    string `json:"version_code"`
		VersionName    string `json:"version_name"`
		LaunchActivity string `json:"launch_activity"`
	} `json:"servers"`
}

func (c *CPHClient) QueryAppDeployServers(appID string, offset int) (*QueryAppDeployServersResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/apps/%s/servers", c.GetProjectID(), appID)
	res := QueryAppDeployServersResponse{}

	query := map[string]string{
		"limit": "100",
		"offset": strconv.Itoa(offset),
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