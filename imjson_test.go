package goimjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImJSON(t *testing.T) {
	start := []byte("{}")

	imj, err := NewWithBody(start)
	assert.NoError(t, err)

	ver1 := imj.Set("field1", "value1")
	ver2 := imj.Set("field2", "value2")

	d, err := imj.Encode()
	assert.NoError(t, err)
	assert.Equal(t, "{\"field1\":\"value1\",\"field2\":\"value2\"}", string(d))

	imjI := imj.Interface()
	imjV := imjI.(map[string]interface{})
	assert.Equal(t, "value1", imjV["field1"])
	assert.Equal(t, "value2", imjV["field2"])

	imj1, err := imj.Get(ver1, "field1")
	assert.NoError(t, err)
	assert.Equal(t, "value1", imj1.Interface())

	imj12, err := imj.Get(ver1, "field2")
	assert.NoError(t, err)
	assert.Equal(t, nil, imj12.Interface())

	imj2, err := imj.Get(ver2, "field2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", imj2.Interface())

	imjN, err := imj.GetLatest("field2")
	assert.NoError(t, err)
	assert.Equal(t, "value2", imjN.Interface())

}
