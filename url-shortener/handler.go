package main

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := urlMap[path]; ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	urlMap := map[string]string{}
	for _, pathURL := range pathUrls {
		urlMap[pathURL.Path] = pathURL.URL
	}
	return MapHandler(urlMap, fallback), nil
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
