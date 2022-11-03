#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

# This is a collection of bash functions used by different scripts

# imports
. scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem
export PEER0_Company_CA=${PWD}/organizations/peerOrganizations/company.example.com/tlsca/tlsca.company.example.com-cert.pem
export PEER0_Person_CA=${PWD}/organizations/peerOrganizations/person.example.com/tlsca/tlsca.person.example.com-cert.pem
export PEER0_Government_CA=${PWD}/organizations/peerOrganizations/government.example.com/tlsca/tlsca.government.example.com-cert.pem
export PEER0_TargetC_CA=${PWD}/organizations/peerOrganizations/targetc.example.com/tlsca/tlsca.targetc.example.com-cert.pem
export ORDERER_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.crt
export ORDERER_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/server.key

ORGARRAY=("zero" "Company" "Person" "Government" "TargetC")

# Set environment variables for the peer org
setGlobals() {
  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  infoln "Using organization ${USING_ORG}"
  if [ $USING_ORG -eq 1 ]; then
    export CORE_PEER_LOCALMSPID="CompanyMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_Company_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/company.example.com/users/Admin@company.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_LOCALMSPID="PersonMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_Person_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/person.example.com/users/Admin@person.example.com/msp
    export CORE_PEER_ADDRESS=localhost:8051

  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_LOCALMSPID="GovernmentMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_Government_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/government.example.com/users/Admin@government.example.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
  elif [ $USING_ORG -eq 4 ]; then
    export CORE_PEER_LOCALMSPID="TargetCMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_TargetC_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/targetc.example.com/users/Admin@targetc.example.com/msp
    export CORE_PEER_ADDRESS=localhost:10051
  else
    errorln "ORG Unknown"
  fi

  if [ "$VERBOSE" == "true" ]; then
    env | grep CORE
  fi
}

# Set environment variables for use in the CLI container
setGlobalsCLI() {
  setGlobals $1

  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  if [ $USING_ORG -eq 1 ]; then
    export CORE_PEER_ADDRESS=peer0.company.example.com:7051
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_ADDRESS=peer0.person.example.com:8051
  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_ADDRESS=peer0.government.example.com:9051
  elif [ $USING_ORG -eq 4 ]; then
    export CORE_PEER_ADDRESS=peer0.targetc.example.com:10051
  else
    errorln "ORG Unknown"
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    ORG=$1
    PEER="peer0.${ORGARRAY[${ORG}]}"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	    PEERS="$PEER"
    else
	    PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    CA=PEER0_${ORGARRAY[${ORG}]}_CA
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
    # shift by one to get to the next organization
    shift
  done
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}

echoParsePeerConnectionParameters() {
  parsePeerConnectionParameters 1 2 3 4
  echo ${PEER_CONN_PARMS[@]} 
}