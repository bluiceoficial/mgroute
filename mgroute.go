// Copyright (C) 2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT
//
// Site: https://mugomes.github.io

package mgroute

import (
	"net/url"
	"regexp"
	"slices"
	"strings"
)

type MGRoute struct {
	fullURL  string
	urlParts []string
	error404 bool
}

func New(requestURI string) *MGRoute {
	path := "/"

	if requestURI != "" {
		if sURL, err := url.Parse(requestURI); err == nil {
			path = strings.TrimRight(sURL.Path, "/")
		}
	}

	if path == "" {
		path = "/"
	}

	parts := strings.Split(path, "/")
	parts = slices.DeleteFunc(parts, func(value string) bool {
		return value == ""
	})

	return &MGRoute{
		fullURL:  path,
		urlParts: parts,
		error404: true,
	}
}

// Retorna as partes da URL
func (r *MGRoute) GetArrayURLs() []string {
	return r.urlParts
}

// Retorna a URL completa
func (r *MGRoute) GetFullURL() string {
	return strings.TrimLeft(r.fullURL, "/")
}

// Verifica se a URL existe
func (r *MGRoute) CheckURL(pattern string) bool {
	re := regexp.MustCompile("^" + pattern + "$")
	return re.MatchString(r.fullURL)
}

// Retorna a parte da URL pelo índice
func (r *MGRoute) GetURL(index int) string {
	if index < 0 || index >= len(r.urlParts) {
		return ""
	}
	return r.urlParts[index]
}

// Retorna a primeira parte da URL
func (r *MGRoute) GetFirstURL() string {
	return r.GetURL(0)
}

// Retorna a penúltima parte da URL
func (r *MGRoute) GetPenultimateURL() string {
	if len(r.urlParts) < 2 {
		return ""
	}
	return r.urlParts[len(r.urlParts)-2]
}

// Retorna a última parte da URL
func (r *MGRoute) GetLastURL() string {
	if len(r.urlParts) == 0 {
		return ""
	}
	return r.urlParts[len(r.urlParts)-1]
}

// Aplica o conceito funcional
func (r *MGRoute) GetPart(pattern string, function func(args ...string)) {
	sURL := "/"
	if r.fullURL != "" {
		sURL = r.fullURL
	}

	re := regexp.MustCompile("^" + pattern + "$")
	matches := re.FindStringSubmatch(sURL)

	if len(matches) > 0 {
		// remove o match completo
		function(matches[1:]...)
		r.error404 = false
	}
}

// Retorna erro caso não for encontrado o pattern
func (r *MGRoute) GetError(handler func()) {
	if r.error404 {
		handler()
	}
}
