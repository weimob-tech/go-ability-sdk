package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuUnitEdit
// @id 266
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑菜品规格)
func (client *Client) SkuUnitEdit(request *SkuUnitEditRequest) (response *SkuUnitEditResponse, err error) {
	rpcResponse := CreateSkuUnitEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuUnitEditRequest struct {
	*api.BaseRequest

	SkuUnitName string `json:"skuUnitName,omitempty"`
	SkuUnitId   int64  `json:"skuUnitId,omitempty"`
	MerchantId  int64  `json:"merchantId,omitempty"`
	ThirdId     string `json:"thirdId,omitempty"`
}

type SkuUnitEditResponse struct {
}

func CreateSkuUnitEditRequest() (request *SkuUnitEditRequest) {
	request = &SkuUnitEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "skuUnit/edit", "api")
	request.Method = api.POST
	return
}

func CreateSkuUnitEditResponse() (response *api.BaseResponse[SkuUnitEditResponse]) {
	response = api.CreateResponse[SkuUnitEditResponse](&SkuUnitEditResponse{})
	return
}
