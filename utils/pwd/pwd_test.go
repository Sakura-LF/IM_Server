package pwd

import (
	"fmt"
	"testing"
)

func TestCreatePws(t *testing.T) {
	pwd := HashPwd("Sakura")
	fmt.Println(pwd)
}

func TestHashPwd(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Pwd", args: args{pwd: "Sakura"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwd := HashPwd(tt.args.pwd)
			if checkPwd := CheckPwd(pwd, tt.args.pwd); checkPwd != tt.want {
				t.Errorf("HashPwd() = %v, want %v", checkPwd, tt.want)
			}
		})
	}
}
