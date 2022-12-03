package commands

import (
	"server/pkg/version"

	log "github.com/sirupsen/logrus"
)

type Version struct{}

func (s *Version) Run(globals *Globals) error {
	log.WithFields(log.Fields{
		"BuildTime": version.BuildTime,
		"Commit":    version.GitCommit,
		"Runtime":   version.Runtime(),
	}).Println(version.BinVersion)
	return nil
}
