package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CategoryAttributeAdd
// @id 3
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品类目属性)
func (client *Client) CategoryAttributeAdd(request *CategoryAttributeAddRequest) (response *CategoryAttributeAddResponse, err error) {
	rpcResponse := CreateCategoryAttributeAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CategoryAttributeAddRequest struct {
	*api.BaseRequest

	CategoryId         string `json:"category_id,omitempty"`
	CategoryAttrName   string `json:"category_attr_name,omitempty"`
	CategoryAttrValues string `json:"category_attr_values,omitempty"`
}

type CategoryAttributeAddResponse struct {
}

func CreateCategoryAttributeAddRequest() (request *CategoryAttributeAddRequest) {
	request = &CategoryAttributeAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "CategoryAttribute/Add", "api")
	request.Method = api.POST
	return
}

func CreateCategoryAttributeAddResponse() (response *api.BaseResponse[CategoryAttributeAddResponse]) {
	response = api.CreateResponse[CategoryAttributeAddResponse](&CategoryAttributeAddResponse{})
	return
}
