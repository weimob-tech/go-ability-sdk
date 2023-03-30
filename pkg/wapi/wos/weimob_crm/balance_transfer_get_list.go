package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// BalanceTransferGetList
// @id 1902
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1902?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询用户余额日志)
func (client *Client) BalanceTransferGetList(request *BalanceTransferGetListRequest) (response *BalanceTransferGetListResponse, err error) {
	rpcResponse := CreateBalanceTransferGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type BalanceTransferGetListRequest struct {
	*api.BaseRequest

	QueryParameter BalanceTransferGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	Vid            int64                                       `json:"vid,omitempty"`
	PageNum        int64                                       `json:"pageNum,omitempty"`
	PageSize       int64                                       `json:"pageSize,omitempty"`
	Wid            int64                                       `json:"wid,omitempty"`
	BalancePlanId  int64                                       `json:"balancePlanId,omitempty"`
}

type BalanceTransferGetListResponse struct {
	DataList   []BalanceTransferGetListResponseDataList `json:"dataList,omitempty"`
	PageNum    int64                                    `json:"pageNum,omitempty"`
	PageSize   int64                                    `json:"pageSize,omitempty"`
	TotalCount int64                                    `json:"totalCount,omitempty"`
}

func CreateBalanceTransferGetListRequest() (request *BalanceTransferGetListRequest) {
	request = &BalanceTransferGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "balance/transfer/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateBalanceTransferGetListResponse() (response *api.BaseResponse[BalanceTransferGetListResponse]) {
	response = api.CreateResponse[BalanceTransferGetListResponse](&BalanceTransferGetListResponse{})
	return
}

type BalanceTransferGetListRequestQueryParameter struct {
	IdentityType   int64   `json:"identityType,omitempty"`
	IdentityNo     string  `json:"identityNo,omitempty"`
	TransBeginDate int64   `json:"transBeginDate,omitempty"`
	TransEndDate   int64   `json:"transEndDate,omitempty"`
	OperateName    string  `json:"operateName,omitempty"`
	TransNo        string  `json:"transNo,omitempty"`
	OutTransNo     string  `json:"outTransNo,omitempty"`
	StoreVidList   []int64 `json:"storeVidList,omitempty"`
	ChangeTypeList []int64 `json:"changeTypeList,omitempty"`
}

type BalanceTransferGetListResponseDataList struct {
	TransNo         string `json:"transNo,omitempty"`
	TransDate       string `json:"transDate,omitempty"`
	Wid             string `json:"wid,omitempty"`
	WidName         string `json:"widName,omitempty"`
	Customer        string `json:"customer,omitempty"`
	ChangeType      string `json:"changeType,omitempty"`
	ChangeTypeDesc  string `json:"changeTypeDesc,omitempty"`
	RuleName        string `json:"ruleName,omitempty"`
	BalanceChange   string `json:"balanceChange,omitempty"`
	Balance         string `json:"balance,omitempty"`
	PrincipalChange string `json:"principalChange,omitempty"`
	Principal       string `json:"principal,omitempty"`
	BonusChange     string `json:"bonusChange,omitempty"`
	Bonus           string `json:"bonus,omitempty"`
	OutTransNo      string `json:"outTransNo,omitempty"`
	ChannelType     int64  `json:"channelType,omitempty"`
	ChannelTypeDesc string `json:"channelTypeDesc,omitempty"`
	Remark          string `json:"remark,omitempty"`
	OccurVName      string `json:"occurVName,omitempty"`
	OperateWName    string `json:"operateWName,omitempty"`
	PhoneNo         string `json:"phoneNo,omitempty"`
	IsPositive      int64  `json:"isPositive,omitempty"`
	OccurVid        int64  `json:"occurVid,omitempty"`
}
