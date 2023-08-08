package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lab18.matbea.xyz/savva/test-backend/internal/delivery/http/v1/chat"
)

type Cors struct {
    	AllowedMethods     []string 
			AllowedOrigins     []string 	
      AllowCredentials   bool     
      AllowedHeaders     []string 
			ExposedHeaders     []string
			Debug              bool
}


type handler struct {
  Engine *gin.Engine
  Cors Cors
}

type Router interface {
  RegisterRoutes(*gin.RouterGroup)
}

type NewHandlerParams struct {
  Cors Cors
}


func NewHttpHandler(params NewHandlerParams) *handler {
  engine := gin.New()

	return &handler{
    Cors: params.Cors,
    Engine: engine,
  }
}

func (h *handler) RegisterRoutes() {

  allowedOrigins := h.Cors.AllowedOrigins

  if h.Cors.Debug {
    allowedOrigins = []string{"*"};
  }

  h.Engine.Use(cors.New(cors.Config{
    AllowOrigins: allowedOrigins,
    AllowHeaders: h.Cors.AllowedHeaders,
    AllowMethods: h.Cors.AllowedMethods,
    ExposeHeaders: h.Cors.ExposedHeaders,
  }))

  mainGroup := h.Engine.Group("/api/v1")

  routers := []Router{
    chat.NewChatHandler(),
  }

  for _,v := range routers {
    v.RegisterRoutes(mainGroup)
  }
}
