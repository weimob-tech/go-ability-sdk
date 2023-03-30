package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StockInventoryGetList
// @id 2288
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2288?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询进销存单据列表)
func (client *Client) StockInventoryGetList(request *StockInventoryGetListRequest) (response *StockInventoryGetListResponse, err error) {
	rpcResponse := CreateStockInventoryGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StockInventoryGetListRequest struct {
	*api.BaseRequest

	BasicInfo      StockInventoryGetListRequestBasicInfo      `json:"basicInfo,omitempty"`
	QueryParameter StockInventoryGetListRequestQueryParameter `json:"queryParameter,omitempty"`
	PageNum        int64                                      `json:"pageNum,omitempty"`
	PageSize       int64                                      `json:"pageSize,omitempty"`
}

type StockInventoryGetListResponse struct {
	PageList   []StockInventoryGetListResponsePageList `json:"pageList,omitempty"`
	PageNum    int64                                   `json:"pageNum,omitempty"`
	PageSize   int64                                   `json:"pageSize,omitempty"`
	TotalCount int64                                   `json:"totalCount,omitempty"`
}

func CreateStockInventoryGetListRequest() (request *StockInventoryGetListRequest) {
	request = &StockInventoryGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "stock/inventory/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateStockInventoryGetListResponse() (response *api.BaseResponse[StockInventoryGetListResponse]) {
	response = api.CreateResponse[StockInventoryGetListResponse](&StockInventoryGetListResponse{})
	return
}

type StockInventoryGetListRequestBasicInfo struct {
	Vid   int64 `json:"vid,omitempty"`
	BosId int64 `json:"bosId,omitempty"`
}

type StockInventoryGetListRequestQueryParameter struct {
	SearchType  int64  `json:"searchType,omitempty"`
	Search      string `json:"search,omitempty"`
	OperateType int64  `json:"operateType,omitempty"`
	BizType     int64  `json:"bizType,omitempty"`
	BeginTime   int64  `json:"beginTime,omitempty"`
	EndTime     int64  `json:"endTime,omitempty"`
}

type StockInventoryGetListResponsePageList struct {
	ReceiptId         int64  `json:"receiptId,omitempty"`
	BusinessReceiptId int64  `json:"businessReceiptId,omitempty"`
	OutOrderId        int64  `json:"outOrderId,omitempty"`
	TotalQuantity     string `json:"totalQuantity,omitempty"`
	OperateId         int64  `json:"operateId,omitempty"`
	OperatePhone      int64  `json:"operatePhone,omitempty"`
	OperateName       string `json:"operateName,omitempty"`
	Remark            string `json:"remark,omitempty"`
	BizType           int64  `json:"bizType,omitempty"`
	BizTypeName       string `json:"bizTypeName,omitempty"`
	CreateTime        string `json:"createTime,omitempty"`
}
