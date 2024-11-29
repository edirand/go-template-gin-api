#!/bin/bash

echo "Running initial setup..."

NEW_BASENAME=$(basename $(pwd))

# Replace the old project name with the new one in all files.
find . -type f | xargs -rn 1 sed -i -e "s/go-template-gin-api/$NEW_BASENAME/g"

# Reset the README, as it currently contains template instructions.
mv docs/DEFAULT_README.md > README.md

echo "Setup complete!"
