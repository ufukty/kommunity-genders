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

Lowercase combined list of member names are filtered for unique [entries](labels/uniq-names.txt) and supplied to an LLM for unisex-excluding classification for [male](labels/male-names.txt) and [female](labels/female-names.txt) names.

```sh
cat data/* | tr '[:upper:]' '[:lower:]' | sort | uniq > labels/uniq-names.txt
```

## Misc.

Scrip to run ratio calculation script for each community member list:

```sh
for data in data/*; do
  echo -n "$data;";
  python3 scripts/ratio-for-kommunity.py --input "$data" 2>/dev/null;
done > stats.txt
```

## Measurements

### Language/framework specific communities

| Subject    | Male:Female | Kommunity                                                                            | Last $n$ Members |
| ---------- | ----------- | ------------------------------------------------------------------------------------ | ---------------- |
| .Net       | 5.7 : 1     | [DotNet Istanbul](https://kommunity.com/dotnet-istanbul)                             | 479              |
| Flutter    | 3.2 : 1     | [Flutter Turkiye](https://kommunity.com/flutter-turkiye)                             | 459              |
| Go         | 4.7 : 1     | [Ankara Gophers](https://kommunity.com/ankara-gophers)                               | 459              |
| Go         | 4.3 : 1     | [GoTurkiye](https://kommunity.com/goturkiye)                                         | 819              |
| Java       | 3.2 : 1     | [Türkiye Java Community](https://kommunity.com/turkiye-java-community)               | 499              |
| JavaScript | 3.2 : 1     | [Istanbul JavaScript Topluluğu](https://kommunity.com/istanbul-javascript-toplulugu) | 499              |
| JavaScript | 4.4 : 1     | [JS İzmir](https://kommunity.com/js-izmir)                                           | 390              |
| PHP        | 5.5 : 1     | [Istanbul PHP User Group](https://kommunity.com/istanbulphp)                         | 479              |
| React      | 5.4 : 1     | [React Turkiye](https://kommunity.com/reacttr)                                       | 459              |
| Ruby       | 5.4 : 1     | [Ruby Turkiye](https://kommunity.com/ruby-turkiye)                                   | 539              |
| Spring     | 4.8 : 1     | [Spring Türkiye](https://kommunity.com/spring-turkiye)                               | 499              |
| Swift      | 4.5 : 1     | [Swift Buddies](https://kommunity.com/swiftbuddies)                                  | 559              |
| TensorFlow | 3.2 : 1     | [TensorFlow Turkey](https://kommunity.com/tensorflow-turkey)                         | 499              |

### Tech focused communities

| Male:Female | Kommunity                                                                            | Last $n$ Members |
| ----------- | ------------------------------------------------------------------------------------ | ---------------- |
| 1 : 0.4     | [SistersLab](https://kommunity.com/sisterslaborg)                                    | 499              |
| 1 : 0.8     | [Kadın Yazılımcı](https://kommunity.com/kadinyazilimci)                              | 499              |
| 1.3 : 1     | [Tech Istanbul](https://kommunity.com/techistanbul)                                  | 999              |
| 1.5 : 1     | [Trendyol Tech Meetup](https://kommunity.com/trendyol)                               | 499              |
| 3.1 : 1     | [Türkiye Açık Kaynak Platformu](https://kommunity.com/tracikkaynak)                  | 999              |
| 3.3 : 1     | [Teknopark Istanbul](https://kommunity.com/teknopark-istanbul-yazilimci-bulusmalari) | 498              |
| 3.4 : 1     | [DevNot](https://kommunity.com/devnot)                                               | 998              |
| 3.4 : 1     | [DevOpsTr](https://kommunity.com/devops-turkiye)                                     | 999              |

## Calculations

| Male:Female | Subject    | Maturity (Initial release)              |
| ----------- | ---------- | --------------------------------------- |
| 3.2 : 1     | Flutter    | ~7 years (2018 (stable))                |
| 3.2 : 1     | Java       | ~30 years (1995)                        |
| 3.2 : 1     | TensorFlow | ~10 years (2015 (public release))       |
| 3.6 : 1     | JavaScript | ~30 years (1995)                        |
| 4.4 : 1     | Go         | ~16 years (2009)                        |
| 4.5 : 1     | Swift      | ~11 years (2014)                        |
| 4.8 : 1     | Spring     | ~23 years (2002 (Spring Framework 1.0)) |
| 5.4 : 1     | React      | ~12 years (2013)                        |
| 5.4 : 1     | Ruby       | ~30 years (1995)                        |
| 5.5 : 1     | PHP        | ~30 years (1995)                        |
| 5.7 : 1     | .Net       | ~23 years (2002)                        |
