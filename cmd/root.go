package main

import (
	"context"
	"encoding/json"
	"fmt"
	"httpproj1/apis"
	"httpproj1/initializers"
	"httpproj1/logger"
	"httpproj1/migrations"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var myslog *slog.Logger = logger.GetLogger()

func run() {
	// myslog := logger.GetLogger()
	// http.HandleFunc("/", apis.GetRoot)
	// http.HandleFunc("/create_product", apis.CreateProduct)

	myslog.Info("Start Listening to 8080")
	servErr := http.ListenAndServe(":8080", nil)
	if servErr != nil {
		myslog.Error("Somthing is Wrong!!!")
	}
}

func main() {
	var cmdMigrate = &cobra.Command{
		Use:   "migrate", //important command
		Short: "Run Migrations",
		Long:  "Run Migrations",
		// Args:cobra.MinimumNArgs(1)
		Run: func(cmd *cobra.Command, args []string) {
			migrations.RunMigrations()
		},
	}
	var cmdRun = &cobra.Command{
		Use:   "runapp", //important command
		Short: "Run the app",
		Long:  "Run the app",
		// Args:cobra.MinimumNArgs(1)
		Run: func(cmd *cobra.Command, args []string) {
			migrations.SetUp()
			router := apis.GetRouter()
			if err := router.Start(":8080"); err != nil {
				myslog.Error(err.Error())
			}

			// Wait for interrupt signal
			quit := make(chan os.Signal, 1)                    // Create a channel to receive OS signals
			signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // Listen for interrupt (Ctrl+C) or termination signals
			<-quit                                             // Wait for the signal to be received

			// Graceful shutdown of the router
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Set a timeout for graceful shutdown
			defer cancel()                                                           // Ensure cancellation of the context when done                                                      // Ensure cancellation of the context when done
			if err := router.Shutdown(ctx); err != nil {                             // Try shutting down the router
				myslog.Error("Server Shutdown Failed: " + err.Error()) // Log an error if shutdown fails
			}

		},
	}
	var showConf = &cobra.Command{
		Use:   "showconf",
		Short: "show configs",
		Long:  "Show Configs of app",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := initializers.LoadConfig(".")
			if err != nil {
				fmt.Println("Error in Load Config File!")
			} else {
				jsoned_config, err := json.MarshalIndent(config, "", "  ")
				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println(string(jsoned_config))

			}
		},
	}

	var rootCmd = &cobra.Command{Use: "jaziapp"}
	rootCmd.AddCommand(cmdRun, cmdMigrate, showConf)
	rootCmd.Execute()
}
