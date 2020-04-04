package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type ListJobsResponse struct {
	RequestID string `json:"request_id"`
	Jobs      []struct {
		PhoneID   string    `json:"phone_id"`
		ServerID  string    `json:"server_id"`
		NodeID    string    `json:"node_id"`
		JobID     string    `json:"job_id"`
		BeginTime time.Time `json:"begin_time"`
		EndTime   time.Time `json:"end_time"`
		Status    int       `json:"status"`
		ErrorCode string    `json:"error_code"`
		ErrorMsg  string    `json:"error_msg"`
	} `json:"jobs"`
}

func (c *CPHClient) ListJobs(requestID string) (*ListJobsResponse, error)  {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/phone-images", c.GetProjectID())

	res := ListJobsResponse{}

	query := map[string]string{
		"request_id": requestID,
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