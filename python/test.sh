#!/usr/bin/env bash

pip show protobuf
rm -f example/*.pyc
rm -f example/*pb2.py
rm -f example/*grpc.py
python -m grpc_tools.protoc -I./example --python_out=./example --grpc_python_out=./example ./example/*.proto
 
python app.py