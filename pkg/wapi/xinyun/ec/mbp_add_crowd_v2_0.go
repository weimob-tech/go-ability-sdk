package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpAddCrowd
// @id 1613
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新建人群)
func (client *Client) MbpAddCrowdV2_0(request *MbpAddCrowdRequestV2_0) (response *MbpAddCrowdResponseV2_0, err error) {
	rpcResponse := CreateMbpAddCrowdResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpAddCrowdRequestV2_0 struct {
	*api.BaseRequest

	StoreId         int64  `json:"storeId,omitempty"`
	CrowdName       string `json:"crowdName,omitempty"`
	Description     string `json:"description,omitempty"`
	CrowdCategoryId int64  `json:"crowdCategoryId,omitempty"`
}

type MbpAddCrowdResponseV2_0 struct {
}

func CreateMbpAddCrowdRequestV2_0() (request *MbpAddCrowdRequestV2_0) {
	request = &MbpAddCrowdRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/addCrowd", "api")
	request.Method = api.POST
	return
}

func CreateMbpAddCrowdResponseV2_0() (response *api.BaseResponse[MbpAddCrowdResponseV2_0]) {
	response = api.CreateResponse[MbpAddCrowdResponseV2_0](&MbpAddCrowdResponseV2_0{})
	return
}
