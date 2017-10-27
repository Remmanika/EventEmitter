package emitter

import "fmt"

func ExampleSimple() {
	application := make(EventEmitter)

	application.On("smth", func(data ...interface{}) {
		fmt.Println(data...)
	})

	application.Emit("smth", struct{ A int }{A: 5})

	// Output: {5}
}
