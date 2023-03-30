package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagEdit
// @id 77
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改标签)
func (client *Client) GoodsTagEdit(request *GoodsTagEditRequest) (response *GoodsTagEditResponse, err error) {
	rpcResponse := CreateGoodsTagEditResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagEditRequest struct {
	*api.BaseRequest

	GoodsTagId int64  `json:"goodsTagId,omitempty"`
	TagName    string `json:"tagName,omitempty"`
	StoreId    int64  `json:"storeId,omitempty"`
	MerchantId int64  `json:"merchantId,omitempty"`
}

type GoodsTagEditResponse struct {
}

func CreateGoodsTagEditRequest() (request *GoodsTagEditRequest) {
	request = &GoodsTagEditRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTag/edit", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagEditResponse() (response *api.BaseResponse[GoodsTagEditResponse]) {
	response = api.CreateResponse[GoodsTagEditResponse](&GoodsTagEditResponse{})
	return
}
