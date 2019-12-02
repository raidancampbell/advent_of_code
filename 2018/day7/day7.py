import re
import struct

strings = open("input", "r").readlines()
wires = {}

for command in strings:
    if 'NOT' in command:
        result = re.search("NOT (\w+) -> (\w+)", command)
        wires[result.group(2)] = struct.unpack('H', struct.pack('h', ~ wires.get(result.group(1), 0)))[0]
        # result is a NOT of a number or variable
    elif 'AND' in command:
        result = re.search("(\w+) AND (\w+) -> (\w+)", command)
        wires[result.group(3)] = wires.get(result.group(1), 0) & wires.get(result.group(2), 0)
        # result is an AND of two numbers or variables
    elif 'RSHIFT' in command:
        result = re.search("(\w+) RSHIFT (\d+) -> (\w+)", command)
        wires[result.group(3)] = struct.unpack('H', struct.pack('h', wires.get(result.group(1), 0) >> int(result.group(2))))[0]
        # result is a number or variable shifted right by a number or variable
    elif 'LSHIFT' in command:
        result = re.search("(\w+) LSHIFT (\d+) -> (\w+)", command)
        wires[result.group(3)] = struct.unpack('H', struct.pack('h', wires.get(result.group(1), 0) << int(result.group(2))))[0]
        # result is a number or variable shifted left by a number or variable
    elif 'OR' in command:
        result = re.search("(\w+) OR (\w+) -> (\w+)", command)
        wires[result.group(3)] = wires.get(result.group(1), 0) | wires.get(result.group(2), 0)
        # result is an OR of two numbers or variables
    else:  # signal definition
        result = re.search("(\d+|\w+) -> (\w+)", command)
        wires[result.group(2)] = int(result.group(1)) if result.group(1).isnumeric() else wires.get(result.group(1), 0)
for key in wires.keys():
    print(key + ": " + str(wires.get(key)))
print(wires.get('a'))