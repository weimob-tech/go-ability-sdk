package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembercardPrivilegeGoodsDelete
// @id 1772
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1772?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除折扣商品信息)
func (client *Client) MembercardPrivilegeGoodsDelete(request *MembercardPrivilegeGoodsDeleteRequest) (response *MembercardPrivilegeGoodsDeleteResponse, err error) {
	rpcResponse := CreateMembercardPrivilegeGoodsDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembercardPrivilegeGoodsDeleteRequest struct {
	*api.BaseRequest

	RuleId      int64   `json:"ruleId,omitempty"`
	GoodsIdList []int64 `json:"goodsIdList,omitempty"`
	Wid         int64   `json:"wid,omitempty"`
	VidType     int64   `json:"vidType,omitempty"`
	Vid         int64   `json:"vid,omitempty"`
}

type MembercardPrivilegeGoodsDeleteResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMembercardPrivilegeGoodsDeleteRequest() (request *MembercardPrivilegeGoodsDeleteRequest) {
	request = &MembercardPrivilegeGoodsDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "membercard/privilege/goods/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateMembercardPrivilegeGoodsDeleteResponse() (response *api.BaseResponse[MembercardPrivilegeGoodsDeleteResponse]) {
	response = api.CreateResponse[MembercardPrivilegeGoodsDeleteResponse](&MembercardPrivilegeGoodsDeleteResponse{})
	return
}
