package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Add(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", answer, err)
	}

	require.Equal(t, int32(2), answer)
}

func TestSub(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Sub(2, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", answer, err)
	}

	require.Equal(t, int32(1), answer)
}

func TestMulti(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multi(2, 3)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", answer, err)
	}

	require.Equal(t, int32(6), answer)
}

func TestDiv(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Div(4, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", answer, err)
	}

	require.Equal(t, int32(2), answer)
}

func TestDivZero(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Div(4, 0)
	require.True(t, err != nil)
	require.Errorf(t, err, "divide by zero")
	require.Equal(t, int32(0), answer)
}
