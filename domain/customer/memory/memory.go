package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ntnghiatn/ddd-go/aggregate"
	"github.com/ntnghiatn/ddd-go/domain/customer"
	"sync"
)

type MemoryRepo struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepo {

	return &MemoryRepo{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryRepo) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepo) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	if _, ok := mr.customers[c.GetID()]; ok {
		return errors.New("")
	}

	//
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()

	return nil
}

func (mr *MemoryRepo) Update(c aggregate.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return errors.New("error")
	}

	//
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
