def has_pair_of_pairs(input_string=""):
    pairs = []
    last_char = ''
    current_char = ''
    for character in input_string:
        last_char = current_char
        current_char = character
        if not last_char:
            continue  # we're on the first iteration. we don't have a pair yet.


strings = open("input", "r").readlines()
nice_strings = 0
for string in strings:
    vowels = string.count("a") + string.count("e") + string.count("i") + string.count("o") + string.count("u") >= 3
    prev_letter = '?'
    double_letter = False
    for letter_num in range(len(string)):
        if prev_letter == string[letter_num]:
            double_letter = True
            break
        prev_letter = string[letter_num]
    bad_string = string.count("ab") + string.count("cd") + string.count("ab") + string.count("pq") + string.count("xy") > 0
    nice_strings += 1 if vowels and double_letter and not bad_string else 0
print(nice_strings)