package main

import "testing"

func TestRepeatingKeyXOR(t *testing.T) {
	type args struct {
		key   []byte
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				key: []byte(`ICE`),
				input: []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`),
			},
			want: `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RepeatingKeyXOR(tt.args.key, tt.args.input); got != tt.want {
				t.Errorf("RepeatingKeyXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
