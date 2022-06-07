package requestid

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const defaultHeader = "X-Request-Id"

type Config struct {
	HeaderName string
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
	}
}

type Demo struct {
	next       http.Handler
	headerName string
	name       string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.HeaderName) == 0 {
		return nil, fmt.Errorf("HeaderName cannot be empty")
	}

	return &Demo{
		headerName: config.HeaderName,
		next:       next,
		name:       name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	id := uuid.New()
	req.Header.Set(a.headerName, id.String())
	a.next.ServeHTTP(rw, req)
}
