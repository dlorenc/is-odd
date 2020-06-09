package odd

import "testing"

func TestIsOdd(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "zero",
			n:    0,
			want: false,
		},
		{
			name: "one",
			n:    1,
			want: true,
		},
		{
			name: "three",
			n:    3,
			want: true,
		},
		{
			name: "ten",
			n:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.n); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
