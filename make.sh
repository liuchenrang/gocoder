#!/bin/env bash
OUTPUT=$2
CONFIG_FILE=$3
./gen --config $CONFIG_FILE generator --type=xdao --name $1 --o $OUTPUT/daos 
./gen --config $CONFIG_FILE generator --type=idao --name $1 --o $OUTPUT/contract/daos

./gen --config $CONFIG_FILE generator --type=service --name $1 --o $OUTPUT/services
./gen --config $CONFIG_FILE generator --type=iservice --name $1 --o $OUTPUT/contract/services

./gen --config $CONFIG_FILE --config coder.yml generator --type=xorm --name $1 --o $OUTPUT/models
