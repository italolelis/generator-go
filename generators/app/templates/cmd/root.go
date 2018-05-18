package cmd

import (
	"context"

	"<%=projectRoot%>/pkg/config"
	"<%=projectRoot%>/pkg/log"
	"github.com/spf13/cobra"
)

type (
	// RootOptions represents the ahoy global options
	RootOptions struct {
		configFile string
		token      string
		org        string
		verbose    bool
	}
)

// NewRootCmd creates the root command
func NewRootCmd() *cobra.Command {
	opts := RootOptions{}

	cmd := cobra.Command{
		Use:   "capitain [command] [--flags]",
		Short: "Welcome to captain",
	}

	cmd.PersistentFlags().StringVarP(&opts.configFile, "config", "c", "", "config file")

	ctx := log.NewContext(context.Background())
	ctx, err := config.NewContext(ctx, opts.configFile)
	if err != nil {
		log.WithContext(ctx).Fatalf("Could not load configuration file: %v", err)
	}

	log.SetLevel(config.WithContext(ctx).LogLevel.Level())

	// Aggregates Root commands
	cmd.AddCommand(NewStartServerCmd(ctx))
	cmd.AddCommand(NewVersionCmd(ctx))

	return &cmd
}
