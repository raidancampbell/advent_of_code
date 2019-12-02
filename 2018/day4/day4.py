secret_key = open("input", "r").read()
import hashlib
m = hashlib.md5()
number = 0
while not str(hashlib.md5((secret_key + str(number)).encode('utf-8')).hexdigest()).startswith("000000"):
    number += 1
print(number)
print(str(m.hexdigest()))
