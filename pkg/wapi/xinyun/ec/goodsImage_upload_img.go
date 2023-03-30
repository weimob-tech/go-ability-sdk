package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"
import "io"

// GoodsImageUploadImg
// @id 369
// @author WeimobCloud
// @create 2023-3-27
// @doc [](https://doc.weimobcloud.com/search?key=上传商品图片信息（文件上传）)
func (client *Client) GoodsImageUploadImg(request *GoodsImageUploadImgRequest) (response *GoodsImageUploadImgResponse, err error) {
	rpcResponse := CreateGoodsImageUploadImgResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsImageUploadImgRequest struct {
	*api.BaseMultipartRequest

	Name string    `json:"name,omitempty"`
	File io.Reader `json:"file,omitempty"`
}

type GoodsImageUploadImgResponse struct {
	CoverPicUrl string                     `json:"coverPicUrl,omitempty"`
	Size        int64                      `json:"size,omitempty"`
	UrlInfo     GoodsImageUploadImgUrlInfo `json:"urlInfo,omitempty"`
}

type GoodsImageUploadImgUrlInfo struct {
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
	LegalStatus int    `json:"legalStatus,omitempty"`
}

func CreateGoodsImageUploadImgRequest() (request *GoodsImageUploadImgRequest) {
	request = &GoodsImageUploadImgRequest{
		BaseMultipartRequest: &api.BaseMultipartRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsImage/uploadImg", "media")
	request.Method = api.POST
	return
}

func CreateGoodsImageUploadImgResponse() (response *api.BaseResponse[GoodsImageUploadImgResponse]) {
	response = api.CreateResponse[GoodsImageUploadImgResponse](&GoodsImageUploadImgResponse{})
	return
}
