#!/bin/bash
PID="process.pid"
./scv -conf=config.json & echo $! > ${PID}
