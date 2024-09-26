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

func InitConfig() {

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
				log.Fatalln(err)
			}
			_, err = os.Create(configPath)
			if err != nil {
				log.Fatalln(err)
			}

		} else {
			log.Fatalln(err)
		}

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
		log.Fatal(err)
	}
}

func setConfig() *ssh.ClientConfig {
	key := viper.GetString("key")
	user := viper.GetString("user")
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

func NewSSHSession() *ssh.Session {
	cfg := setConfig()

	host := viper.GetString("host")
	addr := fmt.Sprintf("%s:22", host)
	client, err := ssh.Dial("tcp", addr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	return session

}

func RunCmd(cmd string, session *ssh.Session) error {
	err := session.Run(cmd)
	if err != nil {
		log.Printf("error while running %s : %q", cmd, err)
		return err
	}
	return nil

}

func RunCmds(cmds []string, session *ssh.Session) error {

	for _, cmd := range cmds {
		err := RunCmd(cmd, session)
		if err != nil {
			log.Fatalln(err)
			return err
		}
	}
	return nil
}
