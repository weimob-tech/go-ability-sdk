package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MaterialLibraryUpdate
// @id 1824
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1824?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=编辑素材)
func (client *Client) MaterialLibraryUpdate(request *MaterialLibraryUpdateRequest) (response *MaterialLibraryUpdateResponse, err error) {
	rpcResponse := CreateMaterialLibraryUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MaterialLibraryUpdateRequest struct {
	*api.BaseRequest

	Channels                []MaterialLibraryUpdateRequestChannels           `json:"channels,omitempty"`
	MaterialDetailExts      []MaterialLibraryUpdateRequestMaterialDetailExts `json:"materialDetailExts,omitempty"`
	BosId                   int64                                            `json:"bosId,omitempty"`
	GroupId                 int64                                            `json:"groupId,omitempty"`
	Id                      int64                                            `json:"id,omitempty"`
	Introduction            string                                           `json:"introduction,omitempty"`
	MaterialType            int64                                            `json:"materialType,omitempty"`
	MerchantId              int64                                            `json:"merchantId,omitempty"`
	SourceProductId         int64                                            `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                            `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                            `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                           `json:"tcode,omitempty"`
	Title                   string                                           `json:"title,omitempty"`
	UpdateUid               int64                                            `json:"updateUid,omitempty"`
	Vid                     int64                                            `json:"vid,omitempty"`
}

type MaterialLibraryUpdateResponse struct {
	MaterialId int64 `json:"materialId,omitempty"`
	Status     int64 `json:"status,omitempty"`
}

func CreateMaterialLibraryUpdateRequest() (request *MaterialLibraryUpdateRequest) {
	request = &MaterialLibraryUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "material/library/update", "apigw")
	request.Method = api.POST
	return
}

func CreateMaterialLibraryUpdateResponse() (response *api.BaseResponse[MaterialLibraryUpdateResponse]) {
	response = api.CreateResponse[MaterialLibraryUpdateResponse](&MaterialLibraryUpdateResponse{})
	return
}

type MaterialLibraryUpdateRequestChannels struct {
	ChannelId      string `json:"channelId,omitempty"`
	ChannelSubType int64  `json:"channelSubType,omitempty"`
	ChannelType    int64  `json:"channelType,omitempty"`
	MediaId        string `json:"mediaId,omitempty"`
}

type MaterialLibraryUpdateRequestMaterialDetailExts struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
