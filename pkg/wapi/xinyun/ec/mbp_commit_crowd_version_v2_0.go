package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpCommitCrowdVersion
// @id 1616
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导入线下人群-提交版本)
func (client *Client) MbpCommitCrowdVersionV2_0(request *MbpCommitCrowdVersionRequestV2_0) (response *MbpCommitCrowdVersionResponseV2_0, err error) {
	rpcResponse := CreateMbpCommitCrowdVersionResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpCommitCrowdVersionRequestV2_0 struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
	CrowdId int64 `json:"crowdId,omitempty"`
	Version int   `json:"version,omitempty"`
}

type MbpCommitCrowdVersionResponseV2_0 struct {
}

func CreateMbpCommitCrowdVersionRequestV2_0() (request *MbpCommitCrowdVersionRequestV2_0) {
	request = &MbpCommitCrowdVersionRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/commitCrowdVersion", "api")
	request.Method = api.POST
	return
}

func CreateMbpCommitCrowdVersionResponseV2_0() (response *api.BaseResponse[MbpCommitCrowdVersionResponseV2_0]) {
	response = api.CreateResponse[MbpCommitCrowdVersionResponseV2_0](&MbpCommitCrowdVersionResponseV2_0{})
	return
}
