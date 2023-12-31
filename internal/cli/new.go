package cli

import (
	"os"

	"github.com/colonyos/pollinator/pkg/project"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&ExecutorName, "executorname", "n", "", "Executor type")
	newCmd.MarkFlagRequired("executorname")
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new project",
	Long:  "Create a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetLevel(log.DebugLevel)
		}

		CheckError(checkIfDirIsEmpty("."))
		CheckError(checkIfDirExists("./cfs"))
		CheckError(checkIfDirExists("./cfs/src"))
		CheckError(checkIfDirExists("./cfs/data"))
		CheckError(checkIfDirExists("./cfs/result"))

		log.WithFields(log.Fields{
			"Dir": "./cfs/src"}).
			Info("Creating directory")
		err := os.MkdirAll("./cfs/src", 0755)
		CheckError(err)

		log.WithFields(log.Fields{
			"Dir": "./cfs/data"}).
			Info("Creating directory")
		err = os.MkdirAll("./cfs/data", 0755)

		CheckError(err)

		log.WithFields(log.Fields{
			"Dir": "./cfs/result"}).
			Info("Creating directory")
		err = os.MkdirAll("./cfs/result", 0755)

		CheckError(err)

		err = project.GenerateProjectConfig([]string{ExecutorName})
		CheckError(err)

		err = project.GenerateProjectData()
		CheckError(err)
	},
}
