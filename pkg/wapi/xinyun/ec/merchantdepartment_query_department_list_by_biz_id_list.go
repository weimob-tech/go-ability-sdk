package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdepartmentQueryDepartmentListByBizIdList
// @id 1493
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据门店storeId回溯上级区域列表)
func (client *Client) MerchantdepartmentQueryDepartmentListByBizIdList(request *MerchantdepartmentQueryDepartmentListByBizIdListRequest) (response *MerchantdepartmentQueryDepartmentListByBizIdListResponse, err error) {
	rpcResponse := CreateMerchantdepartmentQueryDepartmentListByBizIdListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdepartmentQueryDepartmentListByBizIdListRequest struct {
	*api.BaseRequest

	StoreId int64 `json:"storeId,omitempty"`
}

type MerchantdepartmentQueryDepartmentListByBizIdListResponse struct {
}

func CreateMerchantdepartmentQueryDepartmentListByBizIdListRequest() (request *MerchantdepartmentQueryDepartmentListByBizIdListRequest) {
	request = &MerchantdepartmentQueryDepartmentListByBizIdListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdepartment/queryDepartmentListByBizIdList", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdepartmentQueryDepartmentListByBizIdListResponse() (response *api.BaseResponse[MerchantdepartmentQueryDepartmentListByBizIdListResponse]) {
	response = api.CreateResponse[MerchantdepartmentQueryDepartmentListByBizIdListResponse](&MerchantdepartmentQueryDepartmentListByBizIdListResponse{})
	return
}
