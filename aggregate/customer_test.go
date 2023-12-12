package aggregate_test

import (
	"errors"
	"github.com/ntnghiatn/ddd-go/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testcases := []testCase{
		{
			test:        "Empty",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		},
		{
			test:        "Valid Name",
			name:        "Nghia",
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected Error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
