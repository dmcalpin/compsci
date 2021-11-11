package creational

// With a singleton, only one instance
// of the object ever exists.
//
// In Go there are no constructors,
// so the closest I could come is making
// the struct private, and only exposing
// it via GetSingletonInstance. This
// method then ensures that only one instance
// ever exists

// Do not instantiate this object directly
// use GetSingletonInstance. This is only
// a concern from within the package.
type singleton struct {
	Foo interface{}
	Bar interface{}
}

// the single instance
var s = &singleton{}

// GetSingletonInstance is the gatekeeper
// of the singleton. Those that import
// this package will only be able to
// get the reference by calling this
// function.
func GetSingletonInstance() *singleton {
	return s
}
