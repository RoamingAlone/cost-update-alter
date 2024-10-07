package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the CSV file
	file, err := os.Open("new_cost.csv")
	if err != nil {
		log.Fatal("Unable to open CSV file", err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read CSV file", err)
	}

	// Open a text file for writing the generated SQL queries
	outputFile, err := os.Create("new_cost_queries.txt")
	if err != nil {
		log.Fatal("Unable to create output file", err)
	}
	defer outputFile.Close()

	// Write SQL queries into the .txt file
	for i, record := range records[1:] { // Skip header row (index 0)
		// Get the cost (Column D) and SKU (Column E)

		// Remove commas from the cost string
		costStr := strings.ReplaceAll(record[3], ",", "")
		sku := record[4]

		// Generate the SQL query
		query := fmt.Sprintf("UPDATE store_product SET cost = '%s' WHERE sku = '%s';\n", costStr, sku)

		// Write the query into the .txt file
		_, err := outputFile.WriteString(query)
		if err != nil {
			log.Printf("Error writing query for row %d: %v", i+2, err)
		}
	}

	fmt.Println("SQL update queries generated successfully in 'new_cost_queries.txt'")
}
