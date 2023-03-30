package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTastesUnitUpdate
// @id 153
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新口味)
func (client *Client) GoodsTastesUnitUpdate(request *GoodsTastesUnitUpdateRequest) (response *GoodsTastesUnitUpdateResponse, err error) {
	rpcResponse := CreateGoodsTastesUnitUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTastesUnitUpdateRequest struct {
	*api.BaseRequest

	MerchantId        int64  `json:"merchantId,omitempty"`
	GoodsTastesUnitId int64  `json:"goodsTastesUnitId,omitempty"`
	TastesName        string `json:"tastesName,omitempty"`
}

type GoodsTastesUnitUpdateResponse struct {
}

func CreateGoodsTastesUnitUpdateRequest() (request *GoodsTastesUnitUpdateRequest) {
	request = &GoodsTastesUnitUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTastesUnit/update", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTastesUnitUpdateResponse() (response *api.BaseResponse[GoodsTastesUnitUpdateResponse]) {
	response = api.CreateResponse[GoodsTastesUnitUpdateResponse](&GoodsTastesUnitUpdateResponse{})
	return
}
