package auth

import (
	"testing"

	"github.com/pghuy/dobi-oms/domain"
)

func Test_genJWT(t *testing.T) {
	type args struct {
		acc       *domain.Account
		secretJWT string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				acc: &domain.Account{
					ID:       1,
					UserName: "huy",
					Password: "huy",
					Name:     "huy",
				},
				secretJWT: "jwtToken",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := genJWT(tt.args.acc, tt.args.secretJWT)
			if (err != nil) != tt.wantErr {
				t.Errorf("genJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("genJWT() got = %v, want %v", got, tt.want)
			}
		})
	}
}
