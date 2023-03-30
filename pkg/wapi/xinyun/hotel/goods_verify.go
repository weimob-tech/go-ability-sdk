package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsVerify
// @id 504
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=商品消核)
func (client *Client) GoodsVerify(request *GoodsVerifyRequest) (response *GoodsVerifyResponse, err error) {
	rpcResponse := CreateGoodsVerifyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsVerifyRequest struct {
	*api.BaseRequest

	SnNo    bool  `json:"snNo,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type GoodsVerifyResponse struct {
}

func CreateGoodsVerifyRequest() (request *GoodsVerifyRequest) {
	request = &GoodsVerifyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "goods/verify", "api")
	request.Method = api.POST
	return
}

func CreateGoodsVerifyResponse() (response *api.BaseResponse[GoodsVerifyResponse]) {
	response = api.CreateResponse[GoodsVerifyResponse](&GoodsVerifyResponse{})
	return
}
