package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointTransferGetList
// @id 1901
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1901?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户积分日志)
func (client *Client) PointTransferGetList(request *PointTransferGetListRequest) (response *PointTransferGetListResponse, err error) {
	rpcResponse := CreatePointTransferGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointTransferGetListRequest struct {
	*api.BaseRequest

	QueryParameter PointTransferGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	Vid            int64                                     `json:"vid,omitempty"`
	PointPlanId    int64                                     `json:"pointPlanId,omitempty"`
	PageNum        int64                                     `json:"pageNum,omitempty"`
	PageSize       int64                                     `json:"pageSize,omitempty"`
	Wid            int64                                     `json:"wid,omitempty"`
}

type PointTransferGetListResponse struct {
	DataList   []PointTransferGetListResponseDataList `json:"dataList,omitempty"`
	PageNum    int64                                  `json:"pageNum,omitempty"`
	PageSize   int64                                  `json:"pageSize,omitempty"`
	TotalCount int64                                  `json:"totalCount,omitempty"`
}

func CreatePointTransferGetListRequest() (request *PointTransferGetListRequest) {
	request = &PointTransferGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/transfer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreatePointTransferGetListResponse() (response *api.BaseResponse[PointTransferGetListResponse]) {
	response = api.CreateResponse[PointTransferGetListResponse](&PointTransferGetListResponse{})
	return
}

type PointTransferGetListRequestQueryParameter struct {
	IdentityType    int64   `json:"identityType,omitempty"`
	IdentityNo      string  `json:"identityNo,omitempty"`
	TransBeginDate  string  `json:"transBeginDate,omitempty"`
	TransEndDate    string  `json:"transEndDate,omitempty"`
	StoreVidList    []int64 `json:"storeVidList,omitempty"`
	ChannelTypeList []int64 `json:"channelTypeList,omitempty"`
	ChangeTypeList  []int64 `json:"changeTypeList,omitempty"`
	OperateName     string  `json:"operateName,omitempty"`
	TransNo         string  `json:"transNo,omitempty"`
}

type PointTransferGetListResponseDataList struct {
	TransNo         string `json:"transNo,omitempty"`
	TransDate       string `json:"transDate,omitempty"`
	Uid             string `json:"uid,omitempty"`
	ChangeType      string `json:"changeType,omitempty"`
	ChangeTypeDesc  string `json:"changeTypeDesc,omitempty"`
	RuleName        string `json:"ruleName,omitempty"`
	Customer        string `json:"customer,omitempty"`
	VariationRange  string `json:"variationRange,omitempty"`
	Amount          string `json:"amount,omitempty"`
	ChannelType     string `json:"channelType,omitempty"`
	ChannelTypeDesc string `json:"channelTypeDesc,omitempty"`
	OccurVName      string `json:"occurVName,omitempty"`
	Remark          string `json:"remark,omitempty"`
	OperateWName    string `json:"operateWName,omitempty"`
	OccurVid        int64  `json:"occurVid,omitempty"`
}
