package test

import (
	"simigo/simisimi"
	"testing"
)

func TestSimi(t *testing.T) {
	result := simisimi.GetSimi("aku ganteng ga?")

	t.Log(result)
}
