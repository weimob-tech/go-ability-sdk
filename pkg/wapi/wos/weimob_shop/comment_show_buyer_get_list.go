package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerGetList
// @id 1781
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1781?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询买家秀列表)
func (client *Client) CommentShowBuyerGetList(request *CommentShowBuyerGetListRequest) (response *CommentShowBuyerGetListResponse, err error) {
	rpcResponse := CreateCommentShowBuyerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerGetListRequest struct {
	*api.BaseRequest

	BasicInfo            CommentShowBuyerGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	StartTime            string                                  `json:"startTime,omitempty"`
	EndTime              string                                  `json:"endTime,omitempty"`
	GroupId              int64                                   `json:"groupId,omitempty"`
	IsAudit              int64                                   `json:"isAudit,omitempty"`
	IsHide               int64                                   `json:"isHide,omitempty"`
	GroupType            int64                                   `json:"groupType,omitempty"`
	CurrentPage          int64                                   `json:"currentPage,omitempty"`
	PageSize             int64                                   `json:"pageSize,omitempty"`
	IsQueryPubInfoDetail bool                                    `json:"isQueryPubInfoDetail,omitempty"`
}

type CommentShowBuyerGetListResponse struct {
	BuyerShowDetailList []CommentShowBuyerGetListResponseBuyerShowDetailList `json:"buyerShowDetailList,omitempty"`
	CurrentPage         int64                                                `json:"currentPage,omitempty"`
	PageNum             int64                                                `json:"pageNum,omitempty"`
	PageSize            int64                                                `json:"pageSize,omitempty"`
	TotalNum            int64                                                `json:"totalNum,omitempty"`
}

func CreateCommentShowBuyerGetListRequest() (request *CommentShowBuyerGetListRequest) {
	request = &CommentShowBuyerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerGetListResponse() (response *api.BaseResponse[CommentShowBuyerGetListResponse]) {
	response = api.CreateResponse[CommentShowBuyerGetListResponse](&CommentShowBuyerGetListResponse{})
	return
}

type CommentShowBuyerGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type CommentShowBuyerGetListResponseBuyerShowDetailList struct {
	CommentInfo   CommentShowBuyerGetListResponseCommentInfo   `json:"commentInfo,omitempty"`
	RelateGoods   []CommentShowBuyerGetListResponseRelateGoods `json:"relateGoods,omitempty"`
	PubInfo       CommentShowBuyerGetListResponsePubInfo       `json:"pubInfo,omitempty"`
	AuditMsg      string                                       `json:"auditMsg,omitempty"`
	CommentCount  int64                                        `json:"commentCount,omitempty"`
	CreateTime    int64                                        `json:"createTime,omitempty"`
	GoodsCount    int64                                        `json:"goodsCount,omitempty"`
	Id            int64                                        `json:"id,omitempty"`
	IsAnonymity   int64                                        `json:"isAnonymity,omitempty"`
	IsAudit       int64                                        `json:"isAudit,omitempty"`
	IsAvailable   int64                                        `json:"isAvailable,omitempty"`
	IsHide        int64                                        `json:"isHide,omitempty"`
	LikeCount     int64                                        `json:"likeCount,omitempty"`
	OperateSource int64                                        `json:"operateSource,omitempty"`
	UserImageUrl  string                                       `json:"userImageUrl,omitempty"`
	UserNickname  string                                       `json:"userNickname,omitempty"`
	Vid           int64                                        `json:"vid,omitempty"`
	Wid           int64                                        `json:"wid,omitempty"`
	Sort          int64                                        `json:"sort,omitempty"`
}

type CommentShowBuyerGetListResponseCommentInfo struct {
	ImageContent []CommentShowBuyerGetListResponseImageContent `json:"imageContent,omitempty"`
	VideoContent []CommentShowBuyerGetListResponseVideoContent `json:"videoContent,omitempty"`
	TextContent  string                                        `json:"textContent,omitempty"`
}

type CommentShowBuyerGetListResponseImageContent struct {
	Height int64  `json:"height,omitempty"`
	Url    string `json:"url,omitempty"`
	Width  int64  `json:"width,omitempty"`
}

type CommentShowBuyerGetListResponseVideoContent struct {
	ImageUrl    string `json:"imageUrl,omitempty"`
	MediaHeight int64  `json:"mediaHeight,omitempty"`
	MediaName   string `json:"mediaName,omitempty"`
	MediaUrl    string `json:"mediaUrl,omitempty"`
	MediaWidth  int64  `json:"mediaWidth,omitempty"`
}

type CommentShowBuyerGetListResponseRelateGoods struct {
	GoodsId      int64  `json:"goodsId,omitempty"`
	GoodsPicture string `json:"goodsPicture,omitempty"`
	GoodsPrice   int64  `json:"goodsPrice,omitempty"`
	GoodsTitle   string `json:"goodsTitle,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
}

type CommentShowBuyerGetListResponsePubInfo struct {
	PubType      string  `json:"pubType,omitempty"`
	UserNickname string  `json:"userNickname,omitempty"`
	Wid          int64   `json:"wid,omitempty"`
	UnionId      string  `json:"unionId,omitempty"`
	Phones       []int64 `json:"phones,omitempty"`
	AppId        string  `json:"appId,omitempty"`
}
