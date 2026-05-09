#!/bin/bash

# Initialize DustNet local testnet
set -e

if [ -d "data" ]; then
  echo "Data directory already exists. Please remove it if you want to reinitialize."
  exit 1
fi

mkdir -p data

echo "Initializing DustNet node..."
# Placeholder for actual initialization commands.