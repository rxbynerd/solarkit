package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Main(t *testing.T) {
	r := require.New(t)

	r.Equal(1, 1)
}
