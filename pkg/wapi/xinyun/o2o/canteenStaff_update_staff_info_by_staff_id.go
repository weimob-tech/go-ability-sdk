package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffUpdateStaffInfoByStaffId
// @id 1315
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新餐厅员工信息)
func (client *Client) CanteenStaffUpdateStaffInfoByStaffId(request *CanteenStaffUpdateStaffInfoByStaffIdRequest) (response *CanteenStaffUpdateStaffInfoByStaffIdResponse, err error) {
	rpcResponse := CreateCanteenStaffUpdateStaffInfoByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffUpdateStaffInfoByStaffIdRequest struct {
	*api.BaseRequest

	HeadImageUrl       string `json:"headImageUrl,omitempty"`
	Name               string `json:"name,omitempty"`
	Phone              string `json:"phone,omitempty"`
	ProfessionalTitles string `json:"professionalTitles,omitempty"`
	Sex                int64  `json:"sex,omitempty"`
	StaffId            int64  `json:"staffId,omitempty"`
	StageName          string `json:"stageName,omitempty"`
	StoreId            int64  `json:"storeId,omitempty"`
}

type CanteenStaffUpdateStaffInfoByStaffIdResponse struct {
}

func CreateCanteenStaffUpdateStaffInfoByStaffIdRequest() (request *CanteenStaffUpdateStaffInfoByStaffIdRequest) {
	request = &CanteenStaffUpdateStaffInfoByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/updateStaffInfoByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffUpdateStaffInfoByStaffIdResponse() (response *api.BaseResponse[CanteenStaffUpdateStaffInfoByStaffIdResponse]) {
	response = api.CreateResponse[CanteenStaffUpdateStaffInfoByStaffIdResponse](&CanteenStaffUpdateStaffInfoByStaffIdResponse{})
	return
}
