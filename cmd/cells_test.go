package cmd

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"testing"

	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
)

func TestCells(t *testing.T) {
	tests := []struct {
		arg1 sql.NullString
		arg2 sql.NullString
		outp string
		ext  exitcode.ExitCode
	}{
		{arg1: sql.NullString{String: "1", Valid: true}, arg2: sql.NullString{String: "1", Valid: true}, outp: "| TH |\n|----|\n| TD |", ext: exitcode.Normal},
		{arg1: sql.NullString{String: "2", Valid: true}, arg2: sql.NullString{String: "2", Valid: true}, outp: "| TH | TH |\n|----|----|\n| TD | TD |\n| TD | TD |", ext: exitcode.Normal},
		{arg1: sql.NullString{String: "1", Valid: true}, arg2: sql.NullString{Valid: false}, outp: "(cells): Wrong number of arguments", ext: exitcode.Normal},
		{arg1: sql.NullString{Valid: false}, arg2: sql.NullString{Valid: false}, outp: "(cells): Wrong number of arguments", ext: exitcode.Normal},
		{arg1: sql.NullString{String: "1", Valid: true}, arg2: sql.NullString{String: "hoge", Valid: true}, outp: "(cells): Argument format is incorrect", ext: exitcode.Normal},
		{arg1: sql.NullString{String: "1", Valid: true}, arg2: sql.NullString{String: "2hoge", Valid: true}, outp: "(cells): Argument format is incorrect", ext: exitcode.Normal},
	}

	for _, tt := range tests {
		inp := "cells " + tt.arg1.String + " " + tt.arg2.String
		r := strings.NewReader(inp)
		wbuf := &bytes.Buffer{}
		ebuf := &bytes.Buffer{}
		args := []string{"cells"}
		if tt.arg1.Valid {
			args = append(args, tt.arg1.String)
		}
		if tt.arg2.Valid {
			args = append(args, tt.arg2.String)
		}
		ext := Execute(
			rwi.New(
				rwi.WithReader(r),
				rwi.WithWriter(wbuf),
				rwi.WithErrorWriter(ebuf),
			),
			args,
		)
		if ext != tt.ext {
			t.Errorf("Execute() is %v\nwant \"%v\"", ext, tt.ext)
			fmt.Println(ebuf.String())
		}
		// 末尾の改行コードをトリミング
		str := strings.TrimRight(wbuf.String(), "\n")
		if str != tt.outp {
			t.Errorf("Execute() -> \n%v\nwant \n%v", str, tt.outp)
		}
	}
}
