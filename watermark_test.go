package main

import (
	"os"
	"testing"
)

func Test_testWaterMark(t *testing.T) {
	type args struct {
		in   string
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				in:   "git_book1.pdf",
				text: " 斗罗大陆 斗罗大陆 斗罗大陆 唐三 小舞 ",
			},
		}, {
			name: "test2",
			args: args{
				in:   "abc.pdf",
				text: "仅供投资人李武四4727查阅",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := testWaterMark(tt.args.in, tt.args.text)
			f, _ := os.Create(tt.args.in + "-out.pdf")
			f.Write(out)
			f.Close()
		})
	}
}
