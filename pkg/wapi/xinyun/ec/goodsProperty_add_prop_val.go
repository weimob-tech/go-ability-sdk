package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyAddPropVal
// @id 1172
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品属性值)
func (client *Client) GoodsPropertyAddPropVal(request *GoodsPropertyAddPropValRequest) (response *GoodsPropertyAddPropValResponse, err error) {
	rpcResponse := CreateGoodsPropertyAddPropValResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyAddPropValRequest struct {
	*api.BaseRequest

	PropId    int64  `json:"propId,omitempty"`
	PropValue string `json:"propValue,omitempty"`
}

type GoodsPropertyAddPropValResponse struct {
}

func CreateGoodsPropertyAddPropValRequest() (request *GoodsPropertyAddPropValRequest) {
	request = &GoodsPropertyAddPropValRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsProperty/addPropVal", "api")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyAddPropValResponse() (response *api.BaseResponse[GoodsPropertyAddPropValResponse]) {
	response = api.CreateResponse[GoodsPropertyAddPropValResponse](&GoodsPropertyAddPropValResponse{})
	return
}
