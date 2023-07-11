/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// randomTestCmd represents the randomTest command
var randomTestCmd = &cobra.Command{
	Use:   "rt",
	Short: "Compare your solver and a brute-force solver using random test cases",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("randomTest called")

		for i := 1; i <= 10000; i++ {
			out, err := generate("./g.out")
			if err != nil {
				fmt.Println(err)
				return
			}
			o1, err := solve("./a.out", out)
			if err != nil {
				fmt.Println(err)
				return
			}
			o2, err := brute_force("./b.out", out)
			if err != nil {
				fmt.Println(err)
				return
			}
			if o1 != o2 {
				fmt.Printf("Output didn't match at testcase #%d\n", i)
			}
			if i%1000 == 0 {
				fmt.Printf("%d cases passed\n", i)
			}
		}

	},
}

func generate(generator string) (string, error) {
	out, err := exec.Command(generator).Output()
	return string(out), err
}

func solve(solver string, input string) (string, error) {
	cmd := exec.Command(solver)
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()

	return string(out), err
}

func brute_force(brute_forcer string, input string) (string, error) {
	cmd := exec.Command(brute_forcer)
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()

	return string(out), err
}

func init() {
	rootCmd.AddCommand(randomTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
