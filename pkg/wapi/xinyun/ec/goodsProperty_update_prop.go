package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyUpdateProp
// @id 1168
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新商品属性)
func (client *Client) GoodsPropertyUpdateProp(request *GoodsPropertyUpdatePropRequest) (response *GoodsPropertyUpdatePropResponse, err error) {
	rpcResponse := CreateGoodsPropertyUpdatePropResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyUpdatePropRequest struct {
	*api.BaseRequest

	Title           string `json:"title,omitempty"`
	Sort            int    `json:"sort,omitempty"`
	SupportModule   []int  `json:"supportModule,omitempty"`
	PropValueIdList []int  `json:"propValueIdList,omitempty"`
}

type GoodsPropertyUpdatePropResponse struct {
}

func CreateGoodsPropertyUpdatePropRequest() (request *GoodsPropertyUpdatePropRequest) {
	request = &GoodsPropertyUpdatePropRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsProperty/updateProp", "api")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyUpdatePropResponse() (response *api.BaseResponse[GoodsPropertyUpdatePropResponse]) {
	response = api.CreateResponse[GoodsPropertyUpdatePropResponse](&GoodsPropertyUpdatePropResponse{})
	return
}
