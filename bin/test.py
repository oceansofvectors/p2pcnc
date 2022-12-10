#! /bin/env/python

import subprocess

# Run the go test command
subprocess.run(["go", "test", "server.go", "server_test.go"])
