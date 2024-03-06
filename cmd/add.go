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
	Short: "Add a student to the db",
	Long: `takes 4 parameters: first name, last name, email, and date in yyyy-mm-dd format`,
	Args: cobra.ExactArgs(4),
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
}
