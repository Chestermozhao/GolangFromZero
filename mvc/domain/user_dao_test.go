package domain

import (
	"github.com/Chestermozhao/GolangFromZero/mvc/utils"
	"reflect"
	"testing"
)

func Test_userDao_GetUser(t *testing.T) {
	type args struct {
		userId int64
	}
	type want1 struct {
		Message string
		StatusCode int64
		Code string
	}
	var tests = []struct {
		name  string
		args  args
		want  *User
		want1 *utils.ApplicationError
	}{
		// TODO: Add test cases.
		{
			name: "Get User Not Found",
			args: args{
				userId: 1234,
			},
			want1: &utils.ApplicationError{
				"user 1234 does not exists",
				404,
				"not_found",
			},
		},
		{
			name: "Get Exist User",
			args: args{
				userId: 123,
			},
			want: &User{
				123,
				"Chester",
				"Mo",
				"chestermo_test@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userDao{}
			got, got1 := u.GetUser(tt.args.userId)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
