package foo

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"Test Case #1",
			args:args{
				number:2,
			},
			want:true,
		},

		{
			name:"Test Case #2",
			args:args{
				number:3,
			},
			want:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.number); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
