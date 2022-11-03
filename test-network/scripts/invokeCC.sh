#!/bin/bash

source scripts/utils.sh

CHANNEL_NAME=${1:-"mychannel"}
CC_NAME=${2}
CC_INOVKE_FCN=${3}
DELAY=${4:-"3"}
MAX_RETRY=${5:-"1"}

println "executing with the following"
println "- CHANNEL_NAME: ${C_GREEN}${CHANNEL_NAME}${C_RESET}"
println "- CC_NAME: ${C_GREEN}${CC_NAME}${C_RESET}"
println "- DELAY: ${C_GREEN}${DELAY}${C_RESET}"
println "- MAX_RETRY: ${C_GREEN}${MAX_RETRY}${C_RESET}"

FABRIC_CFG_PATH=$PWD/tool/config/


if [ -z "$CC_NAME" ] || [ "$CC_NAME" = "NA" ]; then
  fatalln "No chaincode name was provided. Valid call example: ./network.sh invokeCC -ccn basic -ccf '{\"function\":\"TransferAsset\",\"Args\":[\"asset6\",\"Christopher\"]}'"

elif [ -z "$CC_INOVKE_FCN" ] || [ "$CC_INOVKE_FCN" = "NA" ]; then
  fatalln "No chaincode function was provided. Valid call example: ./network.sh invokeCC -ccn basic -ccf '{\"function\":\"TransferAsset\",\"Args\":[\"asset6\",\"Christopher\"]}'"
fi

# import utils
. scripts/envVar.sh
. scripts/ccutils.sh

infoln "Invoke chaincode by ${CC_INOVKE_FCN}"
chaincodeInvoke ${CC_INOVKE_FCN} 
