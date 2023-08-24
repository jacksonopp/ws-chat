package main

import (
	"log"
	"net/http"

	"github.com/jacksonopp/ws-chat/server"
	"github.com/jacksonopp/ws-chat/ui"
	"golang.org/x/net/websocket"
)

const port = ":8080"

func main() {
	assets, _ := ui.Assets()
	server := server.NewServer()

	mux := http.NewServeMux()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	mux.Handle("/", http.StripPrefix("/", fs))
	mux.Handle("/ws", websocket.Handler(server.HandleWs))

	// Serve the files using the default HTTP server
	log.Printf("Listening on %s...", port)

	// url := fmt.Sprintf("http://localhost%s", port)

	// // open user's browser to login page
	// if err := browser.OpenURL(url); err != nil {
	// 	log.Fatalf("failed to open browser for url %s", err.Error())
	// }

	if err := http.ListenAndServe(port, mux); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
