package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/katallaxie/pkg/logger"
	"github.com/spf13/cobra"
)

// Config ...
type Config struct {
	Flags *Flags
}

// Flags ...
type Flags struct {
	Addr string
}

var cfg = &Config{
	Flags: &Flags{},
}

var rootCmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(cmd.Context())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.Flags.Addr, "addr", ":8080", "addr")

	rootCmd.SilenceUsage = true
}

func run(ctx context.Context) error {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	logger.RedirectStdLog(logger.LogSink)

	app := fiber.New()
	app.Static("/", "./public")

	err := app.Listen(cfg.Flags.Addr)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
