package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpEditAttr
// @id 1663
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑属性)
func (client *Client) MbpEditAttrV2_0(request *MbpEditAttrRequestV2_0) (response *MbpEditAttrResponseV2_0, err error) {
	rpcResponse := CreateMbpEditAttrResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpEditAttrRequestV2_0 struct {
	*api.BaseRequest

	TagId    int64 `json:"tagId,omitempty"`
	AttrId   int64 `json:"attrId,omitempty"`
	AttrName int   `json:"attrName,omitempty"`
}

type MbpEditAttrResponseV2_0 struct {
}

func CreateMbpEditAttrRequestV2_0() (request *MbpEditAttrRequestV2_0) {
	request = &MbpEditAttrRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "2_0", "mbp/editAttr", "api")
	request.Method = api.POST
	return
}

func CreateMbpEditAttrResponseV2_0() (response *api.BaseResponse[MbpEditAttrResponseV2_0]) {
	response = api.CreateResponse[MbpEditAttrResponseV2_0](&MbpEditAttrResponseV2_0{})
	return
}
