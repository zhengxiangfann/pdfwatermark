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
				in:   "tmp_1651997191.pdf",
				text: " 斗罗大陆 斗罗大陆 斗罗大陆 唐三 小舞 ",
			},
		}, {
			name: "test2",
			args: args{
				in:   "b.pdf",
				text: "小舞",
			},
		}, {
			name: "test3",
			args: args{
				in:   "c.pdf",
				text: "波塞西",
			},
		}, {
			name: "test4",
			args: args{
				in:   "abc.pdf",
				text: "胡列娜",
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
