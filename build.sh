#!/bin/bash
set -e
# typically this could be run through go generate...
docker run -v $(pwd):/data vron/compute computeexample.comp