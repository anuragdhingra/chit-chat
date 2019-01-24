#!/usr/bin/env bash

# Having only production for now, might be used when adding dev/staging envs
case $CIRCLE_BRANCH in
    "master")
        export ENVIRONMENT="production"
        export HEROKU_APP="lets-chatt"
        ;;
    "develop")
        export ENVIRONMENT="production"
        export HEROKU_APP="lets-chatt"
        ;;
esac