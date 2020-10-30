package v2

import "github.com/yangyimincn/huaweicloud-sdk/huaweicloud"

type SWRClient struct {
	*huaweicloud.HWClient
}

func NewClient(accessKey, secretKey, region string) *SWRClient {
	client := SWRClient{
		huaweicloud.NewClient(accessKey, secretKey, region, "swr-api"),
	}
	return &client
}