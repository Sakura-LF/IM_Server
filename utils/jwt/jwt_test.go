package jwt

import (
	"log"
	"testing"
)

//func TestParseToken(t *testing.T) {
//	GenToken(JwtPayLoad{
//		UserID:   1,
//		Nickname: "",
//		Role:     0,
//	})
//}

func TestGenToken(t *testing.T) {
	type args struct {
		payload      JwtPayLoad
		accessSecret string
		expires      int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "JwtTest",
			args: args{
				payload: JwtPayLoad{
					UserID:   1,
					Role:     1,
					Nickname: "Sakura",
				},
				accessSecret: "123456",
				expires:      8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenToken(tt.args.payload, tt.args.accessSecret, tt.args.expires)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			token, err := ParseToken(got, tt.args.accessSecret)
			//if token != tt.want {
			//	t.Errorf("GenToken() got = %v, want %v", got, tt.want)
			//}
			log.Println(token)
		})
	}
}
