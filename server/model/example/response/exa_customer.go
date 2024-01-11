package response

import "github.com/Allen9012/ServerManager/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
