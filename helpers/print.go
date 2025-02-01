package helpers

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"tutorial.sqlc.dev/app/tutorial"
)

// Helper function to format the output as a table

func PrintAuthor(author *tutorial.Author) {
	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)

	// Set table headers
	table.SetHeader([]string{"Field", "Value"})

	// Add rows to the table
	table.Append([]string{"ID", fmt.Sprintf("%d", author.ID)})
	table.Append([]string{"Name", author.Name})

	// Handle the Bio field (check if it's valid)
	bio := "N/A"
	if author.Bio.Valid {
		bio = author.Bio.String
	}
	table.Append([]string{"Bio", bio})

	// Configure table styling for better readability
	table.SetBorder(true)                            // Add borders around the table
	table.SetCenterSeparator("|")                    // Use a separator for columns
	table.SetColumnSeparator("|")                    // Use a separator for columns
	table.SetRowSeparator("-")                       // Use a separator for rows
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT) // Align headers to the left
	table.SetAlignment(tablewriter.ALIGN_LEFT)       // Align all columns to the left
	table.SetAutoWrapText(true)                      // Wrap long text in cells

	// Render the table
	table.Render()
}

func PrintAuthors(authors []tutorial.Author) {
	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)

	// Set table headers
	table.SetHeader([]string{"ID", "Name", "Bio"})

	// Add rows to the table for each author
	for _, author := range authors {
		// Handle the Bio field (check if it's valid)
		bio := "N/A"
		if author.Bio.Valid {
			bio = author.Bio.String
		}

		// Append the author's details as a row
		table.Append([]string{
			fmt.Sprintf("%d", author.ID), // ID
			author.Name,                  // Name
			bio,                          // Bio
		})
	}

	// Configure table styling for better readability
	table.SetBorder(true)                            // Add borders around the table
	table.SetCenterSeparator("|")                    // Use a separator for columns
	table.SetColumnSeparator("|")                    // Use a separator for columns
	table.SetRowSeparator("-")                       // Use a separator for rows
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT) // Align headers to the left
	table.SetAlignment(tablewriter.ALIGN_LEFT)       // Align all columns to the left
	table.SetAutoWrapText(true)                      // Wrap long text in cells
	table.SetAutoFormatHeaders(true)                 // Format headers automatically

	// Render the table
	table.Render()
}
