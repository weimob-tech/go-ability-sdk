package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeGetRelationship
// @id 741
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取客户的关系链)
func (client *Client) WeikeGetRelationship(request *WeikeGetRelationshipRequest) (response *WeikeGetRelationshipResponse, err error) {
	rpcResponse := CreateWeikeGetRelationshipResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeGetRelationshipRequest struct {
	*api.BaseRequest

	BizType int   `json:"bizType,omitempty"`
	BizId   int64 `json:"bizId,omitempty"`
	Wid     int64 `json:"wid,omitempty"`
}

type WeikeGetRelationshipResponse struct {
	SuperiorWeikeWid int64 `json:"superiorWeikeWid,omitempty"`
	BindWay          byte  `json:"bindWay,omitempty"`
	BindTime         int64 `json:"bindTime,omitempty"`
	State            byte  `json:"state,omitempty"`
}

func CreateWeikeGetRelationshipRequest() (request *WeikeGetRelationshipRequest) {
	request = &WeikeGetRelationshipRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/getRelationship", "api")
	request.Method = api.POST
	return
}

func CreateWeikeGetRelationshipResponse() (response *api.BaseResponse[WeikeGetRelationshipResponse]) {
	response = api.CreateResponse[WeikeGetRelationshipResponse](&WeikeGetRelationshipResponse{})
	return
}
