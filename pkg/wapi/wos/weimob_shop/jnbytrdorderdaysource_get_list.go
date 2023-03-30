package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// JnbytrdorderdaysourceGetList
// @id 2270
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2270?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=江南布衣访问来源数据接口)
func (client *Client) JnbytrdorderdaysourceGetList(request *JnbytrdorderdaysourceGetListRequest) (response *JnbytrdorderdaysourceGetListResponse, err error) {
	rpcResponse := CreateJnbytrdorderdaysourceGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type JnbytrdorderdaysourceGetListRequest struct {
	*api.BaseRequest

	StatType []int64 `json:"statType,omitempty"`
	BosId    []int64 `json:"bosId,omitempty"`
	StartDd  string  `json:"startDd,omitempty"`
	EndDd    string  `json:"endDd,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
	PageNum  int64   `json:"pageNum,omitempty"`
}

type JnbytrdorderdaysourceGetListResponse struct {
	List      []JnbytrdorderdaysourceGetListResponselist `json:"list,omitempty"`
	PageNum   int64                                      `json:"pageNum,omitempty"`
	PageSize  int64                                      `json:"pageSize,omitempty"`
	TotalSize int64                                      `json:"totalSize,omitempty"`
}

func CreateJnbytrdorderdaysourceGetListRequest() (request *JnbytrdorderdaysourceGetListRequest) {
	request = &JnbytrdorderdaysourceGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "jnbytrdorderdaysource/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateJnbytrdorderdaysourceGetListResponse() (response *api.BaseResponse[JnbytrdorderdaysourceGetListResponse]) {
	response = api.CreateResponse[JnbytrdorderdaysourceGetListResponse](&JnbytrdorderdaysourceGetListResponse{})
	return
}

type JnbytrdorderdaysourceGetListResponselist struct {
	Uv           int64  `json:"uv,omitempty"`
	StartDd      string `json:"startDd,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	PayCnt       int64  `json:"payCnt,omitempty"`
	PayAmt       int64  `json:"payAmt,omitempty"`
	PayVisitRate int64  `json:"payVisitRate,omitempty"`
	DateType     int64  `json:"dateType,omitempty"`
	AreaId       string `json:"areaId,omitempty"`
	PayUsrCnt    int64  `json:"payUsrCnt,omitempty"`
	AreaName     string `json:"areaName,omitempty"`
	OrdCnt       int64  `json:"ordCnt,omitempty"`
	BosId        int64  `json:"bosId,omitempty"`
	StatName     string `json:"statName,omitempty"`
	EndDd        string `json:"endDd,omitempty"`
	StoreName    string `json:"storeName,omitempty"`
	Brand        string `json:"brand,omitempty"`
	StoreCode    string `json:"storeCode,omitempty"`
}
