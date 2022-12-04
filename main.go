package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ROOT_PASS string `yaml:"ROOT_PASS"`
	DB_NAME   string `yaml:"DB_NAME"`
	DB_USER   string `yaml:"DB_USER"`
	DB_PASS   string `yaml:"DB_PASS"`
	DB_PORT   string `yaml:"DB_PORT"`
	TZ        string `yaml:"TZ"`
}

func main() {
	fmt.Print("hello")

	// routing set
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// load env
	var fp string
	flag.StringVar(&fp, "c", ".env", "set yaml file path")
	flag.Parse()

	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	expaneded := os.ExpandEnv(string(b)) // here
	cfg := Config{}
	if err := yaml.Unmarshal([]byte(expaneded), &cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cfg)

	// shutDown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// context shotDown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
