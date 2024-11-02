package foo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo1(t *testing.T) {
	expected := "Foo"
	actual := Foo()

	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
	//assert from testify module
	assert.Equal(t, "Foo", Foo(), "they should be equal")

}

//go test -cover
//go test -coverprofile profile
//go tool cover -html=profile
