package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantprivilegeMerchantprivilege
// @id 1868
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增账号)
func (client *Client) MerchantprivilegeMerchantprivilege(request *MerchantprivilegeMerchantprivilegeRequest) (response *MerchantprivilegeMerchantprivilegeResponse, err error) {
	rpcResponse := CreateMerchantprivilegeMerchantprivilegeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantprivilegeMerchantprivilegeRequest struct {
	*api.BaseRequest

	DepartmentList  MerchantprivilegeMerchantprivilegeRequestDepartmentList `json:"departmentList,omitempty"`
	AccountName     string                                                  `json:"accountName,omitempty"`
	AccountNum      string                                                  `json:"accountNum,omitempty"`
	LoginName       string                                                  `json:"loginName,omitempty"`
	Zone            string                                                  `json:"zone,omitempty"`
	IsCustomService bool                                                    `json:"isCustomService,omitempty"`
}

type MerchantprivilegeMerchantprivilegeResponse struct {
	Wid int64 `json:"wid,omitempty"`
	Id  int64 `json:"id,omitempty"`
}

func CreateMerchantprivilegeMerchantprivilegeRequest() (request *MerchantprivilegeMerchantprivilegeRequest) {
	request = &MerchantprivilegeMerchantprivilegeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantprivilege/merchantprivilege", "api")
	request.Method = api.POST
	return
}

func CreateMerchantprivilegeMerchantprivilegeResponse() (response *api.BaseResponse[MerchantprivilegeMerchantprivilegeResponse]) {
	response = api.CreateResponse[MerchantprivilegeMerchantprivilegeResponse](&MerchantprivilegeMerchantprivilegeResponse{})
	return
}

type MerchantprivilegeMerchantprivilegeRequestDepartmentList struct {
	DepartmentId      int64  `json:"departmentId,omitempty"`
	DepartmentCode    string `json:"departmentCode,omitempty"`
	DepartmentType    int64  `json:"departmentType,omitempty"`
	DepartmentStoreId int64  `json:"departmentStoreId,omitempty"`
}
