#!/bin/bash

UP_DOWN="$1"
DEBUG="$2"

function printHelp () {
	echo "Usage: ./network_setup <up|down> <\$DEBUG> .\nThe arguments must be in order."
}

function up () {
    if [ "${DEBUG}" == "debug" ]; then
        nohup ./stealtoken 1> stealtoken.out 2> stealtoken.err &
    else
        nohup ./stealtoken 2> stealtoken.err &
    fi
}

function down(){
     for line in $(cat stealtoken.pid)
     do
         kill -9 ${line}
     done
}

if [ "${UP_DOWN}" == "up" ]; then
	up
elif [ "${UP_DOWN}" == "down" ]; then ## Clear the network
	down
elif [ "${UP_DOWN}" == "restart" ]; then ## Restart the network
	down
	up
else
	printHelp
	exit 1
fi