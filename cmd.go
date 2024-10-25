package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/victormazeli/Fast-Track-Interview-Test/proto"
	"github.com/victormazeli/Fast-Track-Interview-Test/server"
	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "start-quiz",
	Short: "Start the quiz and answer questions one by one",
	Run: func(cmd *cobra.Command, _ []string) {

		go func() {
			server.NewServer(cmd.Context(), port)
		}()

		time.Sleep(2 * time.Second)

		go func() {
			<-cmd.Context().Done()
			log.Println("Quiz interrupted, shutting down gracefully...")
		}()

		if err := InitClient(port); err != nil {
			log.Printf("Failed to initialize client: %v\n", err)
			return
		}

		res, err := client.GetQuestions(cmd.Context(), &pb.NoRequestParam{})
		if err != nil {
			log.Printf("Failed to get questions: %v\n", err)
			return
		}

		if res.GetQuestions() == nil || len(res.GetQuestions()) == 0 {
			fmt.Println("No questions received from the server.")
			return
		}
		shuffledQuestions := shuffleQuestions(res.GetQuestions())
		questions = shuffledQuestions

		getNextQuestion(cmd.Context())
	},
}


func StartCLI(ctx context.Context) error {
	fmt.Println("Starting Quiz...")
	rootCmd := &cobra.Command{Use: "quiz-cli"}
	quizCmd.PersistentFlags().StringVarP(&port, "port", "p", "50051", "Server port")
	rootCmd.AddCommand(quizCmd)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
