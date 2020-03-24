package v1

import (
	"encoding/json"
	"fmt"
	"time"
)

type QueryJobsResponse struct {
	ErrorMsg   interface{} `json:"error_msg"`
	ExecuteMsg interface{} `json:"execute_msg"`
	JobID      string      `json:"job_id"`
	EndTime    time.Time   `json:"end_time"`
	BeginTime  time.Time   `json:"begin_time"`
	ErrorCode  interface{} `json:"error_code"`
	RequestID  string      `json:"request_id"`
	Status     int         `json:"status"`
}

func (c *CPHClient) QueryJobs(jobID string) (*QueryJobsResponse, error) {
	uri := fmt.Sprintf("/v1/%s/cloud-phone/jobs/%s", c.GetProjectID(), jobID)

	res := QueryJobsResponse{}

	result, err := c.DoRequest("GET", uri, nil, nil)

	if err != nil {
		return &res, err
	}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return &res, err
	}
	return &res, nil
}
