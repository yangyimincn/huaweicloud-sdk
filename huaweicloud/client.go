package huaweicloud

import (
	"bytes"
	"github.com/golang/glog"
	"github.com/yangyimincn/huaweicloud-sdk/huaweicloud/auth"
	hwerror "github.com/yangyimincn/huaweicloud-sdk/huaweicloud/error"
	"io/ioutil"
	"net/http"
	"time"
)

type HWClient struct {
	AccessKey  string
	SecretKey  string
	Region     string
	Service    string
	projectID  string
	global     bool

	httpClient http.Client
	config *Config
}

func (h *HWClient) generateHost() string {
	var url string
	// 全局服务endpoint不需要加region
	if h.global {
		url = "https://" + h.Service + "." + h.config.BaseHost
	} else {
		url = "https://" + h.Service + "." + h.Region + "." + h.config.BaseHost
	}
	return url
}

// do request
func (h *HWClient) DoRequest(method, uri string, query map[string]string, body []byte) ([]byte, error) {
	req, err := h.newRequest(method, uri, query, body)
	if err != nil {
		return nil, err
	}
	res, err := h.doReq(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// do req
func (h *HWClient) doReq(req *http.Request) ([]byte, error) {
	res, err := h.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 300 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			glog.Error("Failed to read response body: ", err)
			return nil, err
		}
		bodyStr := string(bodyBytes)
		return nil, hwerror.NewServerError(res.StatusCode, bodyStr)
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 生成auth headers
func (h *HWClient) generateAuthHeader(req *http.Request) (*http.Request, error) {
	nowTime := time.Unix(time.Now().Unix(), 0).UTC()
	x_sdk_date := nowTime.Format(auth.BasicDateFormat)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set(auth.HeaderXDate, x_sdk_date)

	s := auth.Newsigner(h.Service, h.Region, h.AccessKey, h.SecretKey)
	authkey, err := s.GetAuthorization(req)
	if err != nil {
		glog.Fatal("Get authorization error: ", err)
	}
	req.Header.Set(auth.HeaderAuthorization, authkey)
	return req, err
}

// 新建req
func (h *HWClient) newRequest(method, uri string, query map[string]string, body []byte) (*http.Request, error) {
	fullUrl := h.generateHost() + uri
	b := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest(method, fullUrl, b)
	if err != nil {
		return nil, err
	}

	// query string
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req, err = h.generateAuthHeader(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func NewClient(accessKey, secretKey, region, sevice string) *HWClient {
	cfg := NewConfig()
	client := HWClient{
		AccessKey:  accessKey,
		SecretKey:  secretKey,
		Region:     region,
		Service:    sevice,
		httpClient: http.Client{},
		config:     cfg,
	}
	client.httpClient.Timeout = cfg.Timeout
	return &client
}
