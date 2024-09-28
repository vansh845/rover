package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vansh845/rover/internal"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("initializing rover🚀...")
		//setup config
    err := internal.InitConfig()
    if err != nil{
      log.Fatalln("error occured while initializing config: ",err)
    }
		client := internal.NewSSHClient()
    user := viper.GetString("user")
		//install docker
		dockerSteps := []string{
			"sudo apt-get update -y",
			"sudo apt-get install -y ca-certificates curl",
			"sudo install -m 0755 -d /etc/apt/keyrings", // Removed '-y', it's not valid for `install`
			"sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc", // Removed '-y', it's not valid for `curl`
			"sudo chmod a+r /etc/apt/keyrings/docker.asc",
			`echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null`,
			"sudo apt-get update -y",
			"sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin",
			"docker --version", // Use `docker --version` to verify Docker installation
      fmt.Sprintf("sudo usermod -aG docker %s",user),
		}
		err = internal.RunCmds(dockerSteps, client)
		if err != nil {
			log.Fatalln(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
