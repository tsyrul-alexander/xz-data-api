package identity

import (
	"encoding/json"
	"github.com/tsyrul-alexander/xz-identity-api/model"
	"github.com/tsyrul-alexander/xz-identity-api/model/response"
	"net/http"
	"strconv"
	"time"
)

const (
	userInRoleServicePath string = "/authorization/get-user-in-role"
)

type Service struct {
	Address string
	client http.Client
}

func CreateService(address string, timeout int) *Service {
	var duration = timeout * int(time.Millisecond)
	return &Service{
		Address: address,
		client:  http.Client{Timeout:time.Duration(duration)},
	}
}

func (service *Service) GetUserInRoles(roles []model.UserRole, token string) (bool, error) {
	var req, errRequest = http.NewRequest("GET", service.Address + userInRoleServicePath, nil)
	if errRequest != nil {
		return false, errRequest
	}
	q := req.URL.Query()
	q.Add("token", token)
	for _, role := range roles {
		q.Add("role", strconv.Itoa(int(role)))
	}
	req.URL.RawQuery = q.Encode()
	var resp, errResponse = service.client.Do(req)
	if errResponse != nil {
		return false, errResponse
	}
	var identityResponse = &response.IdentityResponse{}
	var errParse = json.NewDecoder(resp.Body).Decode(identityResponse)
	if errParse != nil {
		return false, errParse
	}
	return identityResponse.Success, nil
}
