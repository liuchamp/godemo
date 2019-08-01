package post

import "reflect"

type Post struct {
}

func (p *Post) GetPackage() string {
	return reflect.TypeOf(p).PkgPath()
}
