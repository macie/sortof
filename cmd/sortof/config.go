package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"reflect"
	"time"
)

const helpMsg = "sortof - sort lines of text files\n" +
	"\n" +
	"Usage:\n" +
	"   sortof <algorithm> [-t <timeout>] [FILE...]\n" +
	"   sortof [-h] [-v]\n" +
	"\n" +
	"Options:\n" +
	"   -t <timeout>  timeout after which the program exits (default: 0).\n" +
	"                 Valid time units: ns, us, ms, s, m, h\n" +
	"   -h            show this help message and exit\n" +
	"   -v            show version information and exit\n" +
	"\n" +
	"Algorithms:\n" +
	"   bogo          Bogosort\n" +
	"   slow          Slowsort\n" +
	"   stalin        Stalinsort\n" +
	"\n" +
	"With no FILE, or when FILE is -, the command reads from standard input"

// AppVersion is the version of the program.
var AppVersion = "local-dev"

// AppConfig contains configuration options for the program provided by the user.
type AppConfig struct {
	SortFunc    func(ctx context.Context, file io.ReadCloser) ([]string, error)
	Files       []string
	Timeout     time.Duration
	ExitMessage string
}

// NewAppConfig creates a new AppConfig from the given command line arguments.
func NewAppConfig(cliArgs []string) (AppConfig, error) {
	config := AppConfig{}

	if len(cliArgs) == 0 {
		return AppConfig{}, fmt.Errorf(helpMsg)
	}

	// global options
	f := flag.NewFlagSet("global args", flag.ContinueOnError)
	f.SetOutput(io.Discard)
	showHelp := f.Bool("h", false, "")
	showVersion := f.Bool("v", false, "")
	if err := f.Parse(cliArgs); err != nil {
		return AppConfig{}, fmt.Errorf("%s. See 'sortof -h' for help", err)
	}
	if *showHelp {
		config.ExitMessage = helpMsg
		return config, nil
	}
	if *showVersion {
		config.ExitMessage = config.Version()
		return config, nil
	}

	// subcommand
	switch cliArgs[0] {
	case "bogo":
		config.SortFunc = BogosortFile
	case "slow":
		config.SortFunc = SlowsortFile
	case "stalin":
		config.SortFunc = StalinsortFile
	default:
		return config, fmt.Errorf("'%s' is not an algorithm. See 'sortof -h' for help", cliArgs[0])
	}

	// subcommand options
	s := flag.NewFlagSet("subcommand args", flag.ContinueOnError)
	s.SetOutput(io.Discard)
	s.DurationVar(&config.Timeout, "t", 0, "")
	showSubcommandHelp := s.Bool("h", false, "")
	if err := s.Parse(cliArgs[1:]); err != nil { // omit subcommand
		return AppConfig{}, fmt.Errorf("%s. See 'sortof -h' for help", err)
	}
	if *showSubcommandHelp {
		config.ExitMessage = helpMsg
		return config, nil
	}

	// files
	if len(s.Args()) > 0 {
		config.Files = s.Args()
	}

	return config, nil
}

// Equal reports whether two AppConfigs are equal. It is used in tests.
func (c AppConfig) Equal(other AppConfig) bool {
	return reflect.ValueOf(c.SortFunc).Pointer() == reflect.ValueOf(other.SortFunc).Pointer() &&
		reflect.DeepEqual(c.Files, other.Files) &&
		c.Timeout == other.Timeout &&
		c.ExitMessage == other.ExitMessage
}

// Version returns string with full version description.
func (c AppConfig) Version() string {
	build := ""
	if IsHardened {
		build = " (hardened)"
	}
	return fmt.Sprintf("sortof %s%s", AppVersion, build)
}

// NewAppContext returns a cancellable context which is:
// - cancelled when the interrupt signal is received
// - cancelled after the timeout (if any).
func NewAppContext(config AppConfig) (context.Context, context.CancelFunc) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	if config.Timeout != 0 {
		return context.WithTimeout(ctx, config.Timeout)
	}

	return ctx, cancel
}
