package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsGet
// @id 1780
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1780?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=根据商品评论ID查询评论详情)
func (client *Client) CommentGoodsGet(request *CommentGoodsGetRequest) (response *CommentGoodsGetResponse, err error) {
	rpcResponse := CreateCommentGoodsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsGetRequest struct {
	*api.BaseRequest

	QueryParameter CommentGoodsGetRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      CommentGoodsGetRequestBasicInfo      `json:"basicInfo,omitempty"`
}

type CommentGoodsGetResponse struct {
	GoodsCommentVoList []CommentGoodsGetResponseGoodsCommentVoList `json:"goodsCommentVoList,omitempty"`
}

func CreateCommentGoodsGetRequest() (request *CommentGoodsGetRequest) {
	request = &CommentGoodsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsGetResponse() (response *api.BaseResponse[CommentGoodsGetResponse]) {
	response = api.CreateResponse[CommentGoodsGetResponse](&CommentGoodsGetResponse{})
	return
}

type CommentGoodsGetRequestQueryParameter struct {
	CommentIds []int64 `json:"commentIds,omitempty"`
}

type CommentGoodsGetRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CommentGoodsGetResponseGoodsCommentVoList struct {
	Id                   int64   `json:"id,omitempty"`
	OrderNo              int64   `json:"orderNo,omitempty"`
	UserImageUrl         string  `json:"userImageUrl,omitempty"`
	UserNickname         string  `json:"userNickname,omitempty"`
	ReceiverName         string  `json:"receiverName,omitempty"`
	ReceiverMobile       string  `json:"receiverMobile,omitempty"`
	IsAnonymity          int64   `json:"isAnonymity,omitempty"`
	GoodsId              int64   `json:"goodsId,omitempty"`
	GoodsTitle           string  `json:"goodsTitle,omitempty"`
	GoodsIndexImage      string  `json:"goodsIndexImage,omitempty"`
	SkuName              string  `json:"skuName,omitempty"`
	Price                int64   `json:"price,omitempty"`
	SkuNum               int64   `json:"skuNum,omitempty"`
	CommentScore         int64   `json:"commentScore,omitempty"`
	CommentLevel         int64   `json:"commentLevel,omitempty"`
	CommentContent       string  `json:"commentContent,omitempty"`
	CommentImageUrlList  []int64 `json:"commentImageUrlList,omitempty"`
	CommentVideoUrl      string  `json:"commentVideoUrl,omitempty"`
	CommentVideoImageUrl string  `json:"commentVideoImageUrl,omitempty"`
	ReplyContent         string  `json:"replyContent,omitempty"`
	ReplyImageUrlList    []int64 `json:"replyImageUrlList,omitempty"`
	IsRights             int64   `json:"isRights,omitempty"`
	IsHide               int64   `json:"isHide,omitempty"`
	IsSelected           int64   `json:"isSelected,omitempty"`
	CommentTime          string  `json:"commentTime,omitempty"`
}
