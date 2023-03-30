package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsReply
// @id 1782
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1782?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家回复商品评论)
func (client *Client) CommentGoodsReply(request *CommentGoodsReplyRequest) (response *CommentGoodsReplyResponse, err error) {
	rpcResponse := CreateCommentGoodsReplyResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsReplyRequest struct {
	*api.BaseRequest

	BasicInfo         CommentGoodsReplyRequestBasicInfo  `json:"basicInfo,omitempty"`
	ExtendInfo        CommentGoodsReplyRequestExtendInfo `json:"extendInfo,omitempty"`
	CommentId         int64                              `json:"commentId,omitempty"`
	ReplyContent      string                             `json:"replyContent,omitempty"`
	ReplyImageUrlList []int64                            `json:"replyImageUrlList,omitempty"`
}

type CommentGoodsReplyResponse struct {
	UpdateResult bool `json:"updateResult,omitempty"`
}

func CreateCommentGoodsReplyRequest() (request *CommentGoodsReplyRequest) {
	request = &CommentGoodsReplyRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/reply", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsReplyResponse() (response *api.BaseResponse[CommentGoodsReplyResponse]) {
	response = api.CreateResponse[CommentGoodsReplyResponse](&CommentGoodsReplyResponse{})
	return
}

type CommentGoodsReplyRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CommentGoodsReplyRequestExtendInfo struct {
	UserIp string `json:"userIp,omitempty"`
}
