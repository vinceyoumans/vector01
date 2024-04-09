package main

import (
	"fmt"
	"log"
	slog "log/slog"
	"net"
	"os"
	"time"

	slogmulti "github.com/samber/slog-multi"
)

func main() {

	// Define log parameters
	v1 := "10"
	v2 := "bada"

	// Create a new logger instance
	// logger := slog.New(slog.Writer(os.Stdout), slog.Level(slog.Info))
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger01 := slog.New(slog.NewTextHandler(os.Stderr, nil))

	// Log the message
	logger01.Info("message", slog.String("k1", v1), slog.String("k2", v2))

	// textHandler := slog.NewTextHandler(os.Stdout)
	// logger := slog.New(textHandler)

	logger01.Info("Usage Statistics",
		slog.Int("current-memory", 50),
		slog.Int("min-memory", 20),
		slog.Int("max-memory", 80),
		slog.Int("cpu", 10),
		slog.String("app-version", "v0.0.1-beta"),
	)

	log.Println("--------    dailing out. --------------")

	logstash, _ := net.Dial("tcp", "localhost:80")
	stderr := os.Stderr

	logger := slog.New(
		slogmulti.Fanout(
			slog.NewJSONHandler(logstash, &slog.HandlerOptions{}), // pass to first handler: logstash over tcp
			slog.NewTextHandler(stderr, &slog.HandlerOptions{}),   // then to second handler: stderr
			// ...
		),
	)

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now()),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")

}
