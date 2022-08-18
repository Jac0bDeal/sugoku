package cli

import (
	"fmt"
	"github.com/alecthomas/kong"
)

// Globals contains the global flags for the CLI.
type Globals struct {
	LogLevel string      `short:"l" help:"Set the logging level (debug|info|warn|error|fatal)" default:"info"`
	Version  versionFlag `name:"version" help:"Print version information and quit"`
}

// versionFlag is the flag for getting version information.
type versionFlag string

func (v versionFlag) Decode(*kong.DecodeContext) error { return nil }
func (v versionFlag) IsBool() bool                     { return true }
func (v versionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(vars["version"])
	app.Exit(0)
	return nil
}

// commands is the set of commands on the App.
type commands struct {
	Globals
}

// App is the CLI application.
type App struct {
	Globals
	ctx *kong.Context
}

// NewApp constructs and returns an App from the passed env variables and flags.
func NewApp(version string) *App {
	cli := commands{
		Globals: Globals{
			Version: versionFlag(version),
		},
	}

	ctx := kong.Parse(&cli,
		kong.Name("sugoku"),
		kong.Description("A tool for generating and solving sudoku puzzles."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": version,
		})

	return &App{
		Globals: cli.Globals,
		ctx:     ctx,
	}
}

// Run runs the CLI application.
func (cli *App) Run() error {
	return cli.ctx.Run(&cli.Globals)
}
