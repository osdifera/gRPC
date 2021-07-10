#!/bin/sh
protoc -I src/ --go_out=go/ src/messages/person.proto