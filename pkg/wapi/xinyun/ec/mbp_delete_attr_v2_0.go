package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpDeleteAttr
// @id 1664
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除属性)
func (client *Client) MbpDeleteAttrV2_0(request *MbpDeleteAttrRequestV2_0) (response *MbpDeleteAttrResponseV2_0, err error) {
	rpcResponse := CreateMbpDeleteAttrResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpDeleteAttrRequestV2_0 struct {
	*api.BaseRequest

	TagId  int64 `json:"tagId,omitempty"`
	AttrId int64 `json:"attrId,omitempty"`
}

type MbpDeleteAttrResponseV2_0 struct {
}

func CreateMbpDeleteAttrRequestV2_0() (request *MbpDeleteAttrRequestV2_0) {
	request = &MbpDeleteAttrRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/deleteAttr", "api")
	request.Method = api.POST
	return
}

func CreateMbpDeleteAttrResponseV2_0() (response *api.BaseResponse[MbpDeleteAttrResponseV2_0]) {
	response = api.CreateResponse[MbpDeleteAttrResponseV2_0](&MbpDeleteAttrResponseV2_0{})
	return
}
