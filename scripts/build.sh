#!/bin/sh

go build ../
sleep 1
mv tsw ../
cd ..
./tsw --help
