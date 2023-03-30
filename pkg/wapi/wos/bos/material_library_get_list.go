package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MaterialLibraryGetList
// @id 1826
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1826?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询素材)
func (client *Client) MaterialLibraryGetList(request *MaterialLibraryGetListRequest) (response *MaterialLibraryGetListResponse, err error) {
	rpcResponse := CreateMaterialLibraryGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MaterialLibraryGetListRequest struct {
	*api.BaseRequest

	SortList                []MaterialLibraryGetListRequestSortList `json:"sortList,omitempty"`
	BosId                   int64                                   `json:"bosId,omitempty"`
	ChannelId               string                                  `json:"channelId,omitempty"`
	ChannelType             int64                                   `json:"channelType,omitempty"`
	Duration                int64                                   `json:"duration,omitempty"`
	GroupIds                []int64                                 `json:"groupIds,omitempty"`
	Keyword                 string                                  `json:"keyword,omitempty"`
	MaterialFileType        []int64                                 `json:"materialFileType,omitempty"`
	MaterialType            int64                                   `json:"materialType,omitempty"`
	MerchantId              int64                                   `json:"merchantId,omitempty"`
	PageNum                 int64                                   `json:"pageNum,omitempty"`
	PageSize                int64                                   `json:"pageSize,omitempty"`
	Size                    int64                                   `json:"size,omitempty"`
	SourceProductId         int64                                   `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                   `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                   `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                  `json:"tcode,omitempty"`
	Vid                     int64                                   `json:"vid,omitempty"`
}

type MaterialLibraryGetListResponse struct {
	Data       MaterialLibraryGetListResponseData       `json:"data,omitempty"`
	PageList   []MaterialLibraryGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                    `json:"pageNum,omitempty"`
	PageSize   int64                                    `json:"pageSize,omitempty"`
	TotalCount int64                                    `json:"totalCount,omitempty"`
}

func CreateMaterialLibraryGetListRequest() (request *MaterialLibraryGetListRequest) {
	request = &MaterialLibraryGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "material/library/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMaterialLibraryGetListResponse() (response *api.BaseResponse[MaterialLibraryGetListResponse]) {
	response = api.CreateResponse[MaterialLibraryGetListResponse](&MaterialLibraryGetListResponse{})
	return
}

type MaterialLibraryGetListRequestSortList struct {
	Field string `json:"field,omitempty"`
	Order string `json:"order,omitempty"`
}

type MaterialLibraryGetListResponseData struct {
}

type MaterialLibraryGetListResponsePageList struct {
	BusinessIdentity        MaterialLibraryGetListResponseBusinessIdentity     `json:"businessIdentity,omitempty"`
	Channels                []MaterialLibraryGetListResponseChannels           `json:"channels,omitempty"`
	MaterialDetailExts      []MaterialLibraryGetListResponseMaterialDetailExts `json:"materialDetailExts,omitempty"`
	BosId                   int64                                              `json:"bosId,omitempty"`
	Cover                   int64                                              `json:"cover,omitempty"`
	CreateTime              string                                             `json:"createTime,omitempty"`
	Duration                int64                                              `json:"duration,omitempty"`
	GroupId                 int64                                              `json:"groupId,omitempty"`
	Id                      int64                                              `json:"id,omitempty"`
	Introduction            string                                             `json:"introduction,omitempty"`
	LegalStatus             int64                                              `json:"legalStatus,omitempty"`
	MaterialFileType        int64                                              `json:"materialFileType,omitempty"`
	MaterialType            int64                                              `json:"materialType,omitempty"`
	MerchantId              int64                                              `json:"merchantId,omitempty"`
	Size                    int64                                              `json:"size,omitempty"`
	SourceProductId         int64                                              `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                              `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                              `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                             `json:"tcode,omitempty"`
	Title                   string                                             `json:"title,omitempty"`
	UpdateTime              string                                             `json:"updateTime,omitempty"`
	Url                     string                                             `json:"url,omitempty"`
	Vid                     int64                                              `json:"vid,omitempty"`
}

type MaterialLibraryGetListResponseBusinessIdentity struct {
}

type MaterialLibraryGetListResponseChannels struct {
	ChannelId      string  `json:"channelId,omitempty"`
	ChannelSubType int64   `json:"channelSubType,omitempty"`
	ChannelType    int64   `json:"channelType,omitempty"`
	ChannelUrls    []int64 `json:"channelUrls,omitempty"`
	MediaId        string  `json:"mediaId,omitempty"`
	SyncMsg        string  `json:"syncMsg,omitempty"`
	SyncStatus     int64   `json:"syncStatus,omitempty"`
}

type MaterialLibraryGetListResponseMaterialDetailExts struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
