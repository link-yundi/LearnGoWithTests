package HelloWorld

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == ""{
		name = "World"
	}
	return englishHelloPrefix + name
}

//func TestHello(t *testing.T) {
	//got := Hello("ZhangYundi")
	//want := "Hello ZhangYundi"
	//t.Logf("%[1]q, %[1]s", got)
	//assert.Equal(t, want, got)
//	t.Run()
//}


func TestHello(t *testing.T) {
	// define assert func
	//assertCorrectMessage := func(t *testing.T, got, want string) {
	//	t.Helper()
	//	if got != want {
	//		t.Errorf("got '%q', want '%q'", got, want)
	//	}
	//}
	assertT := assert.New(t)
	tests := []struct {
		name string
		input string
		want string
	}{
		// TODO: test cases
		{
			name: "saying hello to people",
			input: "ZhangYundi",
			want: "Hello, ZhangYundi",
		},
		{
			name: "say hello world when an empty string is supplied",
			input: "",
			want: "Hello, World",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Hello(test.input)
			//assertCorrectMessage(t, got, test.want)
			assertT.Equal(test.want, got)
		})
	}
}


