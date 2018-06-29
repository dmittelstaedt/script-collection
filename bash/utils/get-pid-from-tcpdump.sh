#!/bin/bash
#
# Get pid from tcpdump stream. Starting tcpdump with -l parameter and pipe
# the output to this script.
#
# Author: David Mittelstaedt <david.mittelstaedt@dataport.de>
# Date: 2018-06-29
read line

host_with_ip=$(cut -d' ' -f3 <<< $line | cut -d. -f5)

pid=$(lsof -Pni:$host_with_ip | grep -v PID | awk '{print $2}' | head -1)
ps $pid

exit 0
