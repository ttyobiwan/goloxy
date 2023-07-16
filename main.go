package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type ContextKey string

type Handler struct {
	settings *Settings
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	// Create a new request with the updated context
	ctx := context.WithValue(r.Context(), ContextKey("settings"), h.settings)
	r = r.WithContext(ctx)

	// Call the default ServeHTTP to process the request
	http.DefaultServeMux.ServeHTTP(w, r)

	// Log request details
	took := time.Since(start)
	log.WithFields(
		log.Fields{
			"method":         r.Method,
			"uri":            r.RequestURI,
			"remote_address": r.RemoteAddr,
			"took":           took.String(),
		},
	).Info(
		fmt.Sprintf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			took,
		),
	)
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value(ContextKey("settings")).(*Settings)
	redirectURL := r.URL.Query().Get("redirect_url")
	authorizeURL := getAuthorizeURL(settings, redirectURL).String()
	http.Redirect(w, r, authorizeURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value(ContextKey("settings")).(*Settings)
	token, err := getToken(r.URL.Query().Get("code"), settings)
	redirectURL := r.URL.Query().Get("redirect_url")
	if err == nil {
		redirectURL = fmt.Sprintf("%s?access_token=%s", redirectURL, token)
	} else {
		log.WithFields(log.Fields{"error": err}).Error("Callback failed")
		redirectURL = fmt.Sprintf("%s?error=%s", redirectURL, err)
	}
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func redirectToLogout(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value(ContextKey("settings")).(*Settings)
	redirectURL := r.URL.Query().Get("redirect_url")
	logoutURL := getLogoutURL(redirectURL, settings)
	http.Redirect(w, r, logoutURL.String(), http.StatusFound)
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	handler := Handler{NewSettings()}
	server := http.Server{
		Addr:         ":8080",
		Handler:      &handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	http.HandleFunc("/login", redirectToLogin)
	http.HandleFunc("/callback", handleCallback)
	http.HandleFunc("/logout", redirectToLogout)

	log.Info("Server listening on 0.0.0.0:8080")
	log.Fatal(server.ListenAndServe())
}
