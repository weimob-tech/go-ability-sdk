package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeImportWeikeRelation
// @id 1289
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改微客-邀请人关系链)
func (client *Client) WeikeImportWeikeRelation(request *WeikeImportWeikeRelationRequest) (response *WeikeImportWeikeRelationResponse, err error) {
	rpcResponse := CreateWeikeImportWeikeRelationResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeImportWeikeRelationRequest struct {
	*api.BaseRequest

	WidList     []int `json:"widList,omitempty"`
	RelationWid int64 `json:"relationWid,omitempty"`
}

type WeikeImportWeikeRelationResponse struct {
}

func CreateWeikeImportWeikeRelationRequest() (request *WeikeImportWeikeRelationRequest) {
	request = &WeikeImportWeikeRelationRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/importWeikeRelation", "api")
	request.Method = api.POST
	return
}

func CreateWeikeImportWeikeRelationResponse() (response *api.BaseResponse[WeikeImportWeikeRelationResponse]) {
	response = api.CreateResponse[WeikeImportWeikeRelationResponse](&WeikeImportWeikeRelationResponse{})
	return
}
