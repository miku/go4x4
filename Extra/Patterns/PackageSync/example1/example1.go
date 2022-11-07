package main

var sync.Pool

func main() {
	buf := bufPool.Get().(bytes.Buffer)
	defer bufPool.Put(buf)

	v:= buf.Get(bytes.Buffer)

	buf.Put(v)
}
