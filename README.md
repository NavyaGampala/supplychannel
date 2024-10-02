# supplychannel
Creating channel 1:
export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export CHANNEL_NAME=channel1

configtxgen -profile TwoOrgsApplication1 -outputBlock ./channel-artifacts/${CHANNEL_NAME}.block -channelID $CHANNEL_NAME
cp ../config/core.yaml ./configtx/.

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER1_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.crt
export ORDERER1_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:8050 --ca-file "$ORDERER_CA" --client-cert "$ORDERER1_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER1_ADMIN_TLS_PRIVATE_KEY"

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER2_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.crt
export ORDERER2_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer2.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:8051 --ca-file "$ORDERER_CA" --client-cert "$ORDERER2_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER2_ADMIN_TLS_PRIVATE_KEY"

export ORDERER_CA=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
export ORDERER3_ADMIN_TLS_SIGN_CERT=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.crt
export ORDERER3_ADMIN_TLS_PRIVATE_KEY=${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer3.example.com/tls/server.key
osnadmin channel join --channelID $CHANNEL_NAME --config-block ./channel-artifacts/${CHANNEL_NAME}.block -o localhost:8052 --ca-file "$ORDERER_CA" --client-cert "$ORDERER3_ADMIN_TLS_SIGN_CERT" --client-key "$ORDERER3_ADMIN_TLS_PRIVATE_KEY"

source ./scripts/setOrgPeerContext.sh 1
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 2
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 3
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 4
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 5
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block

source ./scripts/setOrgPeerContext.sh 6
peer channel join -b ./channel-artifacts/${CHANNEL_NAME}.block



source ./scripts/setOrgPeerContext.sh 1
docker exec cli ./scripts/setAnchorPeer.sh 1 $CHANNEL_NAME

source ./scripts/setOrgPeerContext.sh 3
docker exec cli ./scripts/setAnchorPeer.sh 3 $CHANNEL_NAME

source ./scripts/setOrgPeerContext.sh 5
docker exec cli ./scripts/setAnchorPeer.sh 5 $CHANNEL_NAME





source ./scripts/setChaincodeContext.sh
export FABRIC_CFG_PATH=$PWD/../config/
export FABRIC_CFG_PATH=${PWD}/configtx
export CHANNEL_NAME=channel1
export PATH=${PWD}/../bin:${PWD}:$PATH

source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode package supply1.tar.gz --path ${CC_SRC_PATH} --lang ${CC_RUNTIME_LANGUAGE} --label supply1_${VERSION}

peer lifecycle chaincode install supply1.tar.gz


source ./scripts/setOrgPeerContext.sh 2
peer lifecycle chaincode install supply1.tar.gz

source ./scripts/setOrgPeerContext.sh 3
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 4
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 5
peer lifecycle chaincode install supply1.tar.gz
source ./scripts/setOrgPeerContext.sh 6
peer lifecycle chaincode install supply1.tar.gz

peer lifecycle chaincode queryinstalled 2>&1 | tee outfile

source ./scripts/setPackageID.sh outfile

source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}


peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --sequence ${VERSION} --output json --init-required
source ./scripts/setOrgPeerContext.sh 3
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}
source ./scripts/setOrgPeerContext.sh 5
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --init-required --package-id ${PACKAGE_ID} --sequence ${VERSION}


source ./scripts/setOrgPeerContext.sh 1
peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name supply1 --version ${VERSION} --sequence ${VERSION} --output json --init-required


source ./scripts/setPeerConnectionParam.sh 1 2 3 4 5 6
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA --channelID $CHANNEL_NAME --name supply1 $PEER_CONN_PARAMS --version ${VERSION} --sequence ${VERSION} --init-required

source ./scripts/setOrgPeerContext.sh 1
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C channel1 -n supply1 $PEER_CONN_PARAMS --isInit -c '{"function":"InitLedger","Args":[]}'
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer1.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer1.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C channel1 -n supply1 $PEER_CONN_PARAMS -c '{"function":"GetAllAssets","Args":[]}'
