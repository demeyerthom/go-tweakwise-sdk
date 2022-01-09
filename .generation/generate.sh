#!/bin/bash

projectName="backend"

openapi-generator-cli generate \
  -i "https://navigator-api.tweakwise.com/swagger/docs/v1" \
  -g go \
  -o "$projectName" \
  --skip-validate-spec \
  -c .generation/config.yml \
  --package-name "$projectName"