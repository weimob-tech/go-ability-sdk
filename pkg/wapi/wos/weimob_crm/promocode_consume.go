package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromocodeConsume
// @id 1596
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1596?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=核销优惠码)
func (client *Client) PromocodeConsume(request *PromocodeConsumeRequest) (response *PromocodeConsumeResponse, err error) {
	rpcResponse := CreatePromocodeConsumeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromocodeConsumeRequest struct {
	*api.BaseRequest

	CodeList        []PromocodeConsumeRequestCodeList `json:"codeList,omitempty"`
	UseResourceType int64                             `json:"useResourceType,omitempty"`
	Operator        int64                             `json:"operator,omitempty"`
	UseScene        int64                             `json:"useScene,omitempty"`
	Wid             int64                             `json:"wid,omitempty"`
	VidType         int64                             `json:"vidType,omitempty"`
	Vid             int64                             `json:"vid,omitempty"`
}

type PromocodeConsumeResponse struct {
}

func CreatePromocodeConsumeRequest() (request *PromocodeConsumeRequest) {
	request = &PromocodeConsumeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "promocode/consume", "apigw")
	request.Method = api.POST
	return
}

func CreatePromocodeConsumeResponse() (response *api.BaseResponse[PromocodeConsumeResponse]) {
	response = api.CreateResponse[PromocodeConsumeResponse](&PromocodeConsumeResponse{})
	return
}

type PromocodeConsumeRequestCodeList struct {
	CodeTemplateId int64  `json:"codeTemplateId,omitempty"`
	OrderNo        string `json:"orderNo,omitempty"`
	Code           string `json:"code,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
}
