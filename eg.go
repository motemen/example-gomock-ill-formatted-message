//go:generate mockgen -package eg -destination foo_mock.go . Foo
package eg

type Foo interface {
	Foo(x interface{}) string
}
