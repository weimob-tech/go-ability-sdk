package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUnitInsert
// @id 1626
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1626?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=新增商品单位)
func (client *Client) GoodsUnitInsert(request *GoodsUnitInsertRequest) (response *GoodsUnitInsertResponse, err error) {
	rpcResponse := CreateGoodsUnitInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUnitInsertRequest struct {
	*api.BaseRequest

	BasicInfo GoodsUnitInsertRequestBasicInfo `json:"basicInfo,omitempty"`
	UnitTitle string                          `json:"unitTitle,omitempty"`
}

type GoodsUnitInsertResponse struct {
	Id int64 `json:"id,omitempty"`
}

func CreateGoodsUnitInsertRequest() (request *GoodsUnitInsertRequest) {
	request = &GoodsUnitInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/unit/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUnitInsertResponse() (response *api.BaseResponse[GoodsUnitInsertResponse]) {
	response = api.CreateResponse[GoodsUnitInsertResponse](&GoodsUnitInsertResponse{})
	return
}

type GoodsUnitInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
