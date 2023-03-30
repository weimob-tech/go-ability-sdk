package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SkuAttributeAdd
// @id 39
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品属性)
func (client *Client) SkuAttributeAdd(request *SkuAttributeAddRequest) (response *SkuAttributeAddResponse, err error) {
	rpcResponse := CreateSkuAttributeAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SkuAttributeAddRequest struct {
	*api.BaseRequest

	CategoryId    string   `json:"category_id,omitempty"`
	SkuAttrName   string   `json:"sku_attr_name,omitempty"`
	SkuAttrValues []string `json:"sku_attr_values,omitempty"`
}

type SkuAttributeAddResponse struct {
}

func CreateSkuAttributeAddRequest() (request *SkuAttributeAddRequest) {
	request = &SkuAttributeAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "SkuAttribute/Add", "api")
	request.Method = api.POST
	return
}

func CreateSkuAttributeAddResponse() (response *api.BaseResponse[SkuAttributeAddResponse]) {
	response = api.CreateResponse[SkuAttributeAddResponse](&SkuAttributeAddResponse{})
	return
}
