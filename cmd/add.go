/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"github.com/l004p/3005-assignment-03/db"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		connection, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal(err)
		}
		defer connection.Close(ctx)
		query := db.New(connection)
		const layout = "2006-01-02"
		parsedTime, err := time.Parse(layout, args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(parsedTime)
		newStudent := &db.AddStudentParams{
			FirstName: args[0],
			LastName: args[1],
			Email: args[2],
			EnrollmentDate: pgtype.Date{
				Time: parsedTime,
				Valid: true,
			},
		}
		student, err := query.AddStudent(ctx, *newStudent)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("added student: %v\n", student)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
