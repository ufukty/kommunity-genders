import argparse
parser = argparse.ArgumentParser(description="Example of flag parsing")
parser.add_argument("--input", type=str, required=True, help="Path to the file contains newline separated list of member names")
args = parser.parse_args()

with open(args.input, "r") as file:
    names = [name.strip().lower() for name in file if name]

with open("labels/male.txt", "r") as file:
    labels_male = {name.strip().lower() for name in file if name}

with open("labels/female.txt", "r") as file:
    labels_female = {name.strip().lower() for name in file if name}

import sys

count_males = 0
count_females = 0
count_excluded = 0

for name in names:
    if name in labels_female:
        count_females +=1
    elif name in labels_male:
        count_males +=1
    else:
        print("excluded name:", name,file=sys.stderr)
        count_excluded += 1

count_accounted = count_males+count_females
ratio_females = int(count_females / count_accounted * 100)
ratio_males = 100 - ratio_females

print("excluded", count_excluded, file=sys.stderr)
print("accounted:", count_accounted, file=sys.stderr)
print("ratio_males:", ratio_males, file=sys.stderr)
print("ratio_females:", ratio_females, file=sys.stderr)

if count_females > count_males:
    print("1 : {}".format(int(count_females / count_males * 10) / 10))
else:
    print("{} : 1".format(int(count_males / count_females * 10) / 10))
