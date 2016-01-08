package presentation

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	assert := assert.New(t)
	out := new(bytes.Buffer)

	r := Renderer{}
	r.Image(out, []byte("link"), []byte{}, []byte("alt_to_class"))

	b, err := ioutil.ReadAll(out)
	assert.NoError(err)

	assert.Equal(`<img src="link" class="alt_to_class">`, string(b))
}
