package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsServiceContentInsert
// @id 1657
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1657?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品服务值)
func (client *Client) GoodsServiceContentInsert(request *GoodsServiceContentInsertRequest) (response *GoodsServiceContentInsertResponse, err error) {
	rpcResponse := CreateGoodsServiceContentInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsServiceContentInsertRequest struct {
	*api.BaseRequest

	BasicInfo        GoodsServiceContentInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	ServiceId        int64                                     `json:"serviceId,omitempty"`
	ServiceValueName string                                    `json:"serviceValueName,omitempty"`
}

type GoodsServiceContentInsertResponse struct {
	ReturnResult bool  `json:"returnResult,omitempty"`
	Id           int64 `json:"id,omitempty"`
}

func CreateGoodsServiceContentInsertRequest() (request *GoodsServiceContentInsertRequest) {
	request = &GoodsServiceContentInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/service/content/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsServiceContentInsertResponse() (response *api.BaseResponse[GoodsServiceContentInsertResponse]) {
	response = api.CreateResponse[GoodsServiceContentInsertResponse](&GoodsServiceContentInsertResponse{})
	return
}

type GoodsServiceContentInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
