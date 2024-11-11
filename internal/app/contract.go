package app

import "github.com/CaptainFallaway/XDH/internal"

type Contract interface {
	SelectAndLoadFile()
	GetGroupings() []internal.Grouping
}
