package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsBatchQueryTempGoodsByIdList
// @id 1260
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1260?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=批量查询临时商品信息)
func (client *Client) GoodsBatchQueryTempGoodsByIdList(request *GoodsBatchQueryTempGoodsByIdListRequest) (response *GoodsBatchQueryTempGoodsByIdListResponse, err error) {
	rpcResponse := CreateGoodsBatchQueryTempGoodsByIdListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsBatchQueryTempGoodsByIdListRequest struct {
	*api.BaseRequest

	Input GoodsBatchQueryTempGoodsByIdListRequestInput `json:"input,omitempty"`
}

type GoodsBatchQueryTempGoodsByIdListResponse struct {
	OutPut GoodsBatchQueryTempGoodsByIdListResponseOutPut `json:"outPut,omitempty"`
}

func CreateGoodsBatchQueryTempGoodsByIdListRequest() (request *GoodsBatchQueryTempGoodsByIdListRequest) {
	request = &GoodsBatchQueryTempGoodsByIdListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/batchQueryTempGoodsByIdList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsBatchQueryTempGoodsByIdListResponse() (response *api.BaseResponse[GoodsBatchQueryTempGoodsByIdListResponse]) {
	response = api.CreateResponse[GoodsBatchQueryTempGoodsByIdListResponse](&GoodsBatchQueryTempGoodsByIdListResponse{})
	return
}

type GoodsBatchQueryTempGoodsByIdListRequestInput struct {
}

type GoodsBatchQueryTempGoodsByIdListResponseOutPut struct {
}
