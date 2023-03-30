package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTastesUnitSyn
// @id 240
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=同步菜品口味)
func (client *Client) GoodsTastesUnitSyn(request *GoodsTastesUnitSynRequest) (response *GoodsTastesUnitSynResponse, err error) {
	rpcResponse := CreateGoodsTastesUnitSynResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTastesUnitSynRequest struct {
	*api.BaseRequest

	MerchantId    int64   `json:"merchantId,omitempty"`
	Tastes        []int64 `json:"tastes,omitempty"`
	TastesName    string  `json:"tastesName,omitempty"`
	ThirdTastesId string  `json:"thirdTastesId,omitempty"`
}

type GoodsTastesUnitSynResponse struct {
}

func CreateGoodsTastesUnitSynRequest() (request *GoodsTastesUnitSynRequest) {
	request = &GoodsTastesUnitSynRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTastesUnit/syn", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTastesUnitSynResponse() (response *api.BaseResponse[GoodsTastesUnitSynResponse]) {
	response = api.CreateResponse[GoodsTastesUnitSynResponse](&GoodsTastesUnitSynResponse{})
	return
}
