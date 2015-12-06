houses = {}
x, y = 0, 0
houses[(x, y)] = 1
string = open("input", "r").read()
for char_num in range(len(string)):
    character = string[char_num]
    if character is "^":
        y += 1
    elif character is "v":
        y -= 1
    elif character is "<":
        x -= 1
    elif character is ">":
        x += 1
    houses[(x, y)] = houses.setdefault((x, y), 0) + 1

lucky_children = 0
for house in houses.keys():
    if houses[house] >= 1:
        lucky_children += 1
print(lucky_children)
