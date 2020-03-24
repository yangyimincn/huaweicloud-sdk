package v1

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type ListEncodeServersResponse struct {
	EncodeServers []struct {
		EncodeServerName string `json:"encode_server_name"`
		EncodeServerID   string `json:"encode_server_id"`
		EncodeServerIP   string `json:"encode_server_ip"`
		ServerID         string `json:"server_id"`
		KeypairName      string `json:"keypair_name"`
		Type             int    `json:"type"`
		Status           int    `json:"status"`
		AccessInfos      []struct {
			ListenPort int    `json:"listen_port"`
			IntranetIP string `json:"intranet_ip"`
			AccessPort int    `json:"access_port"`
			AccessIP   string `json:"access_ip"`
			ServerIP   string `json:"server_ip"`
			PublicIP   string `json:"public_ip"`
		} `json:"access_infos"`
	} `json:"encode_servers"`
	RequestID string `json:"request_id"`
}

func (c *CPHClient) ListEncodeServers(offset, limit, encodeType,  status int, serverID string) (*ListEncodeServersResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/encode-servers", c.GetProjectID())

	res := ListEncodeServersResponse{}
	query := map[string]string{}
	if offset > 0 {
		query["offset"] = strconv.Itoa(offset)
	}
	if limit > 0 {
		query["limit"] = strconv.Itoa(limit)
	}
	if encodeType > 0 {
		query["type"] = strconv.Itoa(encodeType)
	}
	if status > 0 {
		query["status"] = strconv.Itoa(status)
	}
	if len(serverID) > 0 {
		query["server_id"] = serverID
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