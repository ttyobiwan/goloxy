package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func getAuthorizeURL(s *Settings, redirectURL string) *url.URL {
	callbackURL := fmt.Sprintf("%s/callback?redirect_url=%s", s.AppDomain, redirectURL)

	params := url.Values{}
	params.Set("client_id", s.AuthClientId)
	params.Set("response_type", "code")
	params.Set("audience", s.AuthAudience)
	params.Set("redirect_uri", callbackURL)

	return &url.URL{
		Scheme:   "https",
		Host:     s.AuthDomain,
		Path:     "/authorize",
		RawQuery: params.Encode(),
	}
}

func getToken(code string, s *Settings) (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", s.AuthDomain)
	body := map[string]string{
		"client_id":     s.AuthClientId,
		"client_secret": s.AuthClientSecret,
		"audience":      s.AuthAudience,
		"redirect_uri":  s.AppDomain,
		"grant_type":    "authorization_code",
		"code":          code,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	bodyReader := bytes.NewReader(jsonBody)
	response, err := http.Post(url, "application/json", bodyReader)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	data := map[string]interface{}{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	return data["access_token"].(string), nil
}

func getLogoutURL(redirectURL string, s *Settings) *url.URL {
	params := url.Values{}
	params.Set("client_id", s.AuthClientId)
	params.Set("returnTo", redirectURL)

	return &url.URL{
		Scheme:   "https",
		Host:     s.AuthDomain,
		Path:     "/logout",
		RawQuery: params.Encode(),
	}
}
