#!/usr/bin/env bash

amp="amp -s localhost"
$amp login --name user --password password
TOKEN=$(cat ~/.config/amp/localhost.credentials)
# TODO: temporarily disabled
return 0
curl -k --header "Authorization: amp $TOKEN" https://gw.local.atomiq.io/v1/stacks | grep "{}"
