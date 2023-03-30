package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagAdd
// @id 76
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加菜品标签)
func (client *Client) GoodsTagAdd(request *GoodsTagAddRequest) (response *GoodsTagAddResponse, err error) {
	rpcResponse := CreateGoodsTagAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagAddRequest struct {
	*api.BaseRequest

	TagName    string `json:"tagName,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
	MerchantId int64  `json:"merchantId,omitempty"`
}

type GoodsTagAddResponse struct {
}

func CreateGoodsTagAddRequest() (request *GoodsTagAddRequest) {
	request = &GoodsTagAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTag/add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagAddResponse() (response *api.BaseResponse[GoodsTagAddResponse]) {
	response = api.CreateResponse[GoodsTagAddResponse](&GoodsTagAddResponse{})
	return
}
