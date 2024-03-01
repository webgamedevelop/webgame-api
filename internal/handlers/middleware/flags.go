package middleware

import (
	"flag"
)

var inspectLevel int

var commandLine flag.FlagSet

func init() {
	commandLine.IntVar(&inspectLevel, "middleware-inspect-level", 3, "Middle ware inspect level")
}

// InitFlags is for explicitly initializing the flags.
// It may get called repeatedly for different flagSets, but not
// twice for the same one. May get called concurrently
// to other goroutines using middleware. However, only some flags
// may get set concurrently.
func InitFlags(flagSet *flag.FlagSet) {
	if flagSet == nil {
		flagSet = flag.CommandLine
	}
	commandLine.VisitAll(func(f *flag.Flag) {
		flagSet.Var(f.Value, f.Name, f.Usage)
	})
}
