package v1

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type CPHClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey, region string) *CPHClient {
	client := CPHClient{
		huaweicloud.NewClient(accessKey, secretKey, region, "cph"),
	}
	return &client
}
