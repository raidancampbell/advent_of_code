import re
total = 0
for present in open("input", "r").readlines():
    match = re.search("(\d+)x(\d+)x(\d+)", present)
    h = int(match.group(1))
    w = int(match.group(2))
    l = int(match.group(3))
    total += 2 * (l * w + w * h + h * l) + min((l * w, w * h, h * l))
print(total)