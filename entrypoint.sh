#!/usr/bin/env sh

echo "Waiting for NATS"
while ! echo exit | nc nats 4222; do sleep 1; done

echo "Starting api-gateway"
/go/bin/api-gateway
