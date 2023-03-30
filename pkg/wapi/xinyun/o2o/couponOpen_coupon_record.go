package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CouponOpenCouponRecord
// @id 1310
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询卡券核销记录)
func (client *Client) CouponOpenCouponRecord(request *CouponOpenCouponRecordRequest) (response *CouponOpenCouponRecordResponse, err error) {
	rpcResponse := CreateCouponOpenCouponRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CouponOpenCouponRecordRequest struct {
	*api.BaseRequest

	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Index     string `json:"index,omitempty"`
}

type CouponOpenCouponRecordResponse struct {
}

func CreateCouponOpenCouponRecordRequest() (request *CouponOpenCouponRecordRequest) {
	request = &CouponOpenCouponRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "couponOpen/couponRecord", "api")
	request.Method = api.POST
	return
}

func CreateCouponOpenCouponRecordResponse() (response *api.BaseResponse[CouponOpenCouponRecordResponse]) {
	response = api.CreateResponse[CouponOpenCouponRecordResponse](&CouponOpenCouponRecordResponse{})
	return
}
