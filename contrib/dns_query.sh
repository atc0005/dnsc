#!/bin/bash

# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/dnsc
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

host=$1

declare -a nameservers

nameservers=(

    # Bind servers
    192.168.2.100
    192.168.2.101

    # DCs
    192.168.2.200
    192.168.2.201

)

echo -e "\n\n"

for nameserver in "${nameservers[@]}"
do

    result=$(echo $(dig ${host} @${nameserver} +short) | sed -r "s/\n//g")

    echo -en "\n\n$nameserver: \t${result}"

done

    echo -e "\n"
