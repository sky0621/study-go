package testableexample

type Foo struct {
	Name string
}

func Hello() string {
	return "Hello, someone."
}

func (f *Foo) Hello() string {
	return "Hello, " + f.Name
}
