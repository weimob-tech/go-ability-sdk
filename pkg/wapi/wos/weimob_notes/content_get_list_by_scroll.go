package weimob_notes

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ContentGetListByScroll
// @id 1799
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1799?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=滚动查询内容列表)
func (client *Client) ContentGetListByScroll(request *ContentGetListByScrollRequest) (response *ContentGetListByScrollResponse, err error) {
	rpcResponse := CreateContentGetListByScrollResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ContentGetListByScrollRequest struct {
	*api.BaseRequest

	StartId  int64 `json:"startId,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
	Vid      int64 `json:"vid,omitempty"`
}

type ContentGetListByScrollResponse struct {
	PageList   ContentGetListByScrollResponsePageList `json:"pageList,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

func CreateContentGetListByScrollRequest() (request *ContentGetListByScrollRequest) {
	request = &ContentGetListByScrollRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_notes", "v2.0", "content/getListByScroll", "apigw")
	request.Method = api.POST
	return
}

func CreateContentGetListByScrollResponse() (response *api.BaseResponse[ContentGetListByScrollResponse]) {
	response = api.CreateResponse[ContentGetListByScrollResponse](&ContentGetListByScrollResponse{})
	return
}

type ContentGetListByScrollResponsePageList struct {
	BosId              string  `json:"bosId,omitempty"`
	GroupNameList      []int64 `json:"groupNameList,omitempty"`
	ContentId          string  `json:"contentId,omitempty"`
	CoverImg           string  `json:"coverImg,omitempty"`
	CreateTime         string  `json:"createTime,omitempty"`
	CreatorId          string  `json:"creatorId,omitempty"`
	CreatorName        string  `json:"creatorName,omitempty"`
	CreatorPortraitPic string  `json:"creatorPortraitPic,omitempty"`
	ProductId          string  `json:"productId,omitempty"`
	PublicTime         string  `json:"publicTime,omitempty"`
	Title              string  `json:"title,omitempty"`
	Vid                string  `json:"vid,omitempty"`
	VidName            string  `json:"vidName,omitempty"`
}
