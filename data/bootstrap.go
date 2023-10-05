package data

type BootstrapParams struct {
	Version string `name:"version"`
	Commit  string `name:"commit"`
}

type SystemConfig struct {
	Debug bool
}
