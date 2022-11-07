package main

func main() {}

type Foo struct{}

func (f *Foo) Read(p []byte) (int, error) {
	return len(p), nil
}
