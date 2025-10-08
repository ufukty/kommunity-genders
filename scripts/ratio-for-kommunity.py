from collections import Counter
import re

# Load file
with open("/mnt/data/names.txt", "r", encoding="utf-8") as f:
    names = [re.sub(r"[^a-zA-ZğüşıöçĞÜŞİÖÇ]+", "", n.strip().lower()) for n in f if n.strip()]

# Remove empty strings
names = [n for n in names if n]

# Define male and female name sets (based on Turkish/common patterns)
male_names = {
    "ahmet", "mehmet", "ali", "mustafa", "murat", "burak", "emre", "enes", "furkan", "ömer", "hasan",
    "hüseyin", "ibrahim", "yusuf", "kadir", "ramazan", "selim", "halil", "fatih", "barış", "berkay",
    "mert", "kerem", "kaan", "alper", "gökhan", "onur", "sinan", "cengiz", "batuhan", "yunus", "recep",
    "emir", "omer", "yasin", "taha", "tuncay", "samed", "samet", "ismail", "abdullah", "abdul", "adem",
    "enes", "hakan", "omer", "ugur", "ahmetcan", "mehmetali", "orhan", "ozan", "tamer", "tolga", "ozgur"
}

female_names = {
    "ayşe", "fatma", "zeynep", "elif", "büşra", "merve", "melike", "ayşegül", "esra", "hilal", "sena",
    "melis", "selin", "kübra", "beyza", "meltem", "yasemin", "özge", "melike", "banu", "duygu", "gül",
    "ece", "sevda", "sümeyye", "seher", "rabia", "hümeyra", "hazal", "ayşenur", "nisa", "ayşe", "yaren",
    "gaye", "leyla", "sema", "seda", "sevil", "tuğçe", "sinem", "özlem", "ayça", "aybüke", "beyzanur"
}

unisex_names = {
    "deniz", "derya", "doğan", "doğa", "evren", "ilhan", "olcay", "umut", "sevgi", "songül", "dilara"
}

# Count occurrences
counts = Counter(names)

# Categorize
male_count = sum(count for name, count in counts.items() if name in male_names)
female_count = sum(count for name, count in counts.items() if name in female_names)
excluded_count = sum(count for name, count in counts.items() if name in unisex_names)

# Calculate totals and ratio
total = male_count + female_count
male_ratio = male_count / total * 100 if total else 0
female_ratio = female_count / total * 100 if total else 0

male_count, female_count, excluded_count, total, male_ratio, female_ratio