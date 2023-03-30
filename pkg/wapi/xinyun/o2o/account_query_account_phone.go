package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AccountQueryAccountPhone
// @id 1339
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据操作员Id查询操作员手机号)
func (client *Client) AccountQueryAccountPhone(request *AccountQueryAccountPhoneRequest) (response *AccountQueryAccountPhoneResponse, err error) {
	rpcResponse := CreateAccountQueryAccountPhoneResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AccountQueryAccountPhoneRequest struct {
	*api.BaseRequest

	MerchantWid int64 `json:"merchantWid,omitempty"`
}

type AccountQueryAccountPhoneResponse struct {
}

func CreateAccountQueryAccountPhoneRequest() (request *AccountQueryAccountPhoneRequest) {
	request = &AccountQueryAccountPhoneRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "account/queryAccountPhone", "api")
	request.Method = api.POST
	return
}

func CreateAccountQueryAccountPhoneResponse() (response *api.BaseResponse[AccountQueryAccountPhoneResponse]) {
	response = api.CreateResponse[AccountQueryAccountPhoneResponse](&AccountQueryAccountPhoneResponse{})
	return
}
