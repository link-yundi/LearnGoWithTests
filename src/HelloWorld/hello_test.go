package HelloWorld

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Hello(name string) string {
	return "Hello World"
}

func TestHello(t *testing.T) {
	got := Hello("ZhangYundi")
	want := "Hello ZhangYundi"
	t.Logf("%[1]q, %[1]s", got)
	assert.Equal(t, want, got)
}

