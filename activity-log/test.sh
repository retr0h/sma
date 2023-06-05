#!/usr/bin/env sh
set -e

echo "=== Insert Test Data ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'

echo "=== Test Descriptions ==="

curl --silent -X GET localhost:8080 -d '{"id": 0}' \
  | jq '.activity.description | contains("christmas eve bike class")'
curl --silent -X GET localhost:8080 -d '{"id": 1}' \
  | jq '.activity.description | contains("cross country skiing")'
curl --silent -X GET localhost:8080 -d '{"id": 2}' \
  | jq '.activity.description | contains("sledding")'
