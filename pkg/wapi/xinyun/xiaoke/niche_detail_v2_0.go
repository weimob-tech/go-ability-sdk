package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NicheDetail
// @id 2671
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商机详情)
func (client *Client) NicheDetailV2_0(request *NicheDetailRequestV2_0) (response *NicheDetailResponseV2_0, err error) {
	rpcResponse := CreateNicheDetailResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NicheDetailRequestV2_0 struct {
	*api.BaseRequest

	Wid int    `json:"wid,omitempty"`
	Key string `json:"key,omitempty"`
}

type NicheDetailResponseV2_0 struct {
}

func CreateNicheDetailRequestV2_0() (request *NicheDetailRequestV2_0) {
	request = &NicheDetailRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "2_0", "niche/detail", "api")
	request.Method = api.POST
	return
}

func CreateNicheDetailResponseV2_0() (response *api.BaseResponse[NicheDetailResponseV2_0]) {
	response = api.CreateResponse[NicheDetailResponseV2_0](&NicheDetailResponseV2_0{})
	return
}
