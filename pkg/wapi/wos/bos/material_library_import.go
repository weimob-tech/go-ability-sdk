package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MaterialLibraryImport
// @id 1825
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1825?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=上传素材)
func (client *Client) MaterialLibraryImport(request *MaterialLibraryImportRequest) (response *MaterialLibraryImportResponse, err error) {
	rpcResponse := CreateMaterialLibraryImportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MaterialLibraryImportRequest struct {
	*api.BaseRequest

	Channels                []MaterialLibraryImportRequestChannels           `json:"channels,omitempty"`
	MaterialDetailExts      []MaterialLibraryImportRequestMaterialDetailExts `json:"materialDetailExts,omitempty"`
	BosId                   int64                                            `json:"bosId,omitempty"`
	CreateUid               int64                                            `json:"createUid,omitempty"`
	Duration                int64                                            `json:"duration,omitempty"`
	GroupId                 int64                                            `json:"groupId,omitempty"`
	Introduction            string                                           `json:"introduction,omitempty"`
	IsSync                  int64                                            `json:"isSync,omitempty"`
	MaterialFileType        int64                                            `json:"materialFileType,omitempty"`
	MaterialType            int64                                            `json:"materialType,omitempty"`
	MerchantId              int64                                            `json:"merchantId,omitempty"`
	Size                    int64                                            `json:"size,omitempty"`
	SourceProductId         int64                                            `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                            `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                            `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                           `json:"tcode,omitempty"`
	Title                   string                                           `json:"title,omitempty"`
	Url                     string                                           `json:"url,omitempty"`
	Vid                     int64                                            `json:"vid,omitempty"`
}

type MaterialLibraryImportResponse struct {
	MaterialId int64 `json:"materialId,omitempty"`
	Status     int64 `json:"status,omitempty"`
}

func CreateMaterialLibraryImportRequest() (request *MaterialLibraryImportRequest) {
	request = &MaterialLibraryImportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "material/library/import", "apigw")
	request.Method = api.POST
	return
}

func CreateMaterialLibraryImportResponse() (response *api.BaseResponse[MaterialLibraryImportResponse]) {
	response = api.CreateResponse[MaterialLibraryImportResponse](&MaterialLibraryImportResponse{})
	return
}

type MaterialLibraryImportRequestChannels struct {
	ChannelId      string `json:"channelId,omitempty"`
	ChannelSubType int64  `json:"channelSubType,omitempty"`
	ChannelType    int64  `json:"channelType,omitempty"`
	MediaId        string `json:"mediaId,omitempty"`
}

type MaterialLibraryImportRequestMaterialDetailExts struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
