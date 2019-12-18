package randomstr

import(
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestRandStringRunes(t *testing.T){
	str1 := RandStringRunes(-1)
	str2 := RandStringRunes(6)
  assert.EqualValues(t, str1, "")
	assert.EqualValues(t, 6, len(str2))
}
