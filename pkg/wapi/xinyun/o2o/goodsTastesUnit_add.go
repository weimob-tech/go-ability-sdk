package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTastesUnitAdd
// @id 152
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加口味)
func (client *Client) GoodsTastesUnitAdd(request *GoodsTastesUnitAddRequest) (response *GoodsTastesUnitAddResponse, err error) {
	rpcResponse := CreateGoodsTastesUnitAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTastesUnitAddRequest struct {
	*api.BaseRequest

	MerchantId int64 `json:"merchantId,omitempty"`
	TastesName int64 `json:"tastesName,omitempty"`
}

type GoodsTastesUnitAddResponse struct {
}

func CreateGoodsTastesUnitAddRequest() (request *GoodsTastesUnitAddRequest) {
	request = &GoodsTastesUnitAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTastesUnit/add", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTastesUnitAddResponse() (response *api.BaseResponse[GoodsTastesUnitAddResponse]) {
	response = api.CreateResponse[GoodsTastesUnitAddResponse](&GoodsTastesUnitAddResponse{})
	return
}
