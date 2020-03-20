package v1

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	iam "github.com/yangyimincn/huaweicloud-sdk/services/iam/v3"
)

type QueryBillSumResponse struct {
	ErrorCode        string  `json:"error_code"`
	ErrorMsg         string  `json:"error_msg"`
	Currency         string  `json:"currency"`
	TotalCount       int     `json:"total_count"`
	TotalAmount      int     `json:"total_amount"`
	DebtAmount       float64 `json:"debt_amount"`
	CouponAmount     float64 `json:"coupon_amount"`
	CashcouponAmount float64 `json:"cashcoupon_amount"`
	StoredcardAmount float64 `json:"storedcard_amount"`
	DebitAmount      float64 `json:"debit_amount"`
	CreditAmount     float64 `json:"credit_amount"`
	MeasureID        int     `json:"measure_id"`
	BillSums         []struct {
		CustomerID           string      `json:"customer_id"`
		ResourceTypeCode     string      `json:"resource_type_code"`
		RegionCode           interface{} `json:"region_code"`
		CloudServiceTypeCode string      `json:"cloud_service_type_code"`
		ConsumeTime          string      `json:"consume_time"`
		PayMethod            string      `json:"pay_method"`
		ConsumeAmount        float64         `json:"consume_amount"`
		Debt                 float64     `json:"debt"`
		Discount             float64         `json:"discount"`
		MeasureID            int         `json:"measure_id"`
		BillType             int         `json:"bill_type"`
		AccountDetails       []struct {
			BalanceTypeID string `json:"balance_type_id"`
			DeductAmount  float64    `json:"deduct_amount"`
		} `json:"account_details"`
		DiscountDetailInfos []struct {
			PromotionType  string      `json:"promotion_type"`
			DiscountAmount float64         `json:"discount_amount"`
			PromotionID    interface{} `json:"promotion_id"`
			MeasureID      int         `json:"measure_id"`
		} `json:"discount_detail_infos"`
		EnterpriseProjectID string `json:"enterpriseProjectId"`
	} `json:"bill_sums"`
}

func (c *BSSClient) QueryBillSum(cycle, cloudServiceTypeCode, accountType, enterpriseProjectID, domainID string) (*QueryBillSumResponse, error) {
	query := map[string]string{
		"cycle": cycle,
	}
	if len(cloudServiceTypeCode) > 0 {
		query["cloud_service_type_code"] = cloudServiceTypeCode
	}
	if len(accountType) > 0 {
		query["type"] = accountType
	}
	if len(enterpriseProjectID) > 0 {
		query["enterpriseProjectId"] = enterpriseProjectID
	}

	var accountID string
	if len(domainID) > 0 {
		accountID = domainID
	} else {
		iamClient := iam.NewClient(c.AccessKey, c.SecretKey)
		iamClient.Global = true
		accInfo, err := iamClient.DescribeDomains()
		if err != nil {
			glog.Error("Failed to query account info.")
			return nil, err
		}
		accountID = accInfo.Domains[0].ID
	}

	uri := fmt.Sprintf("/v1.0/%s/customer/account-mgr/bill/monthly-sum", accountID)

	res := QueryBillSumResponse{}
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
