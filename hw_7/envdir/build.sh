#!/bin/bash

cd .. && go build main.go && ./main ./env sh -c "echo \$A \$B \$C \$D"