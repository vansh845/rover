package internal

import (
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func setConfig() *ssh.ClientConfig {
	fd, err := os.Open("ec2ubuntu.pem")
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
		User: "ubuntu",
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

	client, err := ssh.Dial("tcp", "13.127.48.135:22", cfg)
	if err != nil {
		log.Fatal(err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	return session

}
