package memory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/ntnghiatn/ddd-go/aggregate"
	"github.com/ntnghiatn/ddd-go/domain/customer"
	"testing"
)

func TestMemoryRepo_Get(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()
	repo := MemoryRepo{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("deb6840a-98bc-11ee-b9d1-0242ac120002"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v get %v", tc.expectedErr, err)
			}
		})
	}
}
