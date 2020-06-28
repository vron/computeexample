#!/bin/bash
set -e
mkdir -p data
cp computeexample.comp data/
docker run -v $(pwd)/data:/data vron/compute
cp data/output/* .
rm -rf data