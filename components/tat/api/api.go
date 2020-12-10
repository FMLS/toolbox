package api

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type CommonReq struct {
	Action        string
	Region        string
	Uin           string
	SubAccountUin string
	AppID         int `json:"AppId"`
}

type StartSessionReq struct {
	CommonReq
	InstanceID string `json:"InstanceId"`
}

type StartSessionResp struct {
	Response struct {
		RequestID string `json:"RequestId"`
		SessionID string `json:"SessionId"`
		StreamURL string `json:"StreamUrl"`
	}
}

type API struct {
	Region        string
	AppID         int
	Uin           string
	SubAccountUin string
	url           string
}

func NewAPI(host string, port, appID int, region, uin string) *API {
	url := fmt.Sprintf("http://%s:%d", host, port)
	return &API{
		Region:        region,
		AppID:         appID,
		Uin:           uin,
		SubAccountUin: uin,
		url:           url,
	}
}

func (api *API) StartSession(instanceID string) (startSessionResp *StartSessionResp, err error) {
	client := resty.New()
	r, err := client.R().
		SetBody(StartSessionReq{
			CommonReq: CommonReq{
				Action:        "StartSession",
				Region:        api.Region,
				Uin:           api.Uin,
				SubAccountUin: api.Uin,
				AppID:         api.AppID,
			},
			InstanceID: instanceID,
		}).
		SetResult(&StartSessionResp{}).
		Post(api.url)
	log.Println(r.String())
	if err != nil {
		return nil, err
	}
	return r.Result().(*StartSessionResp), nil
}
