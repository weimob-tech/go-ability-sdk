package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerGet
// @id 1779
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1779?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询买家秀详情)
func (client *Client) CommentShowBuyerGet(request *CommentShowBuyerGetRequest) (response *CommentShowBuyerGetResponse, err error) {
	rpcResponse := CreateCommentShowBuyerGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerGetRequest struct {
	*api.BaseRequest

	Id int64 `json:"id,omitempty"`
}

type CommentShowBuyerGetResponse struct {
	CommentInfo   CommentShowBuyerGetResponseCommentInfo   `json:"commentInfo,omitempty"`
	RelateGoods   []CommentShowBuyerGetResponseRelateGoods `json:"relateGoods,omitempty"`
	Groups        []CommentShowBuyerGetResponseGroups      `json:"groups,omitempty"`
	GoodsCount    int64                                    `json:"goodsCount,omitempty"`
	AuditMsg      string                                   `json:"auditMsg,omitempty"`
	LikeCount     int64                                    `json:"likeCount,omitempty"`
	IsAvailable   int64                                    `json:"isAvailable,omitempty"`
	Wid           int64                                    `json:"wid,omitempty"`
	OperateSource int64                                    `json:"operateSource,omitempty"`
	UserImageUrl  string                                   `json:"userImageUrl,omitempty"`
	UserNickname  string                                   `json:"userNickname,omitempty"`
	IsAnonymity   int64                                    `json:"isAnonymity,omitempty"`
	IsAudit       int64                                    `json:"isAudit,omitempty"`
	IsHide        int64                                    `json:"isHide,omitempty"`
	CommentCount  int64                                    `json:"commentCount,omitempty"`
	CreateTime    string                                   `json:"createTime,omitempty"`
	Id            int64                                    `json:"id,omitempty"`
	Sort          int64                                    `json:"sort,omitempty"`
}

func CreateCommentShowBuyerGetRequest() (request *CommentShowBuyerGetRequest) {
	request = &CommentShowBuyerGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerGetResponse() (response *api.BaseResponse[CommentShowBuyerGetResponse]) {
	response = api.CreateResponse[CommentShowBuyerGetResponse](&CommentShowBuyerGetResponse{})
	return
}

type CommentShowBuyerGetResponseCommentInfo struct {
	ImageContent []CommentShowBuyerGetResponseImageContent `json:"imageContent,omitempty"`
	VideoContent []CommentShowBuyerGetResponseVideoContent `json:"videoContent,omitempty"`
	TextContent  string                                    `json:"textContent,omitempty"`
}

type CommentShowBuyerGetResponseImageContent struct {
	Url    string `json:"url,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
}

type CommentShowBuyerGetResponseVideoContent struct {
	ImageUrl    string `json:"imageUrl,omitempty"`
	MediaUrl    string `json:"mediaUrl,omitempty"`
	MediaWidth  int64  `json:"mediaWidth,omitempty"`
	MediaHeight int64  `json:"mediaHeight,omitempty"`
	MediaName   string `json:"mediaName,omitempty"`
}

type CommentShowBuyerGetResponseRelateGoods struct {
	ActivityTags []CommentShowBuyerGetResponseActivityTags `json:"activityTags,omitempty"`
	GoodsId      int64                                     `json:"goodsId,omitempty"`
	GoodsPicture string                                    `json:"goodsPicture,omitempty"`
	GoodsTitle   string                                    `json:"goodsTitle,omitempty"`
	GoodsPrice   int64                                     `json:"goodsPrice,omitempty"`
	OrgPrice     int64                                     `json:"orgPrice,omitempty"`
	Vid          int64                                     `json:"vid,omitempty"`
}

type CommentShowBuyerGetResponseActivityTags struct {
	ActivityType int64  `json:"activityType,omitempty"`
	TagText      string `json:"tagText,omitempty"`
	Color        string `json:"color,omitempty"`
}

type CommentShowBuyerGetResponseGroups struct {
	GoodsGroups []CommentShowBuyerGetResponseGoodsGroups `json:"goodsGroups,omitempty"`
	GroupName   string                                   `json:"groupName,omitempty"`
	GroupType   int64                                    `json:"groupType,omitempty"`
	Id          int64                                    `json:"id,omitempty"`
}

type CommentShowBuyerGetResponseGoodsGroups struct {
	GroupId   int64  `json:"groupId,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ParentId  int64  `json:"parentId,omitempty"`
}
