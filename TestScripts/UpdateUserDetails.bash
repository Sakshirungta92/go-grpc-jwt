#!/bin/bash
input='{"name":"'$1'", password":"'$2'"}'
cd ../pkg/protobuff
./grpcurl -plaintext -d "$input" -rpc-header "Authorization":"Bearer $AuthToken" -import-path ./ -proto service/server.proto localhost:7777 authservice.service.Auth.UpdateUserDetails
cd ../../TestScripts