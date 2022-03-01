package command

import "testing"

func Test_newWrongNumOfArgsError(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "set",
			args:    args{cmd: "set"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := newWrongNumOfArgsError(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("newWrongNumOfArgsError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
