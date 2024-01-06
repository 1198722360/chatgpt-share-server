package api

import (
	"backend/openai/api/auth"

	"github.com/gogf/gf/v2/frame/g"
)

func init() {

	s := g.Server()
	apiGroup := s.Group("/api")
	apiGroup.GET("/auth/session", auth.Session)
}
