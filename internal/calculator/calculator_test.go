package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive numbers",
			a:    5,
			b:    3,
			want: 8,
		},
		{
			name: "negative numbers",
			a:    -5,
			b:    -3,
			want: -8,
		},
		{
			name: "mixed signs positive result",
			a:    5,
			b:    -3,
			want: 2,
		},
		{
			name: "mixed signs negative result",
			a:    -5,
			b:    3,
			want: -2,
		},
		{
			name: "zero and positive",
			a:    0,
			b:    5,
			want: 5,
		},
		{
			name: "zero and negative",
			a:    0,
			b:    -5,
			want: -5,
		},
		{
			name: "both zeros",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive numbers",
			a:    5,
			b:    3,
			want: 2,
		},
		{
			name: "negative numbers",
			a:    -5,
			b:    -3,
			want: -2,
		},
		{
			name: "mixed signs positive result",
			a:    5,
			b:    -3,
			want: 8,
		},
		{
			name: "mixed signs negative result",
			a:    -5,
			b:    3,
			want: -8,
		},
		{
			name: "zero and positive",
			a:    0,
			b:    5,
			want: -5,
		},
		{
			name: "zero and negative",
			a:    0,
			b:    -5,
			want: 5,
		},
		{
			name: "both zeros",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive numbers",
			a:    5,
			b:    3,
			want: 15,
		},
		{
			name: "negative numbers",
			a:    -5,
			b:    -3,
			want: 15,
		},
		{
			name: "mixed signs",
			a:    5,
			b:    -3,
			want: -15,
		},
		{
			name: "mixed signs reversed",
			a:    -5,
			b:    3,
			want: -15,
		},
		{
			name: "zero and positive",
			a:    0,
			b:    5,
			want: 0,
		},
		{
			name: "zero and negative",
			a:    0,
			b:    -5,
			want: 0,
		},
		{
			name: "both zeros",
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive numbers",
			a:    6,
			b:    3,
			want: 2,
		},
		{
			name: "negative numbers",
			a:    -6,
			b:    -3,
			want: 2,
		},
		{
			name: "mixed signs",
			a:    6,
			b:    -3,
			want: -2,
		},
		{
			name: "mixed signs reversed",
			a:    -6,
			b:    3,
			want: -2,
		},
		{
			name: "zero dividend",
			a:    0,
			b:    5,
			want: 0,
		},
		{
			name: "integer division truncates",
			a:    7,
			b:    3,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDivide_DivisionByZero(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "positive dividend",
			a:    5,
			b:    0,
		},
		{
			name: "negative dividend",
			a:    -5,
			b:    0,
		},
		{
			name: "zero dividend",
			a:    0,
			b:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			require.Error(t, err)
			assert.True(t, errors.Is(err, ErrDivisionByZero))
			assert.Equal(t, 0, got)
		})
	}
}
