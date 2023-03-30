package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreAddStoreAlbum
// @id 524
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增相册)
func (client *Client) StoreAddStoreAlbum(request *StoreAddStoreAlbumRequest) (response *StoreAddStoreAlbumResponse, err error) {
	rpcResponse := CreateStoreAddStoreAlbumResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreAddStoreAlbumRequest struct {
	*api.BaseRequest

	StoreId   int64  `json:"storeId,omitempty"`
	AlbumName string `json:"albumName,omitempty"`
	Status    bool   `json:"status,omitempty"`
}

type StoreAddStoreAlbumResponse struct {
}

func CreateStoreAddStoreAlbumRequest() (request *StoreAddStoreAlbumRequest) {
	request = &StoreAddStoreAlbumRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/addStoreAlbum", "api")
	request.Method = api.POST
	return
}

func CreateStoreAddStoreAlbumResponse() (response *api.BaseResponse[StoreAddStoreAlbumResponse]) {
	response = api.CreateResponse[StoreAddStoreAlbumResponse](&StoreAddStoreAlbumResponse{})
	return
}
