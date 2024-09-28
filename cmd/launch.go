package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	//"os/exec"
	"path/filepath"
	//"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vansh845/rover/internal"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Builds Image using Dockerfile and runs on your VPS.",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		path := filepath.Join(home, ".config/rover", "rover.yaml")
		viper.SetConfigFile(path)
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln(err)
		}
		var user string = viper.GetString("user")
		var host string = viper.GetString("host")
		var key string = viper.GetString("key")
		img := "nginx"
		command := exec.Command("sh", "-s", img, user, host, key)
		command.Stdin = strings.NewReader(internal.LoadToVps)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		err = command.Run()
		if err != nil {
			fmt.Printf("error running command : %q\n", err)
		}
		client := internal.NewSSHClient()
		cmds := []string{
			fmt.Sprintf("docker load -i %s.tar", img),
			fmt.Sprintf("docker run -p 80:80 %s", img),
		}
		err = internal.RunCmds(cmds, client)
		if err != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

}
