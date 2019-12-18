package randomstr

import(
  "testing"
  "os"
  "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M){
  os.Exit(m.Run())
}

func TestRandStringRunesMinusLength(t *testing.T){
	str1 := RandStringRunes(-1)
  assert.EqualValues(t, str1, "")
}

func TestRandStringRunesValidLength(t *testing.T){
  str2 := RandStringRunes(6)
  assert.EqualValues(t, 6, len(str2))
}
