#!/bin/bash

ufw allow 53/tcp
ufw allow 53/udp
ufw allow 443/tcp
ufw allow 8080/tcp

ufw --force enable
