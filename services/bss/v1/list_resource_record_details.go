package v1

import (
	"encoding/json"
)

type ListResourceRecordDetails struct {
	MonthlyRecords []struct {
		Cycle                     string      `json:"cycle"`
		CloudServiceType          string      `json:"cloud_service_type"`
		Region                    string      `json:"region"`
		ResourceTypeCode          string      `json:"resource_Type_code"`
		ResInstanceID             string      `json:"res_instance_id"`
		ResourceName              string      `json:"resource_name"`
		ResourceTag               interface{} `json:"resource_tag"`
		ConsumeAmount             int         `json:"consume_amount"`
		CashAmount                int         `json:"cash_amount"`
		CreditAmount              int         `json:"credit_amount"`
		CouponAmount              int         `json:"coupon_amount"`
		FlexipurchaseCouponAmount int         `json:"flexipurchase_coupon_amount"`
		StoredCardAmount          int         `json:"stored_card_amount"`
		BonusAmount               int         `json:"bonus_amount"`
		DebtAmount                int         `json:"debt_amount"`
		AdjustmentAmount          interface{} `json:"adjustment_amount"`
		OfficialAmount            int         `json:"official_amount"`
		DiscountAmount            int         `json:"discount_amount"`
		MeasureID                 int         `json:"measure_id"`
		EnterpriseProjectID       interface{} `json:"enterprise_project_id"`
		EnterpriseProjectName     string      `json:"enterprise_project_name"`
		ChargeMode                int         `json:"charge_mode"`
		BillType                  int         `json:"bill_type"`
	} `json:"monthly_records"`
	TotalCount int    `json:"total_count"`
	Currency   string `json:"currency"`
}

func (c *BSSClient) ListResourceRecordDetails(cycle string, offset, limit int) (*ListResourceRecordDetails, error) {
	body := map[string]interface{}{
		"cycle": cycle,
		"include_zero_record": "false",
		"offset": offset,
		"limit": limit,
	}
	uri := "/v2/bills/customer-bills/res-records/query"
	res := ListResourceRecordDetails{}

	bodyB, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	response, err := c.DoRequest("POST", uri, nil, bodyB)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}