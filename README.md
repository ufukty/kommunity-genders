# Latest gender ratios among the top Turkish Kommunities

This projects aims to compare the male-to-female ratio among the both language and tech specific communities in Turkiye.

## Motivation

This project has been started with the intention of investigating the correlation between M:F ratio and language maturity.

## Extracting member names

The script only collects the first word from member names extracted from DOM loaded with the list of latest $n$ member. The $n$ is choosen near 1000 for all communities even for those with 10k members for protecting server resources as the change of ratio with time is out of scope of this project.

```js
Array.from(
  document
    .querySelector(
      "#wrapper > div > section > div.custom-page > div.custom-content > section > div > div > div.user-list"
    )
    .getElementsByClassName("full-name")
)
  .map((e) => {
    const fullname = e.innerText;
    const firstSpace = fullname.indexOf(" ");
    return firstSpace === -1 ? fullname : fullname.substring(0, firstSpace);
  })
  .join("\n");
```

## Name categorization

Lowercase combined list of member names are filtered for unique entries and supplied to an LLM for unisex-excluding classification.

```sh
cat data/* | tr '[:upper:]' '[:lower:]' | sort | uniq
```

## Categorization

LLM generated script provided with name-only list of members and instructed to exclude unisex names. The script generated in the first community is provided as an example below. Actual script for each community utilize different list of names.

<details>
<summary>Expand for example script</summary>

```python
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
```

</details>

## Results

### Language specific communities

| Male:Female | Kommunity                                                              | Last $n$ Members |
| ----------- | ---------------------------------------------------------------------- | ---------------- |
| 4.0 : 1     | [GoTurkiye](https://kommunity.com/goturkiye)                           | 820              |
| 5.2 : 1     | [Türkiye Java Community](https://kommunity.com/turkiye-java-community) | 500              |

### Tech focused communities

| Male:Female | Kommunity                                                           | Last $n$ Members |
| ----------- | ------------------------------------------------------------------- | ---------------- |
| 2.0 : 1     | [Trendyol Tech Meetup](https://kommunity.com/trendyol)              | 500              |
| 2.0 : 1     | [Tech Istanbul](https://kommunity.com/techistanbul)                 | 1000             |
| 4.5 : 1     | [Türkiye Açık Kaynak Platformu](https://kommunity.com/tracikkaynak) | 1000             |
| 4.6 : 1     | [DevOpsTr](https://kommunity.com/devops-turkiye)                    | 1000             |
| 4.7 : 1     | [DevNot](https://kommunity.com/devnot)                              | 999              |
