package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffAddStaff
// @id 1322
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加员工)
func (client *Client) FeicanStaffAddStaff(request *FeicanStaffAddStaffRequest) (response *FeicanStaffAddStaffResponse, err error) {
	rpcResponse := CreateFeicanStaffAddStaffResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffAddStaffRequest struct {
	*api.BaseRequest

	TechnicianInfo     FeicanStaffAddStaffRequestTechnicianInfo `json:"technicianInfo,omitempty"`
	HeadImageUrl       string                                   `json:"headImageUrl,omitempty"`
	Identity           int64                                    `json:"identity,omitempty"`
	Name               string                                   `json:"name,omitempty"`
	Phone              string                                   `json:"phone,omitempty"`
	ProfessionalTitles string                                   `json:"professionalTitles,omitempty"`
	Sex                string                                   `json:"sex,omitempty"`
	StageName          string                                   `json:"stageName,omitempty"`
	Status             int64                                    `json:"status,omitempty"`
	StoreId            int64                                    `json:"storeId,omitempty"`
}

type FeicanStaffAddStaffResponse struct {
}

func CreateFeicanStaffAddStaffRequest() (request *FeicanStaffAddStaffRequest) {
	request = &FeicanStaffAddStaffRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/addStaff", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffAddStaffResponse() (response *api.BaseResponse[FeicanStaffAddStaffResponse]) {
	response = api.CreateResponse[FeicanStaffAddStaffResponse](&FeicanStaffAddStaffResponse{})
	return
}

type FeicanStaffAddStaffRequestTechnicianInfo struct {
	AbilityServices    []map[string]any `json:"abilityServices,omitempty"`
	Productions        []map[string]any `json:"productions,omitempty"`
	GoodAtList         []string         `json:"goodAtList,omitempty"`
	Introduction       string           `json:"introduction,omitempty"`
	Permission         int64            `json:"permission,omitempty"`
	PhotoAlbumName     string           `json:"photoAlbumName,omitempty"`
	VirtualReservation int64            `json:"virtualReservation,omitempty"`
	WorkingSeniority   int64            `json:"workingSeniority,omitempty"`
}
