package challenge6

import "testing"

func TestHammingDistance(t *testing.T) {
	type args struct {
		from []byte
		to   []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"proceeding test",
			args{[]byte("this is a test"), []byte("wokka wokka!!!")},
			37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HammingDistance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
