#!/bin/bash

function one_line_pem {
    echo "`awk 'NF {sub(/\\n/, ""); printf "%s\\\\\\\n",$0;}' $1`"
}

function json_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.json
}

function yaml_ccp {
    local PP=$(one_line_pem $4)
    local CP=$(one_line_pem $5)
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${P0PORT}/$2/" \
        -e "s/\${CAPORT}/$3/" \
        -e "s#\${PEERPEM}#$PP#" \
        -e "s#\${CAPEM}#$CP#" \
        organizations/ccp-template.yaml | sed -e $'s/\\\\n/\\\n          /g'
}

ORG=1
P0PORT=7051
CAPORT=7054
PEERPEM=organizations/peerOrganizations/company.example.com/tlsca/tlsca.company.example.com-cert.pem
CAPEM=organizations/peerOrganizations/company.example.com/ca/ca.company.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/company.example.com/connection-company.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/company.example.com/connection-company.yaml

ORG=2
P0PORT=8051
CAPORT=8054
PEERPEM=organizations/peerOrganizations/person.example.com/tlsca/tlsca.person.example.com-cert.pem
CAPEM=organizations/peerOrganizations/person.example.com/ca/ca.person.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/person.example.com/connection-person.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/person.example.com/connection-person.yaml

ORG=3
P0PORT=9051
CAPORT=9054
PEERPEM=organizations/peerOrganizations/government.example.com/tlsca/tlsca.government.example.com-cert.pem
CAPEM=organizations/peerOrganizations/government.example.com/ca/ca.government.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/government.example.com/connection-government.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/government.example.com/connection-government.yaml

ORG=4
P0PORT=10051
CAPORT=10054
PEERPEM=organizations/peerOrganizations/targetc.example.com/tlsca/tlsca.targetc.example.com-cert.pem
CAPEM=organizations/peerOrganizations/targetc.example.com/ca/ca.targetc.example.com-cert.pem

echo "$(json_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/targetc.example.com/connection-targetc.json
echo "$(yaml_ccp $ORG $P0PORT $CAPORT $PEERPEM $CAPEM)" > organizations/peerOrganizations/targetc.example.com/connection-targetc.yaml
