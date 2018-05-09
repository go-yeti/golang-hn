#!/bin/bash

if [ -z "$(which curl)" ]; then
    printf '\e[36m'"We verified that you don't have "'\e[33m'"curl "'\e[36m'" installed";
fi
printf '\e[36m'"We verified that "'\e[33m'"curl "'\e[36m'" is not installed.\nYou can install it by yourself and run this script again or we can install it for you now.\nDo you want to install it now? [Y/n]"'\e[0m'"\n";

read resp

if [ -z "$resp" ]; then
    echo "Empty"
else
    echo "Not empty"
fi


