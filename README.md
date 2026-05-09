# Solve-the-puzzle-

This repository (Solve-the-puzzle-) contains the DustNet project: a Cosmos-SDK-based blockchain prototype with staking, a DVS registry, and governance modules. This branch (release/v1.0.0) contains the final deliverable snapshot for v1.0.0.

## What is included in release/v1.0.0
- x/staking/: Scaffolding for the staking module (types, messages, keeper, and unit tests).
- docs/: Documentation and summary README.
- docker-compose.yml and scripts/init-local-testnet.sh for running a local testnet.

## Run local testnet (Docker)

Prerequisites:
- Docker & Docker Compose installed

Start local testnet:

1. Clone the repo and checkout the branch:

   git clone https://github.com/jloshow2024-blip/Solve-the-puzzle-.git
   cd Solve-the-puzzle-
   git checkout release/v1.0.0

2. Start the Docker Compose environment:

   docker-compose up --build

3. Use provided CLI scripts (or the SDK-based node CLI if integrated) to create accounts and send staking transactions.

Note: This branch includes scaffolding and example commands; full CLI integration requires building the node binary.
