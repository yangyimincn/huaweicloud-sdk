package v3

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type IAMClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey string) *IAMClient {
	client := IAMClient{
		huaweicloud.NewClient(accessKey, secretKey, "", "iam"),
	}
	return &client
}
