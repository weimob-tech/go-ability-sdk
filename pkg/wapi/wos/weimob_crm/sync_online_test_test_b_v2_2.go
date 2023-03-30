package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SyncOnlineTestTestB
// @id 2141
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2141?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=sync-online-test22333444)
func (client *Client) SyncOnlineTestTestBV2_2(request *SyncOnlineTestTestBRequest) (response *SyncOnlineTestTestBResponse, err error) {
	rpcResponse := CreateSyncOnlineTestTestBResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SyncOnlineTestTestBRequest struct {
	*api.BaseRequest

	Test123 string `json:"test123,omitempty"`
}

type SyncOnlineTestTestBResponse struct {
	Abc bool `json:"abc,omitempty"`
}

func CreateSyncOnlineTestTestBRequest() (request *SyncOnlineTestTestBRequest) {
	request = &SyncOnlineTestTestBRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.2", "sync/online/test/testB", "apigw")
	request.Method = api.POST
	return
}

func CreateSyncOnlineTestTestBResponse() (response *api.BaseResponse[SyncOnlineTestTestBResponse]) {
	response = api.CreateResponse[SyncOnlineTestTestBResponse](&SyncOnlineTestTestBResponse{})
	return
}
