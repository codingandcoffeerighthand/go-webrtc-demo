package configs

import _ "embed"

//go:embed config.yaml
var DefaultConfig []byte
