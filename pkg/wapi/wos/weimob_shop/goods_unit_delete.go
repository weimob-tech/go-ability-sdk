package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsUnitDelete
// @id 1855
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1855?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除商品单位)
func (client *Client) GoodsUnitDelete(request *GoodsUnitDeleteRequest) (response *GoodsUnitDeleteResponse, err error) {
	rpcResponse := CreateGoodsUnitDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsUnitDeleteRequest struct {
	*api.BaseRequest

	BasicInfo GoodsUnitDeleteRequestBasicInfo `json:"basicInfo,omitempty"`
	UnitId    int64                           `json:"unitId,omitempty"`
}

type GoodsUnitDeleteResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsUnitDeleteRequest() (request *GoodsUnitDeleteRequest) {
	request = &GoodsUnitDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/unit/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsUnitDeleteResponse() (response *api.BaseResponse[GoodsUnitDeleteResponse]) {
	response = api.CreateResponse[GoodsUnitDeleteResponse](&GoodsUnitDeleteResponse{})
	return
}

type GoodsUnitDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
