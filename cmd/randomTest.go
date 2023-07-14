/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

type RtConfig struct {
	Generator  string `json:"generator"`
	Solver     string `json:"solver"`
	BruteForce string `json:"brute_force"`
	CasesNum   int    `json:"cases_num"`
}

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

		rtConfig := loadConfig("config.json")

		fmt.Println("Running random tests...")

		for i := 1; i <= rtConfig.CasesNum; i++ {
			out, err := generate("./" + rtConfig.Generator)
			if err != nil {
				fmt.Println(err)
				return
			}
			o1, err := solve("./"+rtConfig.Solver, out)
			if err != nil {
				fmt.Println(err)
				return
			}
			o2, err := brute_force("./"+rtConfig.BruteForce, out)
			if err != nil {
				fmt.Println(err)
				return
			}

			if o1 != o2 {

				color.Red.Printf("Output didn't match at testcase #%d\n", i)
				fmt.Printf("Testcase:\n%s\n", out)
				fmt.Printf("Solver:\n%s\n", o1)
				fmt.Printf("Brute force:\n%s\n", o2)

				return
			}
			if i%1000 == 0 {
				fmt.Printf("%d cases passed\n", i)
			}
		}
		color.Green.Println("No mismatch detected.")
	},
}

func loadConfig(config string) RtConfig {
	defaultConfig := RtConfig{
		Generator:  "generator.out",
		Solver:     "solver.out",
		BruteForce: "brute_force.out",
		CasesNum:   10000,
	}
	_, err := os.Stat(config)
	if err != nil {
		fmt.Println("Could not find config.json, using default config")
		return defaultConfig
	}

	file, err := os.Open("config.json")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error opening config.json, using default config")
		return defaultConfig
	}
	var rtConfig RtConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&rtConfig)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error parsing config.json, using default config")
		return defaultConfig
	}
	return RtConfig{
		Generator:  rtConfig.Generator,
		Solver:     rtConfig.Solver,
		BruteForce: rtConfig.BruteForce,
		CasesNum:   rtConfig.CasesNum,
	}
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
