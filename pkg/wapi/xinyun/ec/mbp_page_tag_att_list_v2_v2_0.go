package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpPageTagAttListV2
// @id 1658
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询标签库信息V2)
func (client *Client) MbpPageTagAttListV2V2_0(request *MbpPageTagAttListV2RequestV2_0) (response *MbpPageTagAttListV2ResponseV2_0, err error) {
	rpcResponse := CreateMbpPageTagAttListV2ResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpPageTagAttListV2RequestV2_0 struct {
	*api.BaseRequest

	PageNum        int    `json:"pageNum,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	QueryParameter string `json:"queryParameter,omitempty"`
	Pid            int64  `json:"pid,omitempty"`
}

type MbpPageTagAttListV2ResponseV2_0 struct {
	PageList   string `json:"pageList,omitempty"`
	AttrVoList string `json:"attrVoList,omitempty"`
	TotalCount int64  `json:"totalCount,omitempty"`
}

func CreateMbpPageTagAttListV2RequestV2_0() (request *MbpPageTagAttListV2RequestV2_0) {
	request = &MbpPageTagAttListV2RequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/pageTagAttListV2", "api")
	request.Method = api.POST
	return
}

func CreateMbpPageTagAttListV2ResponseV2_0() (response *api.BaseResponse[MbpPageTagAttListV2ResponseV2_0]) {
	response = api.CreateResponse[MbpPageTagAttListV2ResponseV2_0](&MbpPageTagAttListV2ResponseV2_0{})
	return
}
