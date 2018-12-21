package grpcgw

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fzerorubigd/balloon/pkg/assert"
	"github.com/fzerorubigd/balloon/pkg/log"
)

type security struct {
	In   string `json:"in"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type securityDefinitions struct {
	Security map[string]security
}

type swaggerFile struct {
	Swagger string `json:"swagger"`
	Info    struct {
		Title   string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`

	SecurityDefinitions securityDefinitions    `json:"security_definitions"`
	Host                string                 `json:"host"`
	Schemes             []string               `json:"schemes"`
	Consumes            []string               `json:"consumes"`
	Produces            []string               `json:"produces"`
	Paths               map[string]interface{} `json:"paths"`
	Definitions         map[string]interface{} `json:"definitions"`
}

var (
	swaggerLock sync.RWMutex
	data        = swaggerFile{
		Swagger: "2.0",
		Info: struct {
			Title   string `json:"title"`
			Version string `json:"version"`
		}{Title: "Balloon Swagger", Version: "1.0"},
		SecurityDefinitions: securityDefinitions{
			Security: map[string]security{
				"Authentication": {Name: "Authentication", In: "header", Type: "apiKey"},
			},
		},
		Schemes:     []string{"http", "https"},
		Consumes:    []string{"application/json"},
		Produces:    []string{"application/json"},
		Host:        address(),
		Paths:       make(map[string]interface{}),
		Definitions: make(map[string]interface{}),
	}
)

func address() string {
	p := strings.Split(addr.String(), ":")
	if len(p) != 2 {
		return addr.String()
	}
	if p[0] == "" {
		return "127.0.0.1:" + p[1]
	}

	return addr.String()
}

func swaggerHandler(w http.ResponseWriter, r *http.Request) {
	defer Recover(w)

	swaggerLock.RLock()
	defer swaggerLock.RUnlock()

	fl := strings.TrimPrefix(r.RequestURI, "/v1/swagger/")
	log.Info("Swagger file requested", log.String("file", fl))
	if fl == "index.json" {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Error("Failed to serve swagger", log.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	d, err := Asset(fl)
	if err != nil {
		log.Error("Not found", log.Err(err))
		w.WriteHeader(http.StatusNotFound)
	}
	switch strings.ToLower(filepath.Ext(fl)) {
	case ".json":
		w.Header().Add("Content-Type", "application/json")
	case ".css":
		w.Header().Add("Content-Type", "text/css")
	case ".js":
		w.Header().Add("Content-Type", "text/javascript")
	case ".html":
		w.Header().Add("Content-Type", "text/html")
	}
	_, _ = w.Write(d)
}

// RegisterSwagger register a swagger end point
func RegisterSwagger(paths map[string]interface{}, definitions map[string]interface{}) {
	swaggerLock.Lock()
	defer swaggerLock.Unlock()

	for i := range paths {
		_, ok := data.Paths[i]
		// TODO: Currently one path multiple method is not possible, fix it
		assert.False(ok, "Path is already registered", i)
		data.Paths[i] = appendSecurity(paths[i])
	}

	for i := range definitions {
		_, ok := data.Definitions[i]
		assert.False(ok, "Definition is already registered", i)
		data.Definitions[i] = definitions[i]
	}

}

func appendSecurity(d interface{}) map[string]interface{} {
	v := d.(map[string]interface{})
	for i := range v {
		meth, ok := v[i].(map[string]interface{})
		if !ok {
			continue
		}
		if meth["security"] != nil {
			if p, ok := meth["parameters"].([]interface{}); !ok {
				meth["parameters"] = createParameter(nil)
			} else {
				meth["parameters"] = createParameter(p)
			}
		}
	}
	return v

}

func createParameter(old []interface{}) []interface{} {
	return append(old, map[string]interface{}{
		"description": "the security token, get it from login route",
		"in":          "header",
		"name":        "authorization",
		"required":    true,
		"type":        "string",
	})
}
