package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpPrepareDateForCrowd
// @id 1615
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导入线下人群-准备数据)
func (client *Client) MbpPrepareDateForCrowdV2_0(request *MbpPrepareDateForCrowdRequestV2_0) (response *MbpPrepareDateForCrowdResponseV2_0, err error) {
	rpcResponse := CreateMbpPrepareDateForCrowdResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpPrepareDateForCrowdRequestV2_0 struct {
	*api.BaseRequest

	WidList []int `json:"widList,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
	CrowdId int64 `json:"crowdId,omitempty"`
	Version int   `json:"version,omitempty"`
}

type MbpPrepareDateForCrowdResponseV2_0 struct {
}

func CreateMbpPrepareDateForCrowdRequestV2_0() (request *MbpPrepareDateForCrowdRequestV2_0) {
	request = &MbpPrepareDateForCrowdRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/prepareDateForCrowd", "api")
	request.Method = api.POST
	return
}

func CreateMbpPrepareDateForCrowdResponseV2_0() (response *api.BaseResponse[MbpPrepareDateForCrowdResponseV2_0]) {
	response = api.CreateResponse[MbpPrepareDateForCrowdResponseV2_0](&MbpPrepareDateForCrowdResponseV2_0{})
	return
}
