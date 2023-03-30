package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerGroupGetList
// @id 1776
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1776?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询买家秀分组列表)
func (client *Client) CommentShowBuyerGroupGetList(request *CommentShowBuyerGroupGetListRequest) (response *CommentShowBuyerGroupGetListResponse, err error) {
	rpcResponse := CreateCommentShowBuyerGroupGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerGroupGetListRequest struct {
	*api.BaseRequest
}

type CommentShowBuyerGroupGetListResponse struct {
	CommentGroups []CommentShowBuyerGroupGetListResponseCommentGroups `json:"commentGroups,omitempty"`
}

func CreateCommentShowBuyerGroupGetListRequest() (request *CommentShowBuyerGroupGetListRequest) {
	request = &CommentShowBuyerGroupGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/group/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerGroupGetListResponse() (response *api.BaseResponse[CommentShowBuyerGroupGetListResponse]) {
	response = api.CreateResponse[CommentShowBuyerGroupGetListResponse](&CommentShowBuyerGroupGetListResponse{})
	return
}

type CommentShowBuyerGroupGetListResponseCommentGroups struct {
	GoodsGroups       []CommentShowBuyerGroupGetListResponseGoodsGroups `json:"goodsGroups,omitempty"`
	OperateId         int64                                             `json:"operateId,omitempty"`
	GroupName         string                                            `json:"groupName,omitempty"`
	GroupType         int64                                             `json:"groupType,omitempty"`
	RelateCount       int64                                             `json:"relateCount,omitempty"`
	Id                int64                                             `json:"id,omitempty"`
	MerchantId        int64                                             `json:"merchantId,omitempty"`
	BosId             int64                                             `json:"bosId,omitempty"`
	Vid               int64                                             `json:"vid,omitempty"`
	VidPath           string                                            `json:"vidPath,omitempty"`
	ProductId         int64                                             `json:"productId,omitempty"`
	ProductInstanceId int64                                             `json:"productInstanceId,omitempty"`
}

type CommentShowBuyerGroupGetListResponseGoodsGroups struct {
	GroupId   int64  `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ParentId  int64  `json:"parentId,omitempty"`
}
