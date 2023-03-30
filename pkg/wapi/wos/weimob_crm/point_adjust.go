package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// PointAdjust
// @id 1664
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1664?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=调整积分)
func (client *Client) PointAdjust(request *PointAdjustRequest) (response *PointAdjustResponse, err error) {
	rpcResponse := CreatePointAdjustResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type PointAdjustRequest struct {
	*api.BaseRequest

	AdjustPointInfoList []PointAdjustRequestAdjustPointInfoList `json:"adjustPointInfoList,omitempty"`
	RequestId           string                                  `json:"requestId,omitempty"`
	RequestType         string                                  `json:"requestType,omitempty"`
	ChangePoint         int64                                   `json:"changePoint,omitempty"`
	AdjustType          string                                  `json:"adjustType,omitempty"`
	ChangeReason        string                                  `json:"changeReason,omitempty"`
	OperatorWid         string                                  `json:"operatorWid,omitempty"`
	OperateStoreVid     int64                                   `json:"operateStoreVid,omitempty"`
	SourceVid           int64                                   `json:"sourceVid,omitempty"`
	OutTransNo          string                                  `json:"outTransNo,omitempty"`
	OutTransType        string                                  `json:"outTransType,omitempty"`
	ChangeType          string                                  `json:"changeType,omitempty"`
}

type PointAdjustResponse struct {
	List   []PointAdjustResponselist `json:"list,omitempty"`
	Result bool                      `json:"result,omitempty"`
}

func CreatePointAdjustRequest() (request *PointAdjustRequest) {
	request = &PointAdjustRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "point/adjust", "apigw")
	request.Method = api.POST
	return
}

func CreatePointAdjustResponse() (response *api.BaseResponse[PointAdjustResponse]) {
	response = api.CreateResponse[PointAdjustResponse](&PointAdjustResponse{})
	return
}

type PointAdjustRequestAdjustPointInfoList struct {
	Vid         int64 `json:"vid,omitempty"`
	PointPlanId int64 `json:"pointPlanId,omitempty"`
	Wid         int64 `json:"wid,omitempty"`
}

type PointAdjustResponselist struct {
	Status  int64  `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Wid     int64  `json:"wid,omitempty"`
	WidName string `json:"widName,omitempty"`
}
