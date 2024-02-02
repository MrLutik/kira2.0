package version

import (
	"fmt"
	"os"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
	"github.com/spf13/cobra"
)

const (
	use   = "version"
	short = "short description for version command"
	long  = "long description for version command"
)

func Version(log *logging.Logger) *cobra.Command {
	log.Debugln("Adding `version` command...")
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: short,
		Long:  long,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Fprintf(os.Stdout, "%v\n", types.KiraVersion)
		},
	}
	return versionCmd
}
