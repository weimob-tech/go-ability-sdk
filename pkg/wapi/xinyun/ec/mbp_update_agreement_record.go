package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpUpdateAgreementRecord
// @id 3754
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新协议授权状态)
func (client *Client) MbpUpdateAgreementRecord(request *MbpUpdateAgreementRecordRequest) (response *MbpUpdateAgreementRecordResponse, err error) {
	rpcResponse := CreateMbpUpdateAgreementRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpUpdateAgreementRecordRequest struct {
	*api.BaseRequest
}

type MbpUpdateAgreementRecordResponse struct {
}

func CreateMbpUpdateAgreementRecordRequest() (request *MbpUpdateAgreementRecordRequest) {
	request = &MbpUpdateAgreementRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/updateAgreementRecord", "api")
	request.Method = api.POST
	return
}

func CreateMbpUpdateAgreementRecordResponse() (response *api.BaseResponse[MbpUpdateAgreementRecordResponse]) {
	response = api.CreateResponse[MbpUpdateAgreementRecordResponse](&MbpUpdateAgreementRecordResponse{})
	return
}
