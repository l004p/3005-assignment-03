package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/l004p/3005-assignment-03/db"

	"github.com/jackc/pgx/v5"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get all the students",
	Long: `get the students`,
	Args: cobra.ExactArgs(0),
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
		students, err := query.GetAllStudents(ctx)
		if err != nil {
			log.Fatal(err)
		}
		for _, student := range students {
			fmt.Printf("%v\n", student)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
