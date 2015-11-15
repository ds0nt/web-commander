#!/bin/bash

cd "$(dirname $0)"

cp chat /usr/bin/web-commander
cp chat.service /etc/systemd/system
systemctl enable chat.service
systemctl start chat.service

