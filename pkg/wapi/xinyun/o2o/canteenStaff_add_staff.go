package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenStaffAddStaff
// @id 1316
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加餐厅员工)
func (client *Client) CanteenStaffAddStaff(request *CanteenStaffAddStaffRequest) (response *CanteenStaffAddStaffResponse, err error) {
	rpcResponse := CreateCanteenStaffAddStaffResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenStaffAddStaffRequest struct {
	*api.BaseRequest

	HeadImageUrl       string `json:"headImageUrl,omitempty"`
	Name               string `json:"name,omitempty"`
	Phone              string `json:"phone,omitempty"`
	ProfessionalTitles int    `json:"professionalTitles,omitempty"`
	Sex                int    `json:"sex,omitempty"`
	StageName          string `json:"stageName,omitempty"`
	StoreId            int64  `json:"storeId,omitempty"`
}

type CanteenStaffAddStaffResponse struct {
}

func CreateCanteenStaffAddStaffRequest() (request *CanteenStaffAddStaffRequest) {
	request = &CanteenStaffAddStaffRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "canteenStaff/addStaff", "api")
	request.Method = api.POST
	return
}

func CreateCanteenStaffAddStaffResponse() (response *api.BaseResponse[CanteenStaffAddStaffResponse]) {
	response = api.CreateResponse[CanteenStaffAddStaffResponse](&CanteenStaffAddStaffResponse{})
	return
}
