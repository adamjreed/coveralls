package coveralls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEven(t *testing.T) {
	cases := []struct {
		name  string
		input int64
		pass  bool
	}{
		{
			"even number",
			2,
			true,
		},
		{
			"odd number",
			3,
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.pass, Even(c.input))
		})
	}
}

func TestOdd(t *testing.T) {
	cases := []struct {
		name  string
		input int64
		pass  bool
	}{
		{
			"odd number",
			3,
			true,
		},
		{
			"even number",
			2,
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.pass, Odd(c.input))
		})
	}
}

func TestIsTwo(t *testing.T) {
	cases := []struct {
		name  string
		input int64
		pass  bool
	}{
		{
			"two",
			2,
			true,
		},
		{
			"not two",
			3,
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.pass, IsTwo(c.input))
		})
	}
}
