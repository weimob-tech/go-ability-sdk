package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpGetCrowdVersion
// @id 1614
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导入线下人群-新建版本)
func (client *Client) MbpGetCrowdVersionV2_0(request *MbpGetCrowdVersionRequestV2_0) (response *MbpGetCrowdVersionResponseV2_0, err error) {
	rpcResponse := CreateMbpGetCrowdVersionResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpGetCrowdVersionRequestV2_0 struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
	CrowdId int64 `json:"crowdId,omitempty"`
}

type MbpGetCrowdVersionResponseV2_0 struct {
}

func CreateMbpGetCrowdVersionRequestV2_0() (request *MbpGetCrowdVersionRequestV2_0) {
	request = &MbpGetCrowdVersionRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/getCrowdVersion", "api")
	request.Method = api.POST
	return
}

func CreateMbpGetCrowdVersionResponseV2_0() (response *api.BaseResponse[MbpGetCrowdVersionResponseV2_0]) {
	response = api.CreateResponse[MbpGetCrowdVersionResponseV2_0](&MbpGetCrowdVersionResponseV2_0{})
	return
}
