package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FeicanStaffUpdateStaffInfoByStaffId
// @id 1321
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新员工信息)
func (client *Client) FeicanStaffUpdateStaffInfoByStaffId(request *FeicanStaffUpdateStaffInfoByStaffIdRequest) (response *FeicanStaffUpdateStaffInfoByStaffIdResponse, err error) {
	rpcResponse := CreateFeicanStaffUpdateStaffInfoByStaffIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FeicanStaffUpdateStaffInfoByStaffIdRequest struct {
	*api.BaseRequest

	TechnicianInfo     FeicanStaffUpdateStaffInfoByStaffIdRequestTechnicianInfo `json:"technicianInfo,omitempty"`
	HeadImageUrl       string                                                   `json:"headImageUrl,omitempty"`
	Identity           int64                                                    `json:"identity,omitempty"`
	Name               string                                                   `json:"name,omitempty"`
	Phone              string                                                   `json:"phone,omitempty"`
	ProfessionalTitles string                                                   `json:"professionalTitles,omitempty"`
	Sex                int64                                                    `json:"sex,omitempty"`
	StaffId            int64                                                    `json:"staffId,omitempty"`
	StageName          string                                                   `json:"stageName,omitempty"`
	Status             int64                                                    `json:"status,omitempty"`
	StoreId            int64                                                    `json:"storeId,omitempty"`
}

type FeicanStaffUpdateStaffInfoByStaffIdResponse struct {
}

func CreateFeicanStaffUpdateStaffInfoByStaffIdRequest() (request *FeicanStaffUpdateStaffInfoByStaffIdRequest) {
	request = &FeicanStaffUpdateStaffInfoByStaffIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "feicanStaff/updateStaffInfoByStaffId", "api")
	request.Method = api.POST
	return
}

func CreateFeicanStaffUpdateStaffInfoByStaffIdResponse() (response *api.BaseResponse[FeicanStaffUpdateStaffInfoByStaffIdResponse]) {
	response = api.CreateResponse[FeicanStaffUpdateStaffInfoByStaffIdResponse](&FeicanStaffUpdateStaffInfoByStaffIdResponse{})
	return
}

type FeicanStaffUpdateStaffInfoByStaffIdRequestTechnicianInfo struct {
	AbilityServices    []FeicanStaffUpdateStaffInfoByStaffIdRequestAbilityServices `json:"abilityServices,omitempty"`
	Productions        []FeicanStaffUpdateStaffInfoByStaffIdRequestProductions     `json:"productions,omitempty"`
	GoodAtList         []string                                                    `json:"goodAtList,omitempty"`
	Introduction       string                                                      `json:"introduction,omitempty"`
	Permission         int64                                                       `json:"permission,omitempty"`
	PhotoAlbumName     string                                                      `json:"photoAlbumName,omitempty"`
	VirtualReservation int64                                                       `json:"virtualReservation,omitempty"`
	WorkingSeniority   int64                                                       `json:"workingSeniority,omitempty"`
}

type FeicanStaffUpdateStaffInfoByStaffIdRequestAbilityServices struct {
	GoodsId int64 `json:"goodsId,omitempty"`
}

type FeicanStaffUpdateStaffInfoByStaffIdRequestProductions struct {
	ProductionIntroduce string `json:"productionIntroduce,omitempty"`
	ProductionUrl       string `json:"productionUrl,omitempty"`
}
