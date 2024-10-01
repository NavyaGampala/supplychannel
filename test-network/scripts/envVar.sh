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
export PEER0_ORG1_CA=${PWD}/organizations/peerOrganizations/producer.example.com/tlsca/tlsca.producer.example.com-cert.pem
export PEER1_ORG1_CA=${PWD}/organizations/peerOrganizations/producer.example.com/tlsca/tlsca.producer.example.com-cert.pem
export PEER0_ORG2_CA=${PWD}/organizations/peerOrganizations/supplier.example.com/tlsca/tlsca.supplier.example.com-cert.pem
export PEER1_ORG2_CA=${PWD}/organizations/peerOrganizations/supplier.example.com/tlsca/tlsca.supplier.example.com-cert.pem
export PEER0_ORG3_CA=${PWD}/organizations/peerOrganizations/wholeseller.example.com/tlsca/tlsca.wholeseller.example.com-cert.pem
export PEER1_ORG3_CA=${PWD}/organizations/peerOrganizations/wholeseller.example.com/tlsca/tlsca.wholeseller.example.com-cert.pem
export ORDERER1_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.crt
export ORDERER1_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.key
export ORDERER2_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.crt
export ORDERER2_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.key
export ORDERER3_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.crt
export ORDERER3_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.key

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
    export CORE_PEER_LOCALMSPID="ProducerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/producer.example.com/users/Admin@producer.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7054
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_LOCALMSPID="ProducerMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG1_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/producer.example.com/users/Admin@producer.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7055

  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_LOCALMSPID="SupplierMSP"
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG2_CA
    export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/supplier.example.com/users/Admin@supplier.example.com/msp
    export CORE_PEER_ADDRESS=localhost:7056
  elif [ $USING_ORG -eq 4 ]; then
   export CORE_PEER_LOCALMSPID="SupplierMSP"
   export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG2_CA
   export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/supplier.example.com/users/Admin@supplier.example.com/msp
   export CORE_PEER_ADDRESS=localhost:7057
  elif [ $USING_ORG -eq 5 ]; then
      export CORE_PEER_LOCALMSPID="WholesellerMSP"
      export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_ORG3_CA
      export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/wholeseller.example.com/users/Admin@wholeseller.example.com/msp
      export CORE_PEER_ADDRESS=localhost:7058
   elif [ $USING_ORG -eq 6 ]; then
     export CORE_PEER_LOCALMSPID="WholesellerMSP"
     export CORE_PEER_TLS_ROOTCERT_FILE=$PEER1_ORG3_CA
     export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/wholeseller.example.com/users/Admin@wholeseller.example.com/msp
     export CORE_PEER_ADDRESS=localhost:7059

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
    export CORE_PEER_ADDRESS=peer0.producer.example.com:7054
  elif [ $USING_ORG -eq 2 ]; then
    export CORE_PEER_ADDRESS=peer1.producer.example.com:7055
  elif [ $USING_ORG -eq 3 ]; then
    export CORE_PEER_ADDRESS=peer0.supplier.example.com:7056
  elif [ $USING_ORG -eq 4 ]; then
    export CORE_PEER_ADDRESS=peer1.supplier.example.com:7057
  elif [ $USING_ORG -eq 5 ]; then
    export CORE_PEER_ADDRESS=peer0.wholeseller.example.com:7058
  elif [ $USING_ORG -eq 6 ]; then
    export CORE_PEER_ADDRESS=peer1.wholeseller.example.com:7059
  else
    errorln "ORG Unknown"
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
#parsePeerConnectionParameters() {
#  PEER_CONN_PARMS=()
#  PEERS=""
#  while [ "$#" -gt 0 ]; do
#    setGlobals $1
#    ## Set peer addresses
#    if [ -z "$PEERS" ]
#    then
#	PEERS="$PEER"
#    else
#	PEERS="$PEERS $PEER"
#    fi
#    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
#    ## Set path to TLS certificate
#    CA=PEER0_ORG$1_CA
#    TLSINFO=(--tlsRootCertFiles "${!CA}")
#    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
#    # shift by one to get to the next organization
#    shift
#  done
#}
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    local USING_ORG=""
    if [ -z "$OVERRIDE_ORG" ]; then
      USING_ORG=$1
    else
      USING_ORG="${OVERRIDE_ORG}"
    fi

    if [ $USING_ORG -eq 1 ]; then
      PEER="peer0.producer"
      CA=PEER0_ORG1_CA

    elif [ $USING_ORG -eq 2 ]; then
      PEER="peer1.producer"
      CA=PEER1_ORG1_CA

    elif [ $USING_ORG -eq 3 ]; then
      PEER="peer0.supplier"
      CA=PEER0_ORG2_CA

    elif [ $USING_ORG -eq 4 ]; then
      PEER="peer1.supplier"
      CA=PEER1_ORG2_CA

    elif [ $USING_ORG -eq 5 ]; then
      PEER="peer0.wholeseller"
      CA=PEER0_ORG3_CA

    elif [ $USING_ORG -eq 6 ]; then
      PEER="peer1.wholeseller"
      CA=PEER1_ORG3_CA
    else
      errorln "ORG Unknown"
    fi
    # PEER="peer0.org$1"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	    PEERS="$PEER"
    else
	    PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    # CA=PEER0_ORG$1_CA
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
