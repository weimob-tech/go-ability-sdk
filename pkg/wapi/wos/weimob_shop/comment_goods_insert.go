package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentGoodsInsert
// @id 1568
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1568?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建商品评论)
func (client *Client) CommentGoodsInsert(request *CommentGoodsInsertRequest) (response *CommentGoodsInsertResponse, err error) {
	rpcResponse := CreateCommentGoodsInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentGoodsInsertRequest struct {
	*api.BaseRequest

	GoodsCommentVoList    []CommentGoodsInsertRequestGoodsCommentVoList `json:"goodsCommentVoList,omitempty"`
	BasicInfo             CommentGoodsInsertRequestBasicInfo            `json:"basicInfo,omitempty"`
	ExtendInfo            CommentGoodsInsertRequestExtendInfo           `json:"extendInfo,omitempty"`
	BizWid                int64                                         `json:"bizWid,omitempty"`
	OrderNo               int64                                         `json:"orderNo,omitempty"`
	LogisticsServiceScore int64                                         `json:"logisticsServiceScore,omitempty"`
	ServiceAttitudeScore  int64                                         `json:"serviceAttitudeScore,omitempty"`
	PaymentAmount         int64                                         `json:"paymentAmount,omitempty"`
	OrderCreateTime       int64                                         `json:"orderCreateTime,omitempty"`
	CreateTime            int64                                         `json:"createTime,omitempty"`
	GoodsCount            int64                                         `json:"goodsCount,omitempty"`
	ThirdId               string                                        `json:"thirdId,omitempty"`
	ThirdSource           string                                        `json:"thirdSource,omitempty"`
}

type CommentGoodsInsertResponse struct {
	UpdateResult bool `json:"updateResult,omitempty"`
}

func CreateCommentGoodsInsertRequest() (request *CommentGoodsInsertRequest) {
	request = &CommentGoodsInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/goods/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentGoodsInsertResponse() (response *api.BaseResponse[CommentGoodsInsertResponse]) {
	response = api.CreateResponse[CommentGoodsInsertResponse](&CommentGoodsInsertResponse{})
	return
}

type CommentGoodsInsertRequestGoodsCommentVoList struct {
	ReceiverName        string  `json:"receiverName,omitempty"`
	ReceiverMobile      string  `json:"receiverMobile,omitempty"`
	IsAnonymity         string  `json:"isAnonymity,omitempty"`
	GoodsId             int64   `json:"goodsId,omitempty"`
	SkuId               int64   `json:"skuId,omitempty"`
	Price               int64   `json:"price,omitempty"`
	SkuNum              int64   `json:"skuNum,omitempty"`
	CommentScore        int64   `json:"commentScore,omitempty"`
	CommentImageUrlList []int64 `json:"commentImageUrlList,omitempty"`
	CommentVideoUrl     string  `json:"commentVideoUrl,omitempty"`
	ReplyContent        string  `json:"replyContent,omitempty"`
	ReplyContentTime    int64   `json:"replyContentTime,omitempty"`
	ReplyImageUrl       string  `json:"replyImageUrl,omitempty"`
	CommentType         int64   `json:"commentType,omitempty"`
	IsHide              int64   `json:"isHide,omitempty"`
	IsSelected          int64   `json:"isSelected,omitempty"`
	CommentTime         int64   `json:"commentTime,omitempty"`
	ThirdId             string  `json:"thirdId,omitempty"`
	ThirdSource         string  `json:"thirdSource,omitempty"`
	CommentContent      string  `json:"commentContent,omitempty"`
}

type CommentGoodsInsertRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CommentGoodsInsertRequestExtendInfo struct {
	UserIp string `json:"userIp,omitempty"`
}
