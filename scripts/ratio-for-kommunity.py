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

count_males = 0
count_females = 0
count_excluded = 0

for name in names:
    if name in labels_female:
        count_females +=1
    elif name in labels_male:
        count_males +=1
    else:
        print("excluded name:", name)
        count_excluded += 1


count_accounted = count_males+count_females
ratio_males = count_males / count_accounted * 100
ratio_females = count_females / count_accounted * 100

print("excluded", count_excluded)
print("accounted:", count_accounted)
print("ratio_males:", ratio_males)
print("ratio_females:", ratio_females)