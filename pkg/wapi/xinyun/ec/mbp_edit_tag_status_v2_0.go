package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpEditTagStatus
// @id 1661
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=启用、禁用、删除标签)
func (client *Client) MbpEditTagStatusV2_0(request *MbpEditTagStatusRequestV2_0) (response *MbpEditTagStatusResponseV2_0, err error) {
	rpcResponse := CreateMbpEditTagStatusResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpEditTagStatusRequestV2_0 struct {
	*api.BaseRequest

	TagId int64 `json:"tagId,omitempty"`
	Sign  int   `json:"sign,omitempty"`
}

type MbpEditTagStatusResponseV2_0 struct {
}

func CreateMbpEditTagStatusRequestV2_0() (request *MbpEditTagStatusRequestV2_0) {
	request = &MbpEditTagStatusRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/editTagStatus", "api")
	request.Method = api.POST
	return
}

func CreateMbpEditTagStatusResponseV2_0() (response *api.BaseResponse[MbpEditTagStatusResponseV2_0]) {
	response = api.CreateResponse[MbpEditTagStatusResponseV2_0](&MbpEditTagStatusResponseV2_0{})
	return
}
