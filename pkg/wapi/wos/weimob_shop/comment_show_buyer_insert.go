package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerInsert
// @id 1785
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1785?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=创建买家秀)
func (client *Client) CommentShowBuyerInsert(request *CommentShowBuyerInsertRequest) (response *CommentShowBuyerInsertResponse, err error) {
	rpcResponse := CreateCommentShowBuyerInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerInsertRequest struct {
	*api.BaseRequest

	CommentInfo     CommentShowBuyerInsertRequestCommentInfo   `json:"commentInfo,omitempty"`
	RelateGoods     []CommentShowBuyerInsertRequestRelateGoods `json:"relateGoods,omitempty"`
	IsShow          int64                                      `json:"isShow,omitempty"`
	UserImageUrl    string                                     `json:"userImageUrl,omitempty"`
	UserNickname    string                                     `json:"userNickname,omitempty"`
	BuyerShowGroups []int64                                    `json:"buyerShowGroups,omitempty"`
}

type CommentShowBuyerInsertResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateCommentShowBuyerInsertRequest() (request *CommentShowBuyerInsertRequest) {
	request = &CommentShowBuyerInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerInsertResponse() (response *api.BaseResponse[CommentShowBuyerInsertResponse]) {
	response = api.CreateResponse[CommentShowBuyerInsertResponse](&CommentShowBuyerInsertResponse{})
	return
}

type CommentShowBuyerInsertRequestCommentInfo struct {
	ImageContent []CommentShowBuyerInsertRequestImageContent `json:"imageContent,omitempty"`
	VideoContent []CommentShowBuyerInsertRequestVideoContent `json:"videoContent,omitempty"`
	TextContent  string                                      `json:"textContent,omitempty"`
}

type CommentShowBuyerInsertRequestImageContent struct {
	Url    string `json:"url,omitempty"`
	Width  int64  `json:"width,omitempty"`
	Height int64  `json:"height,omitempty"`
}

type CommentShowBuyerInsertRequestVideoContent struct {
	ImageUrl    string `json:"imageUrl,omitempty"`
	MediaUrl    string `json:"mediaUrl,omitempty"`
	MediaWidth  int64  `json:"mediaWidth,omitempty"`
	MediaHeight int64  `json:"mediaHeight,omitempty"`
	MediaName   string `json:"mediaName,omitempty"`
}

type CommentShowBuyerInsertRequestRelateGoods struct {
	GoodsId int64 `json:"goodsId,omitempty"`
	BosId   int64 `json:"bosId,omitempty"`
	Vid     int64 `json:"vid,omitempty"`
}
