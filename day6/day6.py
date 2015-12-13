import re


def turn_on(start_x, start_y, end_x, end_y, _lights):
    for x in range(int(start_x), int(end_x) + 1):
        for y in range(int(start_y), int(end_y) + 1):
            _lights[x][y] += 1


def turn_off(start_x, start_y, end_x, end_y, _lights):
    for x in range(int(start_x), int(end_x) + 1):
        for y in range(int(start_y), int(end_y) + 1):
            _lights[x][y] -= 1
            if _lights[x][y] < 0:
                _lights[x][y] = 0


def toggle(start_x, start_y, end_x, end_y, _lights):
    for x in range(int(start_x), int(end_x) + 1):
        for y in range(int(start_y), int(end_y) + 1):
            _lights[x][y] += 2


strings = open("input", "r").readlines()
lights = [[0 for x in range(1000)] for x in range(1000)]

for command in strings:
    matches = re.search("(\d\d?\d?),(\d\d?\d?) through (\d\d?\d?),(\d\d?\d?)", command)
    if command.startswith("turn on"):
        turn_on(matches.group(1), matches.group(2), matches.group(3), matches.group(4), lights)
    elif command.startswith("turn off"):
        turn_off(matches.group(1), matches.group(2), matches.group(3), matches.group(4), lights)
    elif command.startswith("toggle"):
        toggle(matches.group(1), matches.group(2), matches.group(3), matches.group(4), lights)

lights_lit = 0
for x in range(0, 1000):
    for y in range(0, 1000):
        lights_lit += lights[x][y]
print(lights_lit)