package Demo19

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2, 3) = %d; want 5", result)
	}

	tests := []struct {
		Name           string
		A, B, Excepted int
	}{
		{"add1", 1, 2, 3},
		{"add2", 3, 5, 8},
		{"add3", 4, 7, 9},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := Add(tt.A, tt.B); got != tt.Excepted {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.A, tt.B, got, tt.Excepted)
			}
		})
	}
	fmt.Print("测试成功")
}
