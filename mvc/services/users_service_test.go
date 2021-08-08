package services

import (
	"github.com/Chestermozhao/GolangFromZero/mvc/domain"
	"github.com/Chestermozhao/GolangFromZero/mvc/utils"
	"net/http"
	"reflect"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userId)
}

func Test_usersService_GetUser(t *testing.T) {
	type args struct {
		userId int64
	}
	tests := []struct {
		name  string
		args  args
		want  *domain.User
		want1 *utils.ApplicationError
	}{
		{
			name: "Test mock get not found user",
			args: args{
				userId: 0,
			},
			want1: &utils.ApplicationError{
				Message: "user 123 does not exists",
				StatusCode: http.StatusNotFound,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usersService{}
			getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
				return nil, &utils.ApplicationError{
					StatusCode: http.StatusNotFound,
					Message: "user 123 does not exists",
				}
			}
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
