package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var ConfigPath string
var ResourceName string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&ConfigPath, "config", "c", "", "Config file path")
	getCmd.Flags().StringVarP(&ResourceName, "name", "n", "", "Resource name to access")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Find path for the request resource",
	Long:  `Return the path for a resource for the requested name`,
	Run: func(cmd *cobra.Command, args []string) {
		records := readCsvFile(ConfigPath)

		for _, record := range records {
			var name, location = record[0], record[1]
			if name == ResourceName {
				fmt.Println(location)
				return
			}
		}
	},
}
