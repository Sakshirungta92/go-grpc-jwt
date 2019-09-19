#!/bin/bash
cd ../pkg/protobuff
./grpcurl -plaintext -d '' -rpc-header "Authorization":"Bearer $AuthToken" -import-path ./ -proto service/server.proto localhost:7777 authservice.service.Auth.GetUserDetails
cd ../../TestScripts