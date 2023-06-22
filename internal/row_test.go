package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidateRow(t *testing.T) {
	type Some struct {
		Kind string `column:"enum" enum:"a,b,c,d"`
	}

	t.Run("valid", func(t *testing.T) {
		some := Some{
			Kind: "a",
		}
		err := ValidateRow(&some)
		require.NoError(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		some := Some{
			Kind: "z",
		}
		err := ValidateRow(&some)
		require.Error(t, err)
	})
}

func TestRowValues(t *testing.T) {
	type Issue struct {
		Id       uint64
		Title    string
		Kind     string `column:"enum" enum:"a,b,c,d"`
		Context  string
		Contents string
		State    string `column:"enum" enum:"a,b,c,d"`
	}
	issue := Issue{
		Id:       1,
		Title:    "a",
		Kind:     "b",
		Context:  "c",
		Contents: "d",
		State:    "e",
	}

	ret := RowValues(&issue)

	require.Equal(t, []string{"1", "a", "b", "c", "d", "e"}, ret)
}
