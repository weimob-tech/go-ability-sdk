package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"
import "io"

// GoodsVideoUploadVideo
// @id 370
// @author WeimobCloud
// @create 2023-3-27
// @doc [](https://doc.weimobcloud.com/search?key=上传商品视频)
func (client *Client) GoodsVideoUploadVideo(request *GoodsVideoUploadVideoRequest) (response *GoodsVideoUploadVideoResponse, err error) {
	rpcResponse := CreateGoodsVideoUploadVideoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsVideoUploadVideoRequest struct {
	*api.BaseMultipartRequest

	Name string    `json:"name,omitempty"`
	File io.Reader `json:"file,omitempty"`
}

type GoodsVideoUploadVideoResponse struct {
	CoverPicUrl string `json:"coverPicUrl,omitempty"`
	UrlInfo     string `json:"urlInfo,omitempty"`
	Size        int64  `json:"size,omitempty"`
}

func CreateGoodsVideoUploadVideoRequest() (request *GoodsVideoUploadVideoRequest) {
	request = &GoodsVideoUploadVideoRequest{
		BaseMultipartRequest: &api.BaseMultipartRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsVideo/uploadVideo", "media")
	request.Method = api.POST
	return
}

func CreateGoodsVideoUploadVideoResponse() (response *api.BaseResponse[GoodsVideoUploadVideoResponse]) {
	response = api.CreateResponse[GoodsVideoUploadVideoResponse](&GoodsVideoUploadVideoResponse{})
	return
}
