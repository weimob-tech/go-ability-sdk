package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyAddProp
// @id 1167
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品属性)
func (client *Client) GoodsPropertyAddProp(request *GoodsPropertyAddPropRequest) (response *GoodsPropertyAddPropResponse, err error) {
	rpcResponse := CreateGoodsPropertyAddPropResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyAddPropRequest struct {
	*api.BaseRequest

	Title         string   `json:"title,omitempty"`
	Sort          int      `json:"sort,omitempty"`
	SupportModule []int    `json:"supportModule,omitempty"`
	PropValList   []string `json:"propValList,omitempty"`
	PropType      int      `json:"propType,omitempty"`
}

type GoodsPropertyAddPropResponse struct {
}

func CreateGoodsPropertyAddPropRequest() (request *GoodsPropertyAddPropRequest) {
	request = &GoodsPropertyAddPropRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsProperty/addProp", "api")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyAddPropResponse() (response *api.BaseResponse[GoodsPropertyAddPropResponse]) {
	response = api.CreateResponse[GoodsPropertyAddPropResponse](&GoodsPropertyAddPropResponse{})
	return
}
