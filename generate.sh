#!/bin/sh
protoc -I src/ --go_out=. src/messages/person.proto