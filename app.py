from __future__ import print_function

import os

print('--------------------------------')
print()
print()
print()
print("run 'htop --pid {}', then hit enter to trigger the protobuf import (watch RES)".format(os.getpid()))
raw_input()
from example.example_pb2 import *
print()
print()
print()
print("from example.example_pb2 import *")
print()
print()
print()
print('hit enter to exit')
raw_input()
