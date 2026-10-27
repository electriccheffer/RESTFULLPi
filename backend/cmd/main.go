package main
import "path/filepath"
import "net/http"
import "time"
import "log"
import "errors"
import "os/signal"
import "syscall"
import "os"

import "restfulpi/internal/server"
import "restfulpi/internal/handler"
import "restfulpi/internal/file_operations"

func main(){
	
	homeDir,err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v", err)
	}
	gpsLoggingPath := filepath.Join(homeDir, "gpsLogs")
	readPath := filepath.Join("/dev","ttyACM0")
	fileOpener := file_operations.NewFileOpener()	
	sessionManager := handler.NewSessionManager(gpsLoggingPath,readPath,fileOpener)
	router := router.NewRouter(sessionManager)
	port := "8080" 
	httpServer := &http.Server{
			Addr: ":" + port,
			Handler: router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second, 
		 }
	go func() {
		log.Printf("Server listening on: %s",port)
		if err := httpServer.ListenAndServe(); err != nil && 
			!errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("HTTP server error: %v", err)
		}

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")		
}
