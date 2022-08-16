package url

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	source := "http://google.com"
	expected := "Fy9z6aGA"

	id, err := Shorten(source)
	if err != nil {
		t.Fail()
	}

	if !assert.Equal(t, id, expected) {
		t.Fail()
	}
}
