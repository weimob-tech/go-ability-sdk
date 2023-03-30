package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyContentInsert
// @id 1630
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1630?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品属性值)
func (client *Client) GoodsPropertyContentInsert(request *GoodsPropertyContentInsertRequest) (response *GoodsPropertyContentInsertResponse, err error) {
	rpcResponse := CreateGoodsPropertyContentInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyContentInsertRequest struct {
	*api.BaseRequest

	BasicInfo     GoodsPropertyContentInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	PropId        int64                                      `json:"propId,omitempty"`
	PropValueName string                                     `json:"propValueName,omitempty"`
}

type GoodsPropertyContentInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsPropertyContentInsertRequest() (request *GoodsPropertyContentInsertRequest) {
	request = &GoodsPropertyContentInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/content/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyContentInsertResponse() (response *api.BaseResponse[GoodsPropertyContentInsertResponse]) {
	response = api.CreateResponse[GoodsPropertyContentInsertResponse](&GoodsPropertyContentInsertResponse{})
	return
}

type GoodsPropertyContentInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
