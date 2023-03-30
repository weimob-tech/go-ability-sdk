package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTastesUnitDelete
// @id 155
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除口味)
func (client *Client) GoodsTastesUnitDelete(request *GoodsTastesUnitDeleteRequest) (response *GoodsTastesUnitDeleteResponse, err error) {
	rpcResponse := CreateGoodsTastesUnitDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTastesUnitDeleteRequest struct {
	*api.BaseRequest

	MerchantId        int64 `json:"merchantId,omitempty"`
	GoodsTastesUnitId int64 `json:"goodsTastesUnitId,omitempty"`
}

type GoodsTastesUnitDeleteResponse struct {
}

func CreateGoodsTastesUnitDeleteRequest() (request *GoodsTastesUnitDeleteRequest) {
	request = &GoodsTastesUnitDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTastesUnit/delete", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTastesUnitDeleteResponse() (response *api.BaseResponse[GoodsTastesUnitDeleteResponse]) {
	response = api.CreateResponse[GoodsTastesUnitDeleteResponse](&GoodsTastesUnitDeleteResponse{})
	return
}
