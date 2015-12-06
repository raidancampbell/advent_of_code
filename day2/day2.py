import re
total = 0
ribbon = 0
for present in open("input", "r").readlines():
    match = re.search("(\d+)x(\d+)x(\d+)", present)
    h = int(match.group(1))
    w = int(match.group(2))
    l = int(match.group(3))
    total += 2 * (l * w + w * h + h * l) + min((l * w, w * h, h * l))
    ribbon += h * w * l  # freaking non-euclidian elves
    ribbon += sorted([h, w, l])[0]*2 + sorted([h, w, l])[1]*2
print(total)
print(ribbon)