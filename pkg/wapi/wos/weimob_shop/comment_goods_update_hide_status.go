package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsUpdateHideStatus
// @id 1783
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1783?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品评论显示状态)
func (client *Client) CommentGoodsUpdateHideStatus(request *CommentGoodsUpdateHideStatusRequest) (response *CommentGoodsUpdateHideStatusResponse, err error) {
	rpcResponse := CreateCommentGoodsUpdateHideStatusResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsUpdateHideStatusRequest struct {
	*api.BaseRequest

	BasicInfo  CommentGoodsUpdateHideStatusRequestBasicInfo `json:"basicInfo,omitempty"`
	CommentIds []int64                                      `json:"commentIds,omitempty"`
	IsHide     int64                                        `json:"isHide,omitempty"`
}

type CommentGoodsUpdateHideStatusResponse struct {
	UpdateResult bool `json:"updateResult,omitempty"`
}

func CreateCommentGoodsUpdateHideStatusRequest() (request *CommentGoodsUpdateHideStatusRequest) {
	request = &CommentGoodsUpdateHideStatusRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/updateHideStatus", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsUpdateHideStatusResponse() (response *api.BaseResponse[CommentGoodsUpdateHideStatusResponse]) {
	response = api.CreateResponse[CommentGoodsUpdateHideStatusResponse](&CommentGoodsUpdateHideStatusResponse{})
	return
}

type CommentGoodsUpdateHideStatusRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
