#!/bin/bash
#
# Create directories from config file
#
# Author: David Mittelstaedt <david.mittelstaedt@dataport.de>
# Date: 2018-06-28

CONFIG_FILE=sample-structure.txt
BASE_DIR=.

while read -r line; do
  mkdir -p $BASE_DIR/$line
done < $CONFIG_FILE

exit 0
