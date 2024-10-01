#!/bin/bash

PACK_ID=$(sed -n "/supply1_${VERSION}/{s/^Package ID: //; s/, Label:.*$//; p;}" $1)
export PACKAGE_ID=$PACK_ID
echo $PACKAGE_ID