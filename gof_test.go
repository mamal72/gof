package gof

import (
	"os"
	"reflect"
	"testing"
)

func TestErrFileNotExists_Error(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrFileNotExists{
				path: tt.fields.path,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrFileNotExists.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrDirectoryNotExists_Error(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrDirectoryNotExists{
				path: tt.fields.path,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrDirectoryNotExists.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrFileAlreadyExists_Error(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrFileAlreadyExists{
				path: tt.fields.path,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrFileAlreadyExists.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrDirectoryAlreadyExists_Error(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrDirectoryAlreadyExists{
				path: tt.fields.path,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("ErrDirectoryAlreadyExists.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_Read(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_Write(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Write(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("gof.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Append(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Append(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("gof.Append() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Prepend(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Prepend(tt.args.path, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("gof.Prepend() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Exists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.Exists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_FileExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.FileExists(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.FileExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_DirectoryExists(t *testing.T) {
	type args struct {
		path    string
		content string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.DirectoryExists(tt.args.path, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.DirectoryExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.DirectoryExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_Mkdir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Mkdir(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("gof.Mkdir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Create(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Create(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("gof.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Touch(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Touch(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("gof.Touch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Rm(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Rm(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("gof.Rm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Chmod(t *testing.T) {
	type args struct {
		path string
		mode os.FileMode
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Chmod(tt.args.path, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("gof.Chmod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_IsDirectory(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.IsDirectory(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.IsDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.IsDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_IsFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.IsFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.IsFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("gof.IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gof_Rename(t *testing.T) {
	type args struct {
		oldPath string
		newPath string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Rename(tt.args.oldPath, tt.args.newPath); (err != nil) != tt.wantErr {
				t.Errorf("gof.Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Copy(t *testing.T) {
	type args struct {
		sourcePath      string
		destinationPath string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Copy(tt.args.sourcePath, tt.args.destinationPath); (err != nil) != tt.wantErr {
				t.Errorf("gof.Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Move(t *testing.T) {
	type args struct {
		sourcePath      string
		destinationPath string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			if err := g.Move(tt.args.sourcePath, tt.args.destinationPath); (err != nil) != tt.wantErr {
				t.Errorf("gof.Move() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_gof_Stat(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		g       *gof
		args    args
		want    os.FileInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &gof{}
			got, err := g.Stat(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("gof.Stat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gof.Stat() = %v, want %v", got, tt.want)
			}
		})
	}
}
