package domain

import (
	"bangking/dto"

	"github.com/nothingprogram/banking-lib/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}

}

type CustomerRepository interface {
	// FindAll status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	// ById id 에 해당하는 고객이 없는 경우 nil 을 보내기 위해 포인터를 리턴합니다.
	ById(string) (*Customer, *errs.AppError)
}
