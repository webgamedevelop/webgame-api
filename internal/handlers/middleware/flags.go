package middleware

import (
	"flag"
	"time"
)

var (
	inspectLevel             int
	cookieName, cookieDomain string
	cookieHTTPOnly           bool
	timeout, maxRefresh      time.Duration
)

var commandLine flag.FlagSet

func init() {
	commandLine.IntVar(&inspectLevel, "middleware-inspect-level", 2, "Middle ware inspect level")
	commandLine.StringVar(&cookieName, "middleware-cookie-name", "token", "Cookie name")
	commandLine.StringVar(&cookieDomain, "middleware-cookie-domain", "", "Cookie domain")
	commandLine.BoolVar(&cookieHTTPOnly, "middleware-cookie-HTTPOnly", true, "Cookie HTTPOnly")
	commandLine.DurationVar(&timeout, "middleware-token-timeout", time.Hour, "Token timeout")
	commandLine.DurationVar(&maxRefresh, "middleware-token-maxRefresh", time.Hour, "Token max refresh time")
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
