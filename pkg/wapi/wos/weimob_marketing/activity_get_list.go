package weimob_marketing

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ActivityGetList
// @id 1762
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1762?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取活动列表)
func (client *Client) ActivityGetList(request *ActivityGetListRequest) (response *ActivityGetListResponse, err error) {
	rpcResponse := CreateActivityGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ActivityGetListRequest struct {
	*api.BaseRequest

	BosId       int64  `json:"bosId,omitempty"`
	CreatorVid  int64  `json:"creatorVid,omitempty"`
	CurVid      int64  `json:"curVid,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	KeyWord     string `json:"keyWord,omitempty"`
	PageIndex   int64  `json:"pageIndex,omitempty"`
	PageSize    int64  `json:"pageSize,omitempty"`
	QueryDetail bool   `json:"queryDetail,omitempty"`
	QueryVNode  bool   `json:"queryVNode,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	Status      int64  `json:"status,omitempty"`
	Type        int64  `json:"type,omitempty"`
	Vid         int64  `json:"vid,omitempty"`
}

type ActivityGetListResponse struct {
}

func CreateActivityGetListRequest() (request *ActivityGetListRequest) {
	request = &ActivityGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_marketing", "v2.0", "activity/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateActivityGetListResponse() (response *api.BaseResponse[ActivityGetListResponse]) {
	response = api.CreateResponse[ActivityGetListResponse](&ActivityGetListResponse{})
	return
}
