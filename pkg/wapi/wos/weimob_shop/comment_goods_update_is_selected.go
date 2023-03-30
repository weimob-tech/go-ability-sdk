package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsUpdateIsSelected
// @id 1784
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1784?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品评论精选状态)
func (client *Client) CommentGoodsUpdateIsSelected(request *CommentGoodsUpdateIsSelectedRequest) (response *CommentGoodsUpdateIsSelectedResponse, err error) {
	rpcResponse := CreateCommentGoodsUpdateIsSelectedResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsUpdateIsSelectedRequest struct {
	*api.BaseRequest

	BasicInfo  CommentGoodsUpdateIsSelectedRequestBasicInfo `json:"basicInfo,omitempty"`
	CommentIds []int64                                      `json:"commentIds,omitempty"`
	IsSelected int64                                        `json:"isSelected,omitempty"`
}

type CommentGoodsUpdateIsSelectedResponse struct {
	UpdateResult bool `json:"updateResult,omitempty"`
}

func CreateCommentGoodsUpdateIsSelectedRequest() (request *CommentGoodsUpdateIsSelectedRequest) {
	request = &CommentGoodsUpdateIsSelectedRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/updateIsSelected", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsUpdateIsSelectedResponse() (response *api.BaseResponse[CommentGoodsUpdateIsSelectedResponse]) {
	response = api.CreateResponse[CommentGoodsUpdateIsSelectedResponse](&CommentGoodsUpdateIsSelectedResponse{})
	return
}

type CommentGoodsUpdateIsSelectedRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
