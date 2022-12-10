#! /usr/bin/python3

import subprocess

# Build the binary
subprocess.run(["go", "build", "server.go"])

# Build the Docker image
subprocess.run(["docker", "build", "-t", "p2p:latest", "."])

# Start the services defined in the docker-compose.yml file
subprocess.run(["docker-compose", "up"])
