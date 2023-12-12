package aggregate

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ntnghiatn/ddd-go/entity"
	"github.com/ntnghiatn/ddd-go/valueobject"
)

var (
	ErrInvalidPerson = errors.New("err in concrete Person - Invalid Person")
)

// Customer - các aggregates (tổng hợp) là không thể lấy dữ liệu trực tiếp và data không thể truy cập được từ bên ngoài.
// không sử dụng bất cứ thẻ `json` nào
type Customer struct {
	// dùng chữ thường để không thể truy cập được từ các miền khác.
	// dùng con trỏ vì entity có thể thay đổi Trạng thái được.
	person   *entity.Person
	products []*entity.Item

	// không dùng con trỏ vì không cho phép thay đổi
	transactions []valueobject.Transaction
}

// NewCustomer - áp dụng Factory pattern để new Customer (concrete product)
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
