package sdp

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// WeikeImportCustomerRelation
// @id 1288
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=绑定客户-微客关系链)
func (client *Client) WeikeImportCustomerRelation(request *WeikeImportCustomerRelationRequest) (response *WeikeImportCustomerRelationResponse, err error) {
	rpcResponse := CreateWeikeImportCustomerRelationResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeikeImportCustomerRelationRequest struct {
	*api.BaseRequest

	WidList     []int `json:"widList,omitempty"`
	RelationWid int64 `json:"relationWid,omitempty"`
}

type WeikeImportCustomerRelationResponse struct {
}

func CreateWeikeImportCustomerRelationRequest() (request *WeikeImportCustomerRelationRequest) {
	request = &WeikeImportCustomerRelationRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("sdp", "1_0", "weike/importCustomerRelation", "api")
	request.Method = api.POST
	return
}

func CreateWeikeImportCustomerRelationResponse() (response *api.BaseResponse[WeikeImportCustomerRelationResponse]) {
	response = api.CreateResponse[WeikeImportCustomerRelationResponse](&WeikeImportCustomerRelationResponse{})
	return
}
