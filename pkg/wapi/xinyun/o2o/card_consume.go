package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardConsume
// @id 206
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠券)
func (client *Client) CardConsume(request *CardConsumeRequest) (response *CardConsumeResponse, err error) {
	rpcResponse := CreateCardConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardConsumeRequest struct {
	*api.BaseRequest

	CardCode        string  `json:"cardCode,omitempty"`
	UseStoreId      int64   `json:"useStoreId,omitempty"`
	ConsumeAmount   float32 `json:"consumeAmount,omitempty"`
	ThirdBusinessId string  `json:"thirdBusinessId,omitempty"`
	ReturnCardInfo  bool    `json:"returnCardInfo,omitempty"`
}

type CardConsumeResponse struct {
}

func CreateCardConsumeRequest() (request *CardConsumeRequest) {
	request = &CardConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "card/consume", "api")
	request.Method = api.POST
	return
}

func CreateCardConsumeResponse() (response *api.BaseResponse[CardConsumeResponse]) {
	response = api.CreateResponse[CardConsumeResponse](&CardConsumeResponse{})
	return
}
