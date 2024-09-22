package main

import (
	"github.com/eduardomassami/gopportunities/config"
	"github.com/eduardomassami/gopportunities/router"
)

var (
  logger *config.Logger
)

func main () {

  logger := config.GetLogger("main")

  err := config.Init()
  if err != nil {
    logger.Errorf("config initialization error: %v", err)
    return
  }
  router.Initialize()
  print("teste, teste, alguém na escuta?")
}
