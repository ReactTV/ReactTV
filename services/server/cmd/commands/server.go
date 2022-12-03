package commands

import (
	log "github.com/sirupsen/logrus"
)

type Server struct {
	DBHost     string `required help:"Postgres hostname or IP address"`
	DBUser     string `required help:"Postgres username"`
	DBPassword string `required help"Postgres password"`
	DBName     string `required help"Postgres database name"`
	Port       string `required help:"Port the server listens on"`
}

func (s *Server) Run(globals *Globals) error {
	l := log.New()
	lvl, err := log.ParseLevel(globals.LogLevel)
	if err != nil {
		l.Fatal(err)
	}
	l.SetLevel(lvl)
	return nil
}
