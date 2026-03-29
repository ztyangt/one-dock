package public

import "embed"

//go:embed sources/*
var SourcesFS embed.FS

//go:embed dist/*
var DistFS embed.FS
