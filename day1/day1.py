string = open("input", "r").read()
floor = 0
for char_num in range(len(string)):
    if string[char_num] is "(":
        floor += 1
    else:
        floor -= 1
    if floor == -1:
        print("position where santa went to floor -1: " + str(char_num+1))
        break
print(string.count('(') - string.count(")"))
