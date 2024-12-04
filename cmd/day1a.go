/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// day1aCmd represents the day1a command
var day1aCmd = &cobra.Command{
	Use:   "day1a",
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

		// Sort the array
		sort.Ints(leftArr)
		sort.Ints(rightArr)

		totalDist := 0
		fmt.Println("Total Length of the array is: ", len(leftArr))
		for i := 0; i < len(leftArr); i++ {
			dist := rightArr[i] - leftArr[i]
			dist = int(math.Abs(float64(dist)))
			totalDist += dist
		}
		fmt.Println("Total distance is: ", totalDist)
	},
}

func init() {
	rootCmd.AddCommand(day1aCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day1aCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day1aCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
