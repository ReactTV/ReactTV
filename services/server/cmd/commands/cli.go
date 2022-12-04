package commands

const (
	Name = "server"
)

type Globals struct {
	LogLevel  string `enum:"panic,fatal,error,warn,info,debug,trace" default:"info"`
	LogFormat string `enum:"text,json" default:"text"`
}

type DataMangementCLI struct {
	Globals
	Server  Server  `cmd:"" help:"Start Server with the given args"`
	Version Version `cmd:"" help:"Current Server version"`
}
