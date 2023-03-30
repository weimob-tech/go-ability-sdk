package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsQueryRightsListByRightsOrderNos
// @id 1147
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1147?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据批量售后单号查询售后单列表)
func (client *Client) RightsQueryRightsListByRightsOrderNos(request *RightsQueryRightsListByRightsOrderNosRequest) (response *RightsQueryRightsListByRightsOrderNosResponse, err error) {
	rpcResponse := CreateRightsQueryRightsListByRightsOrderNosResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsQueryRightsListByRightsOrderNosRequest struct {
	*api.BaseRequest

	Obj RightsQueryRightsListByRightsOrderNosRequestObj `json:"obj,omitempty"`
}

type RightsQueryRightsListByRightsOrderNosResponse struct {
}

func CreateRightsQueryRightsListByRightsOrderNosRequest() (request *RightsQueryRightsListByRightsOrderNosRequest) {
	request = &RightsQueryRightsListByRightsOrderNosRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/queryRightsListByRightsOrderNos", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsQueryRightsListByRightsOrderNosResponse() (response *api.BaseResponse[RightsQueryRightsListByRightsOrderNosResponse]) {
	response = api.CreateResponse[RightsQueryRightsListByRightsOrderNosResponse](&RightsQueryRightsListByRightsOrderNosResponse{})
	return
}

type RightsQueryRightsListByRightsOrderNosRequestObj struct {
}
