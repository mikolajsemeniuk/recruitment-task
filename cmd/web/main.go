package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/mikolajsemeniuk/recruitment-task/pkg/docs"
	"github.com/mikolajsemeniuk/recruitment-task/pkg/index"
)

type config struct {
	Listen       string        `envconfig:"LISTEN"   required:"true"`
	Filepath     string        `envconfig:"FILEPATH" required:"true"`
	LogLevel     string        `default:"Info"       envconfig:"LOG_LEVEL"`
	ReadTimeout  time.Duration `default:"5s"         envconfig:"WRITE_TIMEOUT"`
	WriteTimeout time.Duration `default:"30s"        envconfig:"IDLE_TIMEOUT"`
	IdleTimeout  time.Duration `default:"30s"        envconfig:"IDLE_TIMEOUT"`
}

func main() {
	var c config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatal("error parsing config: %w", err)
	}

	options := &slog.HandlerOptions{Level: slog.LevelInfo}
	if c.LogLevel == "Debug" {
		options.Level = slog.LevelDebug
	}

	if c.LogLevel == "Error" {
		options.Level = slog.LevelError
	}

	// When dependency doesn't have to be injected like database connection
	// which usually has to be kept as singleton I prefer keeping less dependencies than more.
	// Logger for all components can be configured via SetDefault instead of passing it everywhere manually.
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, options)))

	store, err := index.NewMemory(os.DirFS("."), c.Filepath)
	if err != nil {
		log.Fatal("error creating datastore: %w", err)
	}

	router := http.NewServeMux()

	router.Handle("/", docs.NewHandler())
	router.Handle("/index/", http.StripPrefix("/index", index.NewHandler(store)))

	server := &http.Server{
		Addr:         c.Listen,
		Handler:      router,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		IdleTimeout:  c.IdleTimeout,
	}

	log.Printf("Server starting on %s", c.Listen)
	log.Fatal(server.ListenAndServe())
}
