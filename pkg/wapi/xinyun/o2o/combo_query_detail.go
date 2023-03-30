package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ComboQueryDetail
// @id 258
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询套餐详情)
func (client *Client) ComboQueryDetail(request *ComboQueryDetailRequest) (response *ComboQueryDetailResponse, err error) {
	rpcResponse := CreateComboQueryDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ComboQueryDetailRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
	ComboId int64 `json:"comboId,omitempty"`
	SkuId   int64 `json:"skuId,omitempty"`
}

type ComboQueryDetailResponse struct {
}

func CreateComboQueryDetailRequest() (request *ComboQueryDetailRequest) {
	request = &ComboQueryDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "combo/queryDetail", "api")
	request.Method = api.POST
	return
}

func CreateComboQueryDetailResponse() (response *api.BaseResponse[ComboQueryDetailResponse]) {
	response = api.CreateResponse[ComboQueryDetailResponse](&ComboQueryDetailResponse{})
	return
}
