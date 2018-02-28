#!/bin/bash
PID="process.pid"
kill -9 `cat ${PID}` && rm ${PID} 
