package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// QrCodeBatchDownloadQRCode
// @id 2098
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量下载二维码)
func (client *Client) QrCodeBatchDownloadQRCode(request *QrCodeBatchDownloadQRCodeRequest) (response *QrCodeBatchDownloadQRCodeResponse, err error) {
	rpcResponse := CreateQrCodeBatchDownloadQRCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type QrCodeBatchDownloadQRCodeRequest struct {
	*api.BaseRequest

	BatchGenerateList  []QrCodeBatchDownloadQRCodeRequestBatchGenerateList `json:"batchGenerateList,omitempty"`
	Pid                int                                                 `json:"pid,omitempty"`
	SceneId            int                                                 `json:"sceneId,omitempty"`
	FileName           string                                              `json:"fileName,omitempty"`
	NotifyTemplateCode int                                                 `json:"notifyTemplateCode,omitempty"`
	OperatorId         int                                                 `json:"operatorId,omitempty"`
	RequestId          string                                              `json:"requestId,omitempty"`
}

type QrCodeBatchDownloadQRCodeResponse struct {
}

func CreateQrCodeBatchDownloadQRCodeRequest() (request *QrCodeBatchDownloadQRCodeRequest) {
	request = &QrCodeBatchDownloadQRCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "qrCode/batchDownloadQRCode", "api")
	request.Method = api.POST
	return
}

func CreateQrCodeBatchDownloadQRCodeResponse() (response *api.BaseResponse[QrCodeBatchDownloadQRCodeResponse]) {
	response = api.CreateResponse[QrCodeBatchDownloadQRCodeResponse](&QrCodeBatchDownloadQRCodeResponse{})
	return
}

type QrCodeBatchDownloadQRCodeRequestBatchGenerateList struct {
	GenerateDataList []QrCodeBatchDownloadQRCodeRequestGenerateDataList `json:"generateDataList,omitempty"`
	BizId            string                                             `json:"bizId,omitempty"`
}

type QrCodeBatchDownloadQRCodeRequestGenerateDataList struct {
	Scenes    []QrCodeBatchDownloadQRCodeRequestScenes `json:"scenes,omitempty"`
	RequestId []int                                    `json:"requestId,omitempty"`
	Type      int                                      `json:"type,omitempty"`
	Path      string                                   `json:"path,omitempty"`
	AppId     string                                   `json:"appId,omitempty"`
	Name      string                                   `json:"name,omitempty"`
	FileName  string                                   `json:"fileName,omitempty"`
	Width     int                                      `json:"width,omitempty"`
}

type QrCodeBatchDownloadQRCodeRequestScenes struct {
	Key QrCodeBatchDownloadQRCodeRequestKey `json:"key,omitempty"`
}

type QrCodeBatchDownloadQRCodeRequestKey struct {
}
