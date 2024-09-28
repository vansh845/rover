package internal

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

func InitConfig() error {

	home, _ := os.UserHomeDir()
	configHome := fmt.Sprintf("%s/.config/rover/", home)
	configType := "yaml"
	configName := "rover"
	configPath := filepath.Join(configHome, configName+"."+configType)
	_, err := os.Stat(configHome)
	if err != nil {

		if os.IsNotExist(err) {
			err = os.MkdirAll(configHome, 0777)
			if err != nil {
				return err
			}
			_, err = os.Create(configPath)
			if err != nil {
				return err
			}

		}
		return err

	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configHome)

	var info = []string{"host", "user", "key"}

	for _, x := range info {
		var temp string
		fmt.Printf("Enter %s : ", x)
		fmt.Scan(&temp)
		viper.Set(x, temp)
	}
	err = viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func setConfig() *ssh.ClientConfig {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config/rover", "rover.yaml")
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	key := viper.GetString("key")
	user := viper.GetString("user")
	fmt.Println(key, user)
	fd, err := os.Open(key)
	if err != nil {
		log.Fatal(err)
	}
	pemBytes, err := io.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	cfg.SetDefaults()
	return cfg
}

func NewSSHClient() *ssh.Client {
	cfg := setConfig()

	host := viper.GetString("host")
	addr := fmt.Sprintf("%s:22", host)
	client, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		log.Fatalf("error connecting server , err : %q \n", err)
	}

	return client

}

func RunCmd(cmd string, client *ssh.Client) error {
	session, _ := client.NewSession()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	//if err != nil {
	//return fmt.Errorf("error while creating session %s : %q\n", cmd, err)
	//}
	err := session.Run(cmd)
	session.Close()
	return err
}

func RunCmds(cmds []string, client *ssh.Client) error {

	for _, cmd := range cmds {
		err := RunCmd(cmd, client)
		if err != nil {
			return err
		}
	}
	return nil
}
