/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

func newCellsCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cellsCmd := &cobra.Command{
		Use:     "cells",
		Aliases: []string{"ce", "c"},
		Short:   "render tables markdown",
		Long:    "render tables markdown (detail)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return renderMdTable(ui, cmd, args)
		},
	}
	cellsCmd.SetArgs(args) //arguments of command-line
	return cellsCmd
}

func renderMdTable(ui *rwi.RWI, cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		if err := ui.Outputln("(cells): Wrong number of arguments"); err != nil {
			return err
		}
		return nil
	}

	// 引数で指定できるのは1以上の数値のみ
	result := ""
	column := ""
	cc := 0
	partition := ""
	row := ""
	for i, v := range args {
		rex := regexp.MustCompile(`^[0-9]{1,}$`)
		if !rex.MatchString(v) || v == "0" {
			if err := ui.Outputln("(cells): Argument format is incorrect"); err != nil {
				return err
			}
			return nil
		}
		num, _ := strconv.Atoi(v)
		switch i {
		case 0:
			column += "|"
			for n := 0; n < num; n++ {
				column += " TH |"
				cc++
			}
		case 1:
			partition += "|"
			for n := 0; n < cc; n++ {
				partition += "----|"
			}
			for n := 0; n < num; n++ {
				row += "|"
				for m := 0; m < cc; m++ {
					row += " TD |"
				}
				row += "\n"
			}
			row = strings.TrimRight(row, "\n")

		default:
			return nil
		}
		result = column + "\n" + partition + "\n" + row
	}

	if err := ui.Outputln(result); err != nil {
		return err
	}
	return nil
}
