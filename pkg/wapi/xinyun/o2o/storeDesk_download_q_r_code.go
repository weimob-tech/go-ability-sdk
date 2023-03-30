package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreDeskDownloadQRCode
// @id 72
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=桌台下载二维码)
func (client *Client) StoreDeskDownloadQRCode(request *StoreDeskDownloadQRCodeRequest) (response *StoreDeskDownloadQRCodeResponse, err error) {
	rpcResponse := CreateStoreDeskDownloadQRCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreDeskDownloadQRCodeRequest struct {
	*api.BaseRequest

	DeskId int64 `json:"deskId,omitempty"`
}

type StoreDeskDownloadQRCodeResponse struct {
}

func CreateStoreDeskDownloadQRCodeRequest() (request *StoreDeskDownloadQRCodeRequest) {
	request = &StoreDeskDownloadQRCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "storeDesk/downloadQRCode", "api")
	request.Method = api.POST
	return
}

func CreateStoreDeskDownloadQRCodeResponse() (response *api.BaseResponse[StoreDeskDownloadQRCodeResponse]) {
	response = api.CreateResponse[StoreDeskDownloadQRCodeResponse](&StoreDeskDownloadQRCodeResponse{})
	return
}
