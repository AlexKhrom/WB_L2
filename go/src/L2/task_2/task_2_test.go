package task_2

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	str    string
	resStr string
	resErr error
}

type funcRes struct {
	str string
	err error
}

func TestUnpackingString(t *testing.T) {
	tests := []TestStruct{
		{
			str:    "a4bc2d5e",
			resStr: "aaaabccddddde",
			resErr: nil,
		},
		{
			str:    "abcd",
			resStr: "abcd",
			resErr: nil,
		},
		{
			str:    "45",
			resStr: "",
			resErr: errors.New("некорректная строка"),
		},
		{
			str:    "a",
			resStr: "a",
			resErr: nil,
		},
		{
			str:    "v12",
			resStr: "vvvvvvvvvvvv",
			resErr: nil,
		},
		{
			str:    "Ê3Ä10",
			resStr: "ÊÊÊÄÄÄÄÄÄÄÄÄÄ",
			resErr: nil,
		},
		{
			str:    "a0",
			resStr: "",
			resErr: nil,
		},
		{
			str:    "0",
			resStr: "",
			resErr: errors.New("некорректная строка"),
		},
	}

	for _, test := range tests {
		resStr, err := unpackingString(test.str)

		assert.Equal(t, funcRes{str: test.resStr, err: test.resErr}, funcRes{str: resStr, err: err})
	}
}
