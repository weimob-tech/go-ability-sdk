package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreAddAlbumPicture
// @id 522
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传酒店图片)
func (client *Client) StoreAddAlbumPicture(request *StoreAddAlbumPictureRequest) (response *StoreAddAlbumPictureResponse, err error) {
	rpcResponse := CreateStoreAddAlbumPictureResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreAddAlbumPictureRequest struct {
	*api.BaseRequest

	StoreId     int64  `json:"storeId,omitempty"`
	AlbumId     int64  `json:"albumId,omitempty"`
	PhotoName   string `json:"photoName,omitempty"`
	OrderNumber int    `json:"orderNumber,omitempty"`
	PhotoUrl    string `json:"photoUrl,omitempty"`
}

type StoreAddAlbumPictureResponse struct {
}

func CreateStoreAddAlbumPictureRequest() (request *StoreAddAlbumPictureRequest) {
	request = &StoreAddAlbumPictureRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/addAlbumPicture", "api")
	request.Method = api.POST
	return
}

func CreateStoreAddAlbumPictureResponse() (response *api.BaseResponse[StoreAddAlbumPictureResponse]) {
	response = api.CreateResponse[StoreAddAlbumPictureResponse](&StoreAddAlbumPictureResponse{})
	return
}
