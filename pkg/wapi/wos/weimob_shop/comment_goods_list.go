package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsList
// @id 1777
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1777?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商品评论列表)
func (client *Client) CommentGoodsList(request *CommentGoodsListRequest) (response *CommentGoodsListResponse, err error) {
	rpcResponse := CreateCommentGoodsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsListRequest struct {
	*api.BaseRequest

	QueryParameter CommentGoodsListRequestQueryParameter `json:"queryParameter,omitempty"`
	BasicInfo      CommentGoodsListRequestBasicInfo      `json:"basicInfo,omitempty"`
	CurrentPage    int64                                 `json:"currentPage,omitempty"`
	PageSize       int64                                 `json:"pageSize,omitempty"`
}

type CommentGoodsListResponse struct {
	GoodsCommentVoList []CommentGoodsListResponseGoodsCommentVoList `json:"goodsCommentVoList,omitempty"`
	TotalNum           int64                                        `json:"totalNum,omitempty"`
	CurrentPage        int64                                        `json:"currentPage,omitempty"`
	PageSize           int64                                        `json:"pageSize,omitempty"`
	PageNum            int64                                        `json:"pageNum,omitempty"`
}

func CreateCommentGoodsListRequest() (request *CommentGoodsListRequest) {
	request = &CommentGoodsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/list", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsListResponse() (response *api.BaseResponse[CommentGoodsListResponse]) {
	response = api.CreateResponse[CommentGoodsListResponse](&CommentGoodsListResponse{})
	return
}

type CommentGoodsListRequestQueryParameter struct {
	CommentLevel         int64  `json:"commentLevel,omitempty"`
	StoreId              int64  `json:"storeId,omitempty"`
	IsHide               int64  `json:"isHide,omitempty"`
	IsReply              int64  `json:"isReply,omitempty"`
	SearchFile           string `json:"searchFile,omitempty"`
	SearchFileValue      string `json:"searchFileValue,omitempty"`
	NeedAutoCommentOrder int64  `json:"needAutoCommentOrder,omitempty"`
	CommentStartDate     string `json:"commentStartDate,omitempty"`
	CommentEndDate       int64  `json:"commentEndDate,omitempty"`
	CommentType          int64  `json:"commentType,omitempty"`
	CommentContextType   int64  `json:"commentContextType,omitempty"`
}

type CommentGoodsListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CommentGoodsListResponseGoodsCommentVoList struct {
	Id                  int64   `json:"id,omitempty"`
	Vid                 int64   `json:"vid,omitempty"`
	VidName             string  `json:"vidName,omitempty"`
	Wid                 int64   `json:"wid,omitempty"`
	OrderNo             int64   `json:"orderNo,omitempty"`
	UserImageUrl        string  `json:"userImageUrl,omitempty"`
	UserNickname        string  `json:"userNickname,omitempty"`
	ReceiverName        string  `json:"receiverName,omitempty"`
	ReceiverMobile      string  `json:"receiverMobile,omitempty"`
	IsAnonymity         int64   `json:"isAnonymity,omitempty"`
	GoodsId             int64   `json:"goodsId,omitempty"`
	GoodsTitle          string  `json:"goodsTitle,omitempty"`
	GoodsIndexImage     string  `json:"goodsIndexImage,omitempty"`
	SkuName             string  `json:"skuName,omitempty"`
	SkuNum              int64   `json:"skuNum,omitempty"`
	CommentScore        int64   `json:"commentScore,omitempty"`
	CommentLevel        int64   `json:"commentLevel,omitempty"`
	CommentContent      string  `json:"commentContent,omitempty"`
	CommentImageUrlList []int64 `json:"commentImageUrlList,omitempty"`
	ReplyImageUrlList   []int64 `json:"replyImageUrlList,omitempty"`
	CommentStatus       int64   `json:"commentStatus,omitempty"`
	CommentType         int64   `json:"commentType,omitempty"`
	IsHide              int64   `json:"isHide,omitempty"`
	IsSelected          int64   `json:"isSelected,omitempty"`
	CommentTime         string  `json:"commentTime,omitempty"`
}
