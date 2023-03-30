package weimob_guide

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuiderCustomerGetList
// @id 1386
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1386?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询导购的客户)
func (client *Client) GuiderCustomerGetList(request *GuiderCustomerGetListRequest) (response *GuiderCustomerGetListResponse, err error) {
	rpcResponse := CreateGuiderCustomerGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuiderCustomerGetListRequest struct {
	*api.BaseRequest

	QueryParameter GuiderCustomerGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                      `json:"pageNum,omitempty"`
	PageSize       int64                                      `json:"pageSize,omitempty"`
	BizVid         int64                                      `json:"bizVid,omitempty"`
	Tcode          string                                     `json:"tcode,omitempty"`
}

type GuiderCustomerGetListResponse struct {
	PageList   []GuiderCustomerGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                   `json:"pageNum,omitempty"`
	PageSize   int64                                   `json:"pageSize,omitempty"`
	TotalCount int64                                   `json:"totalCount,omitempty"`
}

func CreateGuiderCustomerGetListRequest() (request *GuiderCustomerGetListRequest) {
	request = &GuiderCustomerGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_guide", "v2.0", "guider/customer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateGuiderCustomerGetListResponse() (response *api.BaseResponse[GuiderCustomerGetListResponse]) {
	response = api.CreateResponse[GuiderCustomerGetListResponse](&GuiderCustomerGetListResponse{})
	return
}

type GuiderCustomerGetListRequestQueryParameter struct {
	GuiderWid int64 `json:"guiderWid,omitempty"`
}

type GuiderCustomerGetListResponsePageList struct {
	NickName          string `json:"nickName,omitempty"`
	HeadUrl           string `json:"headUrl,omitempty"`
	BindSceneTypeName string `json:"bindSceneTypeName,omitempty"`
	BindTime          string `json:"bindTime,omitempty"`
	Wid               int64  `json:"wid,omitempty"`
	CustomerPhone     string `json:"customerPhone,omitempty"`
}
