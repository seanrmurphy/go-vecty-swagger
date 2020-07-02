#! /usr/bin/env bash

BUCKET=<insert-bucket-name>

aws s3api put-public-access-block --bucket $BUCKET --public-access-block-configuration file://json_config/public-access-blocked.json

#aws s3api put-bucket-policy --bucket $BUCKET --policy file://policy.json
#aws s3api put-bucket-cors --bucket $BUCKET --cors-configuration file://cors.json
