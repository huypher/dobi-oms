package auth

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pghuy/dobi-oms/domain"
)

func Test_genJWT(t *testing.T) {
	t.Parallel()

	acc := &domain.Account{
		ID:       1,
		UserName: "username",
		Password: "password",
		Name:     "name",
	}

	secretJWT := "jwtToken"

	testCases := []struct {
		name string
		acc  *domain.Account
		err  error
	}{
		{"test", acc, nil},
		{"test", nil, errors.New("user empty")},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := genJWT(tc.acc, secretJWT)
			require.Equal(t, tc.err, err)
			if tc.err == nil {
				require.NotEmpty(t, got)
			}
		})
	}
}
