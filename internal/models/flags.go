package models

import (
	"flag"

	"github.com/spf13/pflag"
)

var (
	address, user, password, database, charset, loc string
	port                                            uint16
	parseTime, debugLogLevel                        bool
)

var commandLine pflag.FlagSet

func init() {
	commandLine.StringVar(&address, "database-address", "localhost", "database address")
	commandLine.Uint16Var(&port, "database-port", 3306, "database port")
	commandLine.StringVar(&user, "database-user", "root", "database user")
	commandLine.StringVar(&password, "database-password", "", "database password")
	commandLine.StringVar(&database, "database-db-name", "webgame", "database name")
	commandLine.StringVar(&charset, "database-charset", "utf8", "database charset")
	commandLine.StringVar(&loc, "database-loc", "Local", "database parameter loc")
	commandLine.BoolVar(&parseTime, "database-parseTime", true, "database parameter parseTime")
	commandLine.BoolVar(&debugLogLevel, "gorm-debug", false, "set gorm log level to Debug")
}

// InitFlags is for explicitly initializing the flags.
// It may get called repeatedly for different flagSets, but not
// twice for the same one. May get called concurrently
// to other goroutines using models. However, only some flags
// may get set concurrently.
func InitFlags(flagSet *flag.FlagSet) {
	if flagSet == nil {
		flagSet = flag.CommandLine
	}
	commandLine.VisitAll(func(f *pflag.Flag) {
		flagSet.Var(f.Value, f.Name, f.Usage)
	})
}
