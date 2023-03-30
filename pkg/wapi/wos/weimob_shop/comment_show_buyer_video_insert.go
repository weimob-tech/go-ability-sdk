package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CommentShowBuyerVideoInsert
// @id 1786
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1786?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=买家秀上传视频)
func (client *Client) CommentShowBuyerVideoInsert(request *CommentShowBuyerVideoInsertRequest) (response *CommentShowBuyerVideoInsertResponse, err error) {
	rpcResponse := CreateCommentShowBuyerVideoInsertResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CommentShowBuyerVideoInsertRequest struct {
	*api.BaseRequest

	VideoUrl string `json:"videoUrl,omitempty"`
}

type CommentShowBuyerVideoInsertResponse struct {
	ReplyVideoUrl      string `json:"replyVideoUrl,omitempty"`
	ReplyVideoImageUrl string `json:"replyVideoImageUrl,omitempty"`
}

func CreateCommentShowBuyerVideoInsertRequest() (request *CommentShowBuyerVideoInsertRequest) {
	request = &CommentShowBuyerVideoInsertRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "comment/show/buyer/video/insert", "apigw")
	request.Method = api.POST
	return
}

func CreateCommentShowBuyerVideoInsertResponse() (response *api.BaseResponse[CommentShowBuyerVideoInsertResponse]) {
	response = api.CreateResponse[CommentShowBuyerVideoInsertResponse](&CommentShowBuyerVideoInsertResponse{})
	return
}
