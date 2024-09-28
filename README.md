# 🚀 Rover

A simple yet powerful CLI tool built in Go that helps you deploy Docker images into your VPS via SSH. This tool handles the setup by loading the Docker image into your VPS and running the container automatically. All you need is your VPS credentials, and the **Rover** takes care of the rest.

## ✨ Features

- 🔐 Securely connects to your VPS using SSH.
- 🐳 Supports both Docker images and Dockerfiles.
- 🪄 Automatically saves and transfers Docker images to the remote server.
- 🛠️ Runs Docker containers on your VPS seamlessly.
- 🖥️ Cross-platform support (Linux, macOS, Windows).

## 📦 Installation

### Prerequisites

- **Go**: [Install Go](https://golang.org/doc/install) if you haven't already.
- **Docker**: Ensure Docker is installed on your local machine.

### Install the CLI
#### Method1:
Clone the repository and build the project:

```bash
git clone https://github.com/vansh845/rover.git
cd rover
go build .
```

## 😏 How to use

### Step 1: Initialize your setup
```bash
./rover init
```
#### You will be asked to provide the following information:

`Hostname`: The VPS address (IP or domain).  
`Username`: SSH username.  
`Public Key`: Path to your SSH public key.   
### Step 2: Launch your image
```bash
./rover launch
```
### 👍That's it! An instance of your docker image will be up and running on your VPS.


### Note : Rover is still under development.

