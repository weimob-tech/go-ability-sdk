package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceInsert
// @id 1656
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1656?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品服务)
func (client *Client) GoodsServiceInsert(request *GoodsServiceInsertRequest) (response *GoodsServiceInsertResponse, err error) {
	rpcResponse := CreateGoodsServiceInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceInsertRequest struct {
	*api.BaseRequest

	BasicInfo   GoodsServiceInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	ServiceName string                             `json:"serviceName,omitempty"`
}

type GoodsServiceInsertResponse struct {
	ReturnResult bool  `json:"returnResult,omitempty"`
	Id           int64 `json:"id,omitempty"`
}

func CreateGoodsServiceInsertRequest() (request *GoodsServiceInsertRequest) {
	request = &GoodsServiceInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceInsertResponse() (response *api.BaseResponse[GoodsServiceInsertResponse]) {
	response = api.CreateResponse[GoodsServiceInsertResponse](&GoodsServiceInsertResponse{})
	return
}

type GoodsServiceInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
