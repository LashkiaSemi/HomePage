package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_IsAdmin(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name string
		in   *User
		out  bool
	}{
		{
			name: "role is admin",
			in: &User{
				Role: "admin",
			},
			out: true,
		},
		{
			name: "role is owner",
			in: &User{
				Role: "owner",
			},
			out: true,
		},
		{
			name: "role is member",
			in: &User{
				Role: "member",
			},
			out: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			out := tt.in.IsAdmin()
			assert.Equal(t, tt.out, out)
		})
	}
}
