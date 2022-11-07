# Generator

Task: Create a function that could be used to fetch an infinite amount of data
with limited memory.

Specifically, write a function that would a new random id on every call.

Usage could be like this:

```
func main() {
	for v := range generateId() {
		fmt.Println(v)
	}
}
```