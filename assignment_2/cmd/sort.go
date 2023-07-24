package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"sort"

	"strconv"

	"strings"
)

var intType, stringType bool

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort element",
	Long:  "Sort element by string or integer",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide element to be sort")
			return
		}

		if intType && stringType {
			fmt.Println("Not support mix type yet")
		} else if intType {
			intSlice := make([]int, len(args))
			for i, elem := range args {
				num, err := strconv.Atoi(strings.TrimSpace(elem))
				if err != nil {
					fmt.Printf("Error converting '%s' to an integer: %v\n", elem, err)
					return
				}
				intSlice[i] = num
			}

			sort.Ints(intSlice)

			fmt.Println("Sorted integers:", intSlice)
		} else if stringType {
			sort.Strings(args)

			fmt.Println("Sorted strings:", args)
		} else {
			fmt.Println("Invalid sort type. Use 'int' or 'string'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)

	sortCmd.Flags().BoolVarP(&intType, "integer", "i", false, "Sort by integer")
	sortCmd.Flags().BoolVarP(&stringType, "string", "s", false, "Sort by string")
}
