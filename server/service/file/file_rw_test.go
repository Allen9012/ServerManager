package file

import (
	"github.com/Allen9012/ServerManager/server/model/file"
	"testing"
)

func TestFileRWService_CheckPermission(t *testing.T) {
	type args struct {
		authID   uint
		filePath string
		flag     file.RWAction
	}
	//Do,     ^/?([^/]+/)*?allen(/[^/]+)*?/[^/]+\.sh$
	//Up,     ^/allen/123(/[^/]+)*$
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "上传文件", args: args{
			authID:   9528,
			filePath: "/home/allen/proxy.sh",
			flag:     file.Upload,
		}, want: true},
		{name: "下载文件", args: args{
			authID:   9528,
			filePath: "/home/allen/proxy.sh",
			flag:     file.Download,
		}, want: true},
		{name: "上传文件false", args: args{
			authID:   9528,
			filePath: "/home/allen/file1/123",
			flag:     file.Upload}, want: false},
		{name: "下载文件false", args: args{
			authID:   9528,
			filePath: "/home/allen/123",
			flag:     file.Download,
		}, want: false},
	}
	var permissions []file.FilePermission = []file.FilePermission{
		{PermissionState: 3, Regexp: "^/?([^/]+/)*?allen(/[^/]+)*?/[^/]+\\.sh$"},
		{PermissionState: 2, Regexp: "^/?([^/]+/)*?allen/123(/[^/]+)*$"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _checkPermission(permissions, tt.args.filePath, tt.args.flag); got != tt.want {
				//if got := FileRWService_CheckPermission2(permissions, tt.args.filePath, tt.args.flag); got != tt.want {
				t.Errorf("CheckPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func FileRWService_CheckPermission2(permissions []file.FilePermission, filePath string, flag file.RWAction) bool {
//	for _, permission := range permissions {
//		re := regexp.MustCompile(permission.Regexp)
//		if re.MatchString(filePath) {
//			return true
//		} else {
//			continue
//		}
//	}
//	return false
//}
