/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day2aCmd represents the day2a command
var day2aCmd = &cobra.Command{
	Use:   "day2a",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("day_2.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		curLine := 0
		safeLine := 0
		for scanner.Scan() {
			line := scanner.Text()
			split := strings.Split(line, " ")
			state := ""
			safe := false

			// Safeline condition
			// All elements in the split must all be ascending or descending
			// All elements must differ by at least 1 and at most 3 from each other
			// If the above conditions are met, increment safeLine

			for i := 0; i < len(split)-1; i++ {
				// Check if the elemnent and the next element is the same. If it's the same it's not safe
				if split[i] == split[i+1] {
					safe = false
					break
				}
				a, _ := strconv.Atoi(split[i])
				b, _ := strconv.Atoi(split[i+1])
				if i == 0 {
					// Find out whether the elements are ascending or descending
					if a < b {
						state = "ascending"
					} else {
						state = "descending"
					}
				}

				if state == "ascending" {
					if a > b {
						fmt.Printf("%s | Breaking at %d Because %d and %d\n", state, curLine, a, b)
						safe = false
						break
					}
				} else {
					if a < b {
						fmt.Printf("%s | Breaking at %d Because %d and %d\n", state, curLine, a, b)
						safe = false
						break
					}
				}

				// Check if the difference between the elements is at least 1 and at most 3
				diff := int(math.Abs(float64(b - a)))
				if diff < 1 || diff > 3 {
					fmt.Printf("diff | Breaking at %d Because %d and %d = %d\n", curLine, a, b, diff)
					safe = false
					break
				}

				safe = true
			}

			if safe {
				safeLine++
			}
			curLine++
		}

		fmt.Println("Safe lines: ", safeLine)
	},
}

func init() {
	rootCmd.AddCommand(day2aCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2aCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2aCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
