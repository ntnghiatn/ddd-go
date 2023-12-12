package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ntnghiatn/ddd-go/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("not found")
	ErrFailedToAddCustomer = errors.New("failed to add customer")
	ErrUpdateCustomer      = errors.New("error update customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
