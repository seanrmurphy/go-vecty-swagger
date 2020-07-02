#! /usr/bin/env bash

BUCKET=<insert-bucket-name>

echo
echo "Uploading content from this directory to AWS..."
aws s3 sync . s3://$BUCKET

echo
echo "Giving public access to this bocket..."
aws s3api put-public-access-block --bucket $BUCKET --public-access-block-configuration file://json_config/public-access-open.json

echo
echo "Adding bucket policy..."
aws s3api put-bucket-policy --bucket $BUCKET --policy file://json_config/policy.json

echo
echo "Setting up CORS configuration on bucket..."
aws s3api put-bucket-cors --bucket $BUCKET --cors-configuration file://json_config/cors.json

echo
echo "Setting up static website hosting..."
aws s3api put-bucket-website --bucket $BUCKET --website-configuration file://json_config/static-website.json

LOCATION=$(aws s3api get-bucket-location --bucket $BUCKET | jq -r '.LocationConstraint')

echo "Website being served on..."
echo
echo "http://$BUCKET.s3-website.$LOCATION.amazonaws.com"
echo
echo "*** Please disable access to this bucket when you are finished using the script"
echo
echo "./s3_block_access.sh"




