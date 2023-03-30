package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryAgreementProtocol
// @id 1998
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询同意协议记录)
func (client *Client) MbpQueryAgreementProtocol(request *MbpQueryAgreementProtocolRequest) (response *MbpQueryAgreementProtocolResponse, err error) {
	rpcResponse := CreateMbpQueryAgreementProtocolResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryAgreementProtocolRequest struct {
	*api.BaseRequest

	Pid int64 `json:"pid,omitempty"`
	Wid int64 `json:"wid,omitempty"`
}

type MbpQueryAgreementProtocolResponse struct {
}

func CreateMbpQueryAgreementProtocolRequest() (request *MbpQueryAgreementProtocolRequest) {
	request = &MbpQueryAgreementProtocolRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryAgreementProtocol", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryAgreementProtocolResponse() (response *api.BaseResponse[MbpQueryAgreementProtocolResponse]) {
	response = api.CreateResponse[MbpQueryAgreementProtocolResponse](&MbpQueryAgreementProtocolResponse{})
	return
}
