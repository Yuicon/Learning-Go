package main

func main() {
	foo := func() (ret string) {
		defer func() {
			if p := recover(); p != nil {
				ret = p.(string)
			}
		}()
		panic("asdasd")
	}
	println(foo())
}
