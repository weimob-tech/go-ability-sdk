package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GuideGetCommissionBillList
// @id 1135
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询提成统计)
func (client *Client) GuideGetCommissionBillList(request *GuideGetCommissionBillListRequest) (response *GuideGetCommissionBillListResponse, err error) {
	rpcResponse := CreateGuideGetCommissionBillListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GuideGetCommissionBillListRequest struct {
	*api.BaseRequest

	PageNum     int    `json:"pageNum,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	GuiderWid   int64  `json:"guiderWid,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	GuiderPhone string `json:"guiderPhone,omitempty"`
	JobNumber   string `json:"jobNumber,omitempty"`
}

type GuideGetCommissionBillListResponse struct {
}

func CreateGuideGetCommissionBillListRequest() (request *GuideGetCommissionBillListRequest) {
	request = &GuideGetCommissionBillListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "guide/getCommissionBillList", "api")
	request.Method = api.POST
	return
}

func CreateGuideGetCommissionBillListResponse() (response *api.BaseResponse[GuideGetCommissionBillListResponse]) {
	response = api.CreateResponse[GuideGetCommissionBillListResponse](&GuideGetCommissionBillListResponse{})
	return
}
