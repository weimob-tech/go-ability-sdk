package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// DiscountSaveDiscountInfo
// @id 2133
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=限时折扣创建(作废))
func (client *Client) DiscountSaveDiscountInfo(request *DiscountSaveDiscountInfoRequest) (response *DiscountSaveDiscountInfoResponse, err error) {
	rpcResponse := CreateDiscountSaveDiscountInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type DiscountSaveDiscountInfoRequest struct {
	*api.BaseRequest
}

type DiscountSaveDiscountInfoResponse struct {
}

func CreateDiscountSaveDiscountInfoRequest() (request *DiscountSaveDiscountInfoRequest) {
	request = &DiscountSaveDiscountInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "discount/saveDiscountInfo", "api")
	request.Method = api.POST
	return
}

func CreateDiscountSaveDiscountInfoResponse() (response *api.BaseResponse[DiscountSaveDiscountInfoResponse]) {
	response = api.CreateResponse[DiscountSaveDiscountInfoResponse](&DiscountSaveDiscountInfoResponse{})
	return
}
