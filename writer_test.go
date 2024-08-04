package qcsv

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriter_Write(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		inputs [][]string
		wants  string
	}{
		{
			name: "case1",
			inputs: [][]string{
				{"hello", "world"},
			},
			wants: `"hello","world"`,
		},
		{
			name: "case2",
			inputs: [][]string{
				{"hello", `"my" world`},
			},
			wants: `"hello","""my"" world"`,
		},
		{
			name: "case3",
			inputs: [][]string{
				{"hello", "my \nworld"},
			},
			wants: "\"hello\",\"my \nworld\"",
		},
		{
			name: "case4",
			inputs: [][]string{
				{"hello", "", "world", ""},
			},
			wants: `"hello","","world",""`,
		},
		{
			name: "case5",
			inputs: [][]string{
				{"message1", "message2", "date"},
				{"hello", "world", `="20060102150406"`},
				{"hello", "world", `=""`},
			},
			wants: strings.Join([]string{
				`"message1","message2","date"`,
				`"hello","world","=""20060102150406"""`,
				`"hello","world","="""""`,
			}, "\n"),
		},
		{
			name: "case-all",
			inputs: [][]string{
				{"hello", "world"},
				{"hello", `"my" world`},
				{"hello", "my \nworld"},
				{"hello", "", "world", ""},
				{"hello", "world", `="20060102150406"`},
			},
			wants: strings.Join([]string{
				`"hello","world"`,
				`"hello","""my"" world"`,
				"\"hello\",\"my \nworld\"",
				`"hello","","world",""`,
				`"hello","world","=""20060102150406"""`,
			}, "\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotBuf := &bytes.Buffer{}
			w := NewWriter(gotBuf, `"`)
			if err := w.WriteAll(tt.inputs); err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.wants+"\n", gotBuf.String())
		})
	}
}
