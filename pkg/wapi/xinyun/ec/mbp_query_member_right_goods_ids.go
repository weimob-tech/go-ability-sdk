package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MbpQueryMemberRightGoodsIds
// @id 1782
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员权益商品信息)
func (client *Client) MbpQueryMemberRightGoodsIds(request *MbpQueryMemberRightGoodsIdsRequest) (response *MbpQueryMemberRightGoodsIdsResponse, err error) {
	rpcResponse := CreateMbpQueryMemberRightGoodsIdsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MbpQueryMemberRightGoodsIdsRequest struct {
	*api.BaseRequest

	RuleId   int64   `json:"ruleId,omitempty"`
	GoodsIds []int64 `json:"goodsIds,omitempty"`
}

type MbpQueryMemberRightGoodsIdsResponse struct {
	GoodsIds []int64 `json:"goodsIds,omitempty"`
}

func CreateMbpQueryMemberRightGoodsIdsRequest() (request *MbpQueryMemberRightGoodsIdsRequest) {
	request = &MbpQueryMemberRightGoodsIdsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "mbp/queryMemberRightGoodsIds", "api")
	request.Method = api.POST
	return
}

func CreateMbpQueryMemberRightGoodsIdsResponse() (response *api.BaseResponse[MbpQueryMemberRightGoodsIdsResponse]) {
	response = api.CreateResponse[MbpQueryMemberRightGoodsIdsResponse](&MbpQueryMemberRightGoodsIdsResponse{})
	return
}
