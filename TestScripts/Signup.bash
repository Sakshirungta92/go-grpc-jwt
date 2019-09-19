#!/bin/bash:
input='{"name":"'$1'", "email":"'$2'","password":"'$3'"}'
cd ../pkg/protobuff/ && tk=$( ./grpcurl -plaintext -d "$input" -import-path ./ -proto service/server.proto localhost:7777 authservice.service.Auth.SignUp | grep -oP '(?<="token": ")[^"]*' ) 
cd ../../TestScripts
export AuthToken=$tk
