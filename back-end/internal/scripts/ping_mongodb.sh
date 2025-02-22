#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo "Pinging MongoDB..."

# Try to ping MongoDB using mongosh
docker exec mongodb mongosh --eval '
  try {
    db.runCommand("ping");
    print("\n✓ MongoDB is running and responsive!");
  } catch (error) {
    print("\n✗ Failed to ping MongoDB: " + error);
    quit(1);
  }
'

# Check the exit status
if [ $? -eq 0 ]; then
    echo -e "${GREEN}MongoDB connection successful!${NC}"
else
    echo -e "${RED}Failed to connect to MongoDB${NC}"
    echo "Make sure the MongoDB container is running with:"
    echo "docker ps | grep mongodb"
fi