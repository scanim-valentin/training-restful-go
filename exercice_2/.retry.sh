#!/bin/bash
docker compose down
cd src && docker build --tag exercice_2-api . && cd ..
# cd front/app && docker build --tag exercice_2-front . && cd ../..
docker compose up --force-recreate 
