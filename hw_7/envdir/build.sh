#!/bin/bash

go build ../main.go && ./main ../env sh -c "echo $A"