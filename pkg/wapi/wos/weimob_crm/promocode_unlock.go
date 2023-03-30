package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PromocodeUnlock
// @id 1595
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1595?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=优惠码解锁)
func (client *Client) PromocodeUnlock(request *PromocodeUnlockRequest) (response *PromocodeUnlockResponse, err error) {
	rpcResponse := CreatePromocodeUnlockResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PromocodeUnlockRequest struct {
	*api.BaseRequest

	CodeList []PromocodeUnlockRequestCodeList `json:"codeList,omitempty"`
	UseScene int64                            `json:"useScene,omitempty"`
	Wid      int64                            `json:"wid,omitempty"`
	VidType  int64                            `json:"vidType,omitempty"`
	Vid      int64                            `json:"vid,omitempty"`
}

type PromocodeUnlockResponse struct {
}

func CreatePromocodeUnlockRequest() (request *PromocodeUnlockRequest) {
	request = &PromocodeUnlockRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "promocode/unlock", "apigw")
	request.Method = api.POST
	return
}

func CreatePromocodeUnlockResponse() (response *api.BaseResponse[PromocodeUnlockResponse]) {
	response = api.CreateResponse[PromocodeUnlockResponse](&PromocodeUnlockResponse{})
	return
}

type PromocodeUnlockRequestCodeList struct {
	CodeTemplateId int64  `json:"codeTemplateId,omitempty"`
	OrderNo        string `json:"orderNo,omitempty"`
	Code           string `json:"code,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
}
