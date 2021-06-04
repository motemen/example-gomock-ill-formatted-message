package eg

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func Test(t *testing.T) {
	ctrl := gomock.NewController(t)
	foo := NewMockFoo(ctrl)
	foo.EXPECT().Foo("%s")

	foo.Foo("x")
	//     eg_test.go:14: Unexpected call to *eg.MockFoo.Foo([x]) at /src/eg_test.go:14 because:
	//         expected call at /src/eg_test.go:12 doesn't match the argument at index 0.
	//         Got: x
	//         Want: is equal to %!s(MISSING)
	//     controller.go:137: missing call(s) to *eg.MockFoo.Foo(is equal to %s) /src/eg_test.go:12
	//     controller.go:137: aborting test due to missing call(s)
}
