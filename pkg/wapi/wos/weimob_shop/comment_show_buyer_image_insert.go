package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerImageInsert
// @id 1787
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1787?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=买家秀上传图片)
func (client *Client) CommentShowBuyerImageInsert(request *CommentShowBuyerImageInsertRequest) (response *CommentShowBuyerImageInsertResponse, err error) {
	rpcResponse := CreateCommentShowBuyerImageInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerImageInsertRequest struct {
	*api.BaseRequest

	ImageUrl string `json:"imageUrl,omitempty"`
}

type CommentShowBuyerImageInsertResponse struct {
	ReplyImageUrl string `json:"replyImageUrl,omitempty"`
}

func CreateCommentShowBuyerImageInsertRequest() (request *CommentShowBuyerImageInsertRequest) {
	request = &CommentShowBuyerImageInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/image/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerImageInsertResponse() (response *api.BaseResponse[CommentShowBuyerImageInsertResponse]) {
	response = api.CreateResponse[CommentShowBuyerImageInsertResponse](&CommentShowBuyerImageInsertResponse{})
	return
}
