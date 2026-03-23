package public

import "embed"

//go:embed *
var HomeFS embed.FS

//go:embed static/*
var StaticFS embed.FS
