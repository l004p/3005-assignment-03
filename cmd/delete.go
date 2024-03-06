package cmd

import (
	//"fmt"
	"context"
	//"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/l004p/3005-assignment-03/db"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a student",
	Long: `deletes the student referred to by the student id`,
	Args: cobra.ExactArgs(1),
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
		studentId, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		err = query.DeleteStudent(ctx, int32(studentId))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
