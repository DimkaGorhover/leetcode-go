package p0071

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	t.Parallel()
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{path: "/home/"},
			want: "/home",
		},
		{
			args: args{path: "/../"},
			want: "/",
		},
		{
			args: args{path: "/home//foo/"},
			want: "/home/foo",
		},
		{
			args: args{path: "/a/./b/../../c/"},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, simplifyPath(tt.args.path))
		})
	}
}

func Test_splitter_next(t *testing.T) {
	t.Parallel()

	t.Run(``, func(t *testing.T) {
		s := newSplitter("/path/../to//dir/")

		for i := 0; i < 10; i++ {
			text, found := s.next()
			fmt.Printf("text = [%s]; found = %v\n", text, found)
		}

	})

}
