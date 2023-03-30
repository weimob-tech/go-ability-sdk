package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyInsert
// @id 1629
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1629?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=添加商品属性)
func (client *Client) GoodsPropertyInsert(request *GoodsPropertyInsertRequest) (response *GoodsPropertyInsertResponse, err error) {
	rpcResponse := CreateGoodsPropertyInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyInsertRequest struct {
	*api.BaseRequest

	BasicInfo    GoodsPropertyInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	PropName     string                              `json:"propName,omitempty"`
	Sort         int64                               `json:"sort,omitempty"`
	IsSearchable bool                                `json:"isSearchable,omitempty"`
	PropType     int64                               `json:"propType,omitempty"`
}

type GoodsPropertyInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsPropertyInsertRequest() (request *GoodsPropertyInsertRequest) {
	request = &GoodsPropertyInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/property/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyInsertResponse() (response *api.BaseResponse[GoodsPropertyInsertResponse]) {
	response = api.CreateResponse[GoodsPropertyInsertResponse](&GoodsPropertyInsertResponse{})
	return
}

type GoodsPropertyInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
