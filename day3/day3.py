houses = {}
santa_x, santa_y = 0, 0
robosanta_x, robosanta_y = 0, 0
houses[(santa_x, santa_y)] = 1
string = open("input", "r").read()
ROBO_MODE_ENGAGED = False
for char_num in range(len(string)):
    character = string[char_num]
    if character is "^":
        if ROBO_MODE_ENGAGED:
            robosanta_y += 1
        else:
            santa_y += 1
    elif character is "v":
        if ROBO_MODE_ENGAGED:
            robosanta_y -= 1
        else:
            santa_y -= 1
    elif character is "<":
        if ROBO_MODE_ENGAGED:
            robosanta_x -= 1
        else:
            santa_x -= 1
    elif character is ">":
        if ROBO_MODE_ENGAGED:
            robosanta_x += 1
        else:
            santa_x += 1
    if ROBO_MODE_ENGAGED:
        houses[(robosanta_x, robosanta_y)] = houses.setdefault((robosanta_x, robosanta_y), 0) + 1
    else:
        houses[(santa_x, santa_y)] = houses.setdefault((santa_x, santa_y), 0) + 1
    ROBO_MODE_ENGAGED = not ROBO_MODE_ENGAGED

lucky_children = 0
for house in houses.keys():
    if houses[house] >= 1:
        lucky_children += 1
print(lucky_children)
