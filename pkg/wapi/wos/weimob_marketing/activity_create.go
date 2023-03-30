package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityCreate
// @id 1766
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1766?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建活动)
func (client *Client) ActivityCreate(request *ActivityCreateRequest) (response *ActivityCreateResponse, err error) {
	rpcResponse := CreateActivityCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityCreateRequest struct {
	*api.BaseRequest

	VNodes            []ActivityCreateRequestVNodes `json:"vNodes,omitempty"`
	BosId             int64                         `json:"bosId,omitempty"`
	ProductId         int64                         `json:"productId,omitempty"`
	ProductInstanceId int64                         `json:"productInstanceId,omitempty"`
	Vid               int64                         `json:"vid,omitempty"`
	VidType           int64                         `json:"vidType,omitempty"`
	ActivityId        int64                         `json:"activityId,omitempty"`
	ActivityName      string                        `json:"activityName,omitempty"`
	ActivityPic       string                        `json:"activityPic,omitempty"`
	ActivityDesc      string                        `json:"activityDesc,omitempty"`
	StartTime         string                        `json:"startTime,omitempty"`
	EndTime           string                        `json:"endTime,omitempty"`
	MatchAllVid       bool                          `json:"matchAllVid,omitempty"`
	Status            int64                         `json:"status,omitempty"`
}

type ActivityCreateResponse struct {
	VNodes            []ActivityCreateResponseVNodes `json:"vNodes,omitempty"`
	Id                string                         `json:"id,omitempty"`
	BosId             string                         `json:"bosId,omitempty"`
	ProductId         string                         `json:"productId,omitempty"`
	ProductInstanceId string                         `json:"productInstanceId,omitempty"`
	Vid               string                         `json:"vid,omitempty"`
	VidType           string                         `json:"vidType,omitempty"`
	ActivityId        string                         `json:"activityId,omitempty"`
	ActivityTitle     string                         `json:"activityTitle,omitempty"`
	ActivityName      string                         `json:"activityName,omitempty"`
	ActivityPic       string                         `json:"activityPic,omitempty"`
	ActivityDesc      string                         `json:"activityDesc,omitempty"`
	StartTime         string                         `json:"startTime,omitempty"`
	EndTime           string                         `json:"endTime,omitempty"`
	MatchAllVid       string                         `json:"matchAllVid,omitempty"`
	Status            string                         `json:"status,omitempty"`
	CreateTime        string                         `json:"createTime,omitempty"`
	UpdateTime        string                         `json:"updateTime,omitempty"`
}

func CreateActivityCreateRequest() (request *ActivityCreateRequest) {
	request = &ActivityCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/create", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityCreateResponse() (response *api.BaseResponse[ActivityCreateResponse]) {
	response = api.CreateResponse[ActivityCreateResponse](&ActivityCreateResponse{})
	return
}

type ActivityCreateRequestVNodes struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type ActivityCreateResponseVNodes struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
