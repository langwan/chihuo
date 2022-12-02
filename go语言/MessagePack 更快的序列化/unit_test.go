package msgpack

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
	"testing"
)

func Test_MsgpackMarshal(t *testing.T) {
	marshal, err := msgpack.Marshal(msg2000)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}

func Test_JsonMarshal(t *testing.T) {
	marshal, err := json.Marshal(msg2000)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}
func Test_YamlMarshal(t *testing.T) {
	marshal, err := yaml.Marshal(msg2000)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}

func Test_MsgpackMarshalMessages(t *testing.T) {
	marshal, err := msgpack.Marshal(messages)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}

func Test_JsonMarshalMessages(t *testing.T) {
	marshal, err := json.Marshal(messages)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}
func Test_YamlMarshalMessages(t *testing.T) {
	marshal, err := yaml.Marshal(messages)
	assert.NoError(t, err)
	t.Logf("len = %d", len(marshal))
}
