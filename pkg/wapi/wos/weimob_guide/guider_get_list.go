package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderGetList
// @id 1387
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1387?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购列表)
func (client *Client) GuiderGetList(request *GuiderGetListRequest) (response *GuiderGetListResponse, err error) {
	rpcResponse := CreateGuiderGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderGetListRequest struct {
	*api.BaseRequest

	QueryParameter GuiderGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                              `json:"pageNum,omitempty"`
	PageSize       int64                              `json:"pageSize,omitempty"`
}

type GuiderGetListResponse struct {
	PageList   []GuiderGetListResponsePageList `json:"pageList,omitempty"`
	PageSize   int64                           `json:"pageSize,omitempty"`
	TotalCount int64                           `json:"totalCount,omitempty"`
	PageNum    int64                           `json:"pageNum,omitempty"`
}

func CreateGuiderGetListRequest() (request *GuiderGetListRequest) {
	request = &GuiderGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderGetListResponse() (response *api.BaseResponse[GuiderGetListResponse]) {
	response = api.CreateResponse[GuiderGetListResponse](&GuiderGetListResponse{})
	return
}

type GuiderGetListRequestQueryParameter struct {
	IsUsed          int64   `json:"isUsed,omitempty"`
	GuiderVid       int64   `json:"guiderVid,omitempty"`
	GuiderWidList   []int64 `json:"guiderWidList,omitempty"`
	GuiderPhoneList []int64 `json:"guiderPhoneList,omitempty"`
	GuiderIdList    []int64 `json:"guiderIdList,omitempty"`
	GuiderName      string  `json:"guiderName,omitempty"`
}

type GuiderGetListResponsePageList struct {
	GuiderName  string `json:"guiderName,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
	VidName     string `json:"vidName,omitempty"`
	Zone        string `json:"zone,omitempty"`
	GuiderVid   string `json:"guiderVid,omitempty"`
}
