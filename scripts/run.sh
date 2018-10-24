#!/usr/bin/env bash

# Export the necessary environment variables
export $(cat .env | xargs)

# Install the server executable
go install ./cmd/gopx

# Run the server
gopx