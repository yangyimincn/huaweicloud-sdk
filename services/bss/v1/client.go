package v1

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type BSSClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey, region string) *BSSClient {
	client := BSSClient{
		huaweicloud.NewClient(accessKey, secretKey, region, "bss"),
	}
	client.Global = true
	return &client
}

