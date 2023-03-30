package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpEditTag
// @id 1660
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑标签)
func (client *Client) MbpEditTagV2_0(request *MbpEditTagRequestV2_0) (response *MbpEditTagResponseV2_0, err error) {
	rpcResponse := CreateMbpEditTagResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpEditTagRequestV2_0 struct {
	*api.BaseRequest

	TagId   int64  `json:"tagId,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

type MbpEditTagResponseV2_0 struct {
}

func CreateMbpEditTagRequestV2_0() (request *MbpEditTagRequestV2_0) {
	request = &MbpEditTagRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/editTag", "api")
	request.Method = api.POST
	return
}

func CreateMbpEditTagResponseV2_0() (response *api.BaseResponse[MbpEditTagResponseV2_0]) {
	response = api.CreateResponse[MbpEditTagResponseV2_0](&MbpEditTagResponseV2_0{})
	return
}
