#!/bin/bash

file=$1
subject=$2

cat $file | age -a -e -r $(dig +short TXT brad.w8rbt.org | tr --delete \") | mutt -i - -s $subject brad@w8rbt.org
