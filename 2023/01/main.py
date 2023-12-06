f = open("input.txt", "r")
input = []
for x in f:
    input.append(x)

input_test = [
    "1abc2",
"pqr3stu8vwx",
"a1b2c3d4e5f",
"treb7uchet"
]

input_test_2 = [
    "two1nine",
"eightwothree",
"abcone2threexyz",
"xtwone3four",
"4nineeightseven2",
"zoneight234",
"7pqrstsixteen"
]

def replace_all(replace_line):
    replace_line = replace_line.replace("one", "1")
    replace_line = replace_line.replace("two", "2")
    replace_line = replace_line.replace("three", "3")
    replace_line = replace_line.replace("four", "4")
    replace_line = replace_line.replace("five", "5")
    replace_line = replace_line.replace("six", "6")
    replace_line = replace_line.replace("seven", "7")
    replace_line = replace_line.replace("eight", "8")
    replace_line = replace_line.replace("nine", "9")
    return replace_line


all_code = []
for code in input_test_2:
    first = ""
    last = ""
    code = replace_all(code)
    print(code)
    for char in code:
        if (char.isdigit()):
            if (first ==""):
                first = char
            else:
                last = char
    if (last == ""):
        last = first
    all_code.append(first+last)

print(all_code)
product = 0

for number in all_code:
    product = product + int(number)

print(product)