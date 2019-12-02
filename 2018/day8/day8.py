strings = open("input", "r").readlines()
characters = 0
memory_characters = 0

for string in strings:
    string = string.replace("\n", '')
    memory_characters += len(string)
    characters += len(string) - 2  # beginning and ending quotes
    characters -= 3 * string.count("\\x")  # each hex encoding takes 4 memory characters to make 1 character
    characters += 3 * string.count("\\\\x")
    # this is a problem.  Because we can have an escaped slash, followed by an `x` which is not a hex encoding
    characters -= (string.count("\\\"") + string.count("\\\\"))  # another character for the slash in a delimited quote.
    print(memory_characters)
    print(characters)

print(memory_characters)
print(characters)
print(memory_characters - characters)
