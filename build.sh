#!/bin/bash
set -e
docker run -v $(pwd):/data vron/compute shaders/computeexample.comp
