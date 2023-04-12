#!/bin/bash

# Run 500 GET requests with page=1 parameter
ab -n 50000 -c 2000 "http://localhost:8080/links"