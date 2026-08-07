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

func main(){
	
	gpsLoggingPath := filepath.Join("/","var","lib","restfulpi","/gpsLogs")
	router := router.NewRouter(gpsLoggingPath)
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
