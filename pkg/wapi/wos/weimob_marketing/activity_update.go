package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityUpdate
// @id 1760
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1760?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新活动)
func (client *Client) ActivityUpdate(request *ActivityUpdateRequest) (response *ActivityUpdateResponse, err error) {
	rpcResponse := CreateActivityUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityUpdateRequest struct {
	*api.BaseRequest

	VNodes       []ActivityUpdateRequestVNodes `json:"vNodes,omitempty"`
	Id           int64                         `json:"id,omitempty"`
	BosId        int64                         `json:"bosId,omitempty"`
	Vid          int64                         `json:"vid,omitempty"`
	UpdateTime   string                        `json:"updateTime,omitempty"`
	ActivityName string                        `json:"activityName,omitempty"`
	ActivityDesc string                        `json:"activityDesc,omitempty"`
	StartTime    string                        `json:"startTime,omitempty"`
	EndTime      string                        `json:"endTime,omitempty"`
}

type ActivityUpdateResponse struct {
	VNodes            []ActivityUpdateResponseVNodes `json:"vNodes,omitempty"`
	Id                int64                          `json:"id,omitempty"`
	BosId             int64                          `json:"bosId,omitempty"`
	ProductId         int64                          `json:"productId,omitempty"`
	ProductInstanceId int64                          `json:"productInstanceId,omitempty"`
	Vid               int64                          `json:"vid,omitempty"`
	VidType           int64                          `json:"vidType,omitempty"`
	ActivityId        int64                          `json:"activityId,omitempty"`
	ActivityTitle     string                         `json:"activityTitle,omitempty"`
	ActivityName      string                         `json:"activityName,omitempty"`
	ActivityPic       string                         `json:"activityPic,omitempty"`
	ActivityDesc      string                         `json:"activityDesc,omitempty"`
	StartTime         string                         `json:"startTime,omitempty"`
	EndTime           string                         `json:"endTime,omitempty"`
	MatchAllVid       bool                           `json:"matchAllVid,omitempty"`
	Status            int64                          `json:"status,omitempty"`
	CreateTime        string                         `json:"createTime,omitempty"`
	UpdateTime        string                         `json:"updateTime,omitempty"`
}

func CreateActivityUpdateRequest() (request *ActivityUpdateRequest) {
	request = &ActivityUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/update", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityUpdateResponse() (response *api.BaseResponse[ActivityUpdateResponse]) {
	response = api.CreateResponse[ActivityUpdateResponse](&ActivityUpdateResponse{})
	return
}

type ActivityUpdateRequestVNodes struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type ActivityUpdateResponseVNodes struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}
