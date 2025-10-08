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

### Language/framework specific communities

| Subject    | Male:Female | Kommunity                                                                            | Last $n$ Members | Subject maturity |
| ---------- | ----------- | ------------------------------------------------------------------------------------ | ---------------- | ---------------- |
| .Net       |             | [DotNet Istanbul](https://kommunity.com/dotnet-istanbul)                             | 480              |                  |
| Flutter    |             | [Flutter Turkiye](https://kommunity.com/flutter-turkiye)                             | 460              |                  |
| Go         |             | [Ankara Gophers](https://kommunity.com/ankara-gophers)                               | 460              |                  |
| Go         |             | [GoTurkiye](https://kommunity.com/goturkiye)                                         | 820              |                  |
| Java       |             | [Türkiye Java Community](https://kommunity.com/turkiye-java-community)               | 500              |                  |
| JavaScript |             | [Istanbul JavaScript Topluluğu](https://kommunity.com/istanbul-javascript-toplulugu) | 500              |                  |
| JavaScript |             | [JS İzmir](https://kommunity.com/js-izmir)                                           | 391              |                  |
| PHP        |             | [Istanbul PHP User Group](https://kommunity.com/istanbulphp)                         | 480              |                  |
| React      |             | [React Turkiye](https://kommunity.com/reacttr)                                       | 460              |                  |
| Ruby       |             | [Ruby Turkiye](https://kommunity.com/ruby-turkiye)                                   | 540              |                  |
| Spring     |             | [Spring Türkiye](https://kommunity.com/spring-turkiye)                               | 500              |                  |
| Swift      |             | [Swift Buddies](https://kommunity.com/swiftbuddies)                                  | 560              |                  |
| TensorFlow |             | [TensorFlow Turkey](https://kommunity.com/tensorflow-turkey)                         | 540              |                  |

### Tech focused communities

| Male:Female | Kommunity                                                                                            | Last $n$ Members |
| ----------- | ---------------------------------------------------------------------------------------------------- | ---------------- |
|             | [Teknopark Istanbul](https://kommunity.com/teknopark-istanbul-yazilimci-bulusmalari/members?page=25) | 499              |
|             | [Kadın Yazılımcı](https://kommunity.com/kadinyazilimci)                                              | 500              |
|             | [SistersLab](https://kommunity.com/sisterslaborg)                                                    | 500              |
|             | [Trendyol Tech Meetup](https://kommunity.com/trendyol)                                               | 500              |
|             | [Tech Istanbul](https://kommunity.com/techistanbul)                                                  | 1000             |
|             | [Türkiye Açık Kaynak Platformu](https://kommunity.com/tracikkaynak)                                  | 1000             |
|             | [DevOpsTr](https://kommunity.com/devops-turkiye)                                                     | 1000             |
|             | [DevNot](https://kommunity.com/devnot)                                                               | 999              |
