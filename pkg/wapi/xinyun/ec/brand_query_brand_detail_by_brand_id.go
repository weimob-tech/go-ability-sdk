package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BrandQueryBrandDetailByBrandId
// @id 1625
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询品牌详情)
func (client *Client) BrandQueryBrandDetailByBrandId(request *BrandQueryBrandDetailByBrandIdRequest) (response *BrandQueryBrandDetailByBrandIdResponse, err error) {
	rpcResponse := CreateBrandQueryBrandDetailByBrandIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BrandQueryBrandDetailByBrandIdRequest struct {
	*api.BaseRequest

	BrandId int64 `json:"brandId,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type BrandQueryBrandDetailByBrandIdResponse struct {
}

func CreateBrandQueryBrandDetailByBrandIdRequest() (request *BrandQueryBrandDetailByBrandIdRequest) {
	request = &BrandQueryBrandDetailByBrandIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "brand/queryBrandDetailByBrandId", "api")
	request.Method = api.POST
	return
}

func CreateBrandQueryBrandDetailByBrandIdResponse() (response *api.BaseResponse[BrandQueryBrandDetailByBrandIdResponse]) {
	response = api.CreateResponse[BrandQueryBrandDetailByBrandIdResponse](&BrandQueryBrandDetailByBrandIdResponse{})
	return
}
