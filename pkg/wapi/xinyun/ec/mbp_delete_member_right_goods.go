package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpDeleteMemberRightGoods
// @id 1785
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除商品折扣信息)
func (client *Client) MbpDeleteMemberRightGoods(request *MbpDeleteMemberRightGoodsRequest) (response *MbpDeleteMemberRightGoodsResponse, err error) {
	rpcResponse := CreateMbpDeleteMemberRightGoodsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpDeleteMemberRightGoodsRequest struct {
	*api.BaseRequest

	RuleId   int64   `json:"ruleId,omitempty"`
	GoodsIds []int64 `json:"goodsIds,omitempty"`
}

type MbpDeleteMemberRightGoodsResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMbpDeleteMemberRightGoodsRequest() (request *MbpDeleteMemberRightGoodsRequest) {
	request = &MbpDeleteMemberRightGoodsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/deleteMemberRightGoods", "api")
	request.Method = api.POST
	return
}

func CreateMbpDeleteMemberRightGoodsResponse() (response *api.BaseResponse[MbpDeleteMemberRightGoodsResponse]) {
	response = api.CreateResponse[MbpDeleteMemberRightGoodsResponse](&MbpDeleteMemberRightGoodsResponse{})
	return
}
