#! /bin/bash
aws ssm get-parameter --with-decryption --name /souflair/env | jq -r .Parameter.Value | echo -e "$(cat -)" > /var/app/.env
