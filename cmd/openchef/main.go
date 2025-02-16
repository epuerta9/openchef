package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/epuerta9/openchef/internal/config"
	"github.com/epuerta9/openchef/internal/database"
	"github.com/epuerta9/openchef/internal/nats"
	"github.com/epuerta9/openchef/internal/server"
	"github.com/epuerta9/openchef/internal/services/communicator"
	"github.com/epuerta9/openchef/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "openchef",
		Short: "OpenChef - AI Agent Orchestrator",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")

	// Add run command flags
	runCmd.Flags().Int("port", 8080, "API server port")
	runCmd.Flags().String("nats-url", "nats://localhost:4222", "NATS server URL")
	runCmd.Flags().String("db-url", "sqlite://local.db", "Database URL")
	runCmd.Flags().String("openai-key", "", "OpenAI API Key")

	viper.BindPFlag("PORT", runCmd.Flags().Lookup("port"))
	viper.BindPFlag("NATS_URL", runCmd.Flags().Lookup("nats-url"))
	viper.BindPFlag("DB_URL", runCmd.Flags().Lookup("db-url"))
	viper.BindPFlag("OPENAI_KEY", runCmd.Flags().Lookup("openai-key"))

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the OpenChef server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start embedded NATS server
		ns, err := nats.New()
		if err != nil {
			log.Fatalf("Failed to start NATS: %v", err)
		}
		defer ns.Shutdown(ctx)

		// Initialize DB
		db, err := database.New(viper.GetString("DB_URL"))
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Create communicator using embedded NATS URL
		comm, err := communicator.New(ns.URL())
		if err != nil {
			log.Fatalf("Failed to create communicator: %v", err)
		}
		defer comm.Close()

		// Create API server
		cfg := &config.Config{
			Port:        viper.GetInt("PORT"),
			OpenAIKey:   viper.GetString("OPENAI_KEY"),
			Environment: "development",
		}

		srv := server.New(cfg, db)

		// Handle shutdown gracefully
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigChan
			fmt.Println("\nShutting down gracefully...")
			cancel()
			srv.Shutdown(ctx)
		}()

		fmt.Printf("OpenChef %s starting...\n", version.Version)
		fmt.Printf("API server listening on :%d\n", cfg.Port)

		if err := srv.Start(); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OpenChef", version.Info())
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
