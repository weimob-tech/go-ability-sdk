package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuUnitAdd
// @id 264
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加菜品规格名称)
func (client *Client) SkuUnitAdd(request *SkuUnitAddRequest) (response *SkuUnitAddResponse, err error) {
	rpcResponse := CreateSkuUnitAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuUnitAddRequest struct {
	*api.BaseRequest

	SkuUnitName string `json:"skuUnitName,omitempty"`
	MerchantId  int64  `json:"merchantId,omitempty"`
	ThirdId     string `json:"thirdId,omitempty"`
}

type SkuUnitAddResponse struct {
}

func CreateSkuUnitAddRequest() (request *SkuUnitAddRequest) {
	request = &SkuUnitAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "skuUnit/add", "api")
	request.Method = api.POST
	return
}

func CreateSkuUnitAddResponse() (response *api.BaseResponse[SkuUnitAddResponse]) {
	response = api.CreateResponse[SkuUnitAddResponse](&SkuUnitAddResponse{})
	return
}
