package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"golang.org/x/tour/tree"
)

func TestSameWithSameNumberOfNodes(t *testing.T) {
	tests := []struct {
		name string
		k1   int
		k2   int
		want bool
	}{
		{"Equivalent trees", 1, 1, true},
		{"Tree with different values", 1, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1 := tree.New(tt.k1)
			t2 := tree.New(tt.k2)
			got := Same(t1, t2)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSameWithDifferentNumberOfNodes(t *testing.T) {
	type args struct {
		k1, n1 int
		k2, n2 int
	}
	tests := []struct {
		name string
		arg  args
		want bool
	}{
		{"Equivalent trees", args{k1: 1, n1: 100, k2: 1, n2: 100}, true},
		{"Tree with different number of nodes", args{k1: 1, n1: 100, k2: 1, n2: 101}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1 := NewTreeVarNodes(tt.arg.k1, tt.arg.n1)
			t2 := NewTreeVarNodes(tt.arg.k2, tt.arg.n2)
			got := Same(t1, t2)
			assert.Equal(t, tt.want, got)
		})
	}
}
