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
		ConsumeAmount             float64     `json:"consume_amount"`
		CashAmount                float64     `json:"cash_amount"`
		CreditAmount              float64     `json:"credit_amount"`
		CouponAmount              float64     `json:"coupon_amount"`
		FlexipurchaseCouponAmount float64     `json:"flexipurchase_coupon_amount"`
		StoredCardAmount          float64     `json:"stored_card_amount"`
		BonusAmount               float64     `json:"bonus_amount"`
		DebtAmount                float64     `json:"debt_amount"`
		AdjustmentAmount          interface{} `json:"adjustment_amount"`
		OfficialAmount            float64     `json:"official_amount"`
		DiscountAmount            float64     `json:"discount_amount"`
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
		"cycle":               cycle,
		"include_zero_record": "false",
		"offset":              offset,
		"limit":               limit,
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
