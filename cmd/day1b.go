/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day1bCmd represents the day1b command
var day1bCmd = &cobra.Command{
	Use:   "day1b",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var leftArr []int
		var rightArr []int
		file, err := os.Open("day_1.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			res := strings.Split(line, " ")
			leftNum, err := strconv.Atoi(res[0])
			if err != nil {
				panic(err)
			}
			leftArr = append(leftArr, leftNum)

			rightNum, err := strconv.Atoi(res[len(res)-1])
			if err != nil {
				panic(err)
			}
			rightArr = append(rightArr, rightNum)
		}

		memory := make(map[int]int)
		sum := 0
		for i := 0; i < len(leftArr); i++ {
			key := leftArr[i]
			// If key is not present in memory, count the appearence of key in rightArr
			if _, ok := memory[key]; !ok {
				memory[key] = 0
				for j := 0; j < len(rightArr); j++ {
					if rightArr[j] == key {
						memory[key]++
					}
				}

				memory[key] = memory[key] * key
			}

			sum += memory[key]
		}
		println("Sum is: ", sum)
	},
}

func init() {
	rootCmd.AddCommand(day1bCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1bCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1bCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
