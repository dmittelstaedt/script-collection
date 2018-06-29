#!/bin/bash
#
# Get the pid of a port given as first argument to this script
#
# Author: David Mittelstaedt <david.mittelstaedt@dataport.de>
# Date: 2018-06-29

pid=$(lsof -Pni:$1 | grep -v PID | awk '{print $2}' | head -1)
ps $pid

exit 0
