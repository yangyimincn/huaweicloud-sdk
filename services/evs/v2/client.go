package v2

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type EVSClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey, region string) *EVSClient {
	client := EVSClient{
		huaweicloud.NewClient(accessKey, secretKey, region, "evs"),
	}
	return &client
}