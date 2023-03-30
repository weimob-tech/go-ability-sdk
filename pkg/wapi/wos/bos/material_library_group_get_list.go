package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MaterialLibraryGroupGetList
// @id 1827
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1827?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询素材分组)
func (client *Client) MaterialLibraryGroupGetList(request *MaterialLibraryGroupGetListRequest) (response *MaterialLibraryGroupGetListResponse, err error) {
	rpcResponse := CreateMaterialLibraryGroupGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MaterialLibraryGroupGetListRequest struct {
	*api.BaseRequest

	BosId        int64  `json:"bosId,omitempty"`
	GroupKeyword string `json:"groupKeyword,omitempty"`
	MaterialType int64  `json:"materialType,omitempty"`
	Vid          int64  `json:"vid,omitempty"`
}

type MaterialLibraryGroupGetListResponse struct {
	GroupList     []MaterialLibraryGroupGetListResponseGroupList `json:"groupList,omitempty"`
	MaterialTotal int64                                          `json:"materialTotal,omitempty"`
	Total         int64                                          `json:"total,omitempty"`
}

func CreateMaterialLibraryGroupGetListRequest() (request *MaterialLibraryGroupGetListRequest) {
	request = &MaterialLibraryGroupGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "material/library/group/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateMaterialLibraryGroupGetListResponse() (response *api.BaseResponse[MaterialLibraryGroupGetListResponse]) {
	response = api.CreateResponse[MaterialLibraryGroupGetListResponse](&MaterialLibraryGroupGetListResponse{})
	return
}

type MaterialLibraryGroupGetListResponseGroupList struct {
	BusinessIdentity        MaterialLibraryGroupGetListResponseBusinessIdentity    `json:"businessIdentity,omitempty"`
	MaterialGroupList       []MaterialLibraryGroupGetListResponseMaterialGroupList `json:"materialGroupList,omitempty"`
	ParentGroup             MaterialLibraryGroupGetListResponseParentGroup2        `json:"parentGroup,omitempty"`
	BosId                   int64                                                  `json:"bosId,omitempty"`
	CreateTime              string                                                 `json:"createTime,omitempty"`
	Id                      int64                                                  `json:"id,omitempty"`
	Level                   int64                                                  `json:"level,omitempty"`
	MerchantId              int64                                                  `json:"merchantId,omitempty"`
	Name                    string                                                 `json:"name,omitempty"`
	ParentId                int64                                                  `json:"parentId,omitempty"`
	SortId                  int64                                                  `json:"sortId,omitempty"`
	SourceProductId         int64                                                  `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                                  `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                                  `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                                 `json:"tcode,omitempty"`
	Vid                     int64                                                  `json:"vid,omitempty"`
}

type MaterialLibraryGroupGetListResponseBusinessIdentity struct {
}

type MaterialLibraryGroupGetListResponseMaterialGroupList struct {
	BusinessIdentity        MaterialLibraryGroupGetListResponseBusinessIdentity2    `json:"businessIdentity,omitempty"`
	MaterialGroupList       []MaterialLibraryGroupGetListResponseMaterialGroupList2 `json:"materialGroupList,omitempty"`
	ParentGroup             MaterialLibraryGroupGetListResponseParentGroup          `json:"parentGroup,omitempty"`
	BosId                   int64                                                   `json:"bosId,omitempty"`
	CreateTime              string                                                  `json:"createTime,omitempty"`
	Id                      int64                                                   `json:"id,omitempty"`
	Level                   int64                                                   `json:"level,omitempty"`
	MerchantId              int64                                                   `json:"merchantId,omitempty"`
	Name                    string                                                  `json:"name,omitempty"`
	ParentId                int64                                                   `json:"parentId,omitempty"`
	SortId                  int64                                                   `json:"sortId,omitempty"`
	SourceProductId         int64                                                   `json:"sourceProductId,omitempty"`
	SourceProductInstanceId int64                                                   `json:"sourceProductInstanceId,omitempty"`
	SourceProductVersionId  int64                                                   `json:"sourceProductVersionId,omitempty"`
	Tcode                   string                                                  `json:"tcode,omitempty"`
	Vid                     int64                                                   `json:"vid,omitempty"`
}

type MaterialLibraryGroupGetListResponseBusinessIdentity2 struct {
}

type MaterialLibraryGroupGetListResponseMaterialGroupList2 struct {
}

type MaterialLibraryGroupGetListResponseParentGroup struct {
}

type MaterialLibraryGroupGetListResponseParentGroup2 struct {
}
