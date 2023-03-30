package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreGetStoreAlbumList
// @id 521
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=酒店相册列表)
func (client *Client) StoreGetStoreAlbumList(request *StoreGetStoreAlbumListRequest) (response *StoreGetStoreAlbumListResponse, err error) {
	rpcResponse := CreateStoreGetStoreAlbumListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreGetStoreAlbumListRequest struct {
	*api.BaseRequest

	StoreId   int64 `json:"storeId,omitempty"`
	PageIndex int   `json:"pageIndex,omitempty"`
	PageSize  int   `json:"pageSize,omitempty"`
}

type StoreGetStoreAlbumListResponse struct {
}

func CreateStoreGetStoreAlbumListRequest() (request *StoreGetStoreAlbumListRequest) {
	request = &StoreGetStoreAlbumListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "store/getStoreAlbumList", "api")
	request.Method = api.POST
	return
}

func CreateStoreGetStoreAlbumListResponse() (response *api.BaseResponse[StoreGetStoreAlbumListResponse]) {
	response = api.CreateResponse[StoreGetStoreAlbumListResponse](&StoreGetStoreAlbumListResponse{})
	return
}
