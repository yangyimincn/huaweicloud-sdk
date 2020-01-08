package v1

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type VPCClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey, region string) *VPCClient {
	client := VPCClient{
		huaweicloud.NewClient(accessKey, secretKey, region, "vpc"),
	}
	return &client
}
