package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cobrich/url-shortener/handler"
	"github.com/cobrich/url-shortener/storage"
)

func main() {

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	port := ":8080"

	storage := storage.NewStorage()
	handler := handler.NewHandler(storage)

	router := http.NewServeMux()
	router.HandleFunc("GET /{short_code}", handler.GetLongURLHundler)
	router.HandleFunc("POST /shorten", handler.CreateShortURLHundler)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		log.Print("Server running on localhost port:", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	<-quit
	log.Println("Получен сигнал на завершение работы. Начинаем graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}

	log.Println("Все соединения закрыты.")

	log.Println("Сервер успешно остановлен.")
}
