# Comparing masculinity-to-maturity ratios of top Turkish Programming Language Kommunities

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
| Flutter    | 3.2 : 1     | [Flutter Turkiye](https://kommunity.com/flutter-turkiye)                             | 459              |
| Java       | 3.2 : 1     | [Türkiye Java Community](https://kommunity.com/turkiye-java-community)               | 499              |
| JavaScript | 3.2 : 1     | [Istanbul JavaScript Topluluğu](https://kommunity.com/istanbul-javascript-toplulugu) | 499              |
| TensorFlow | 3.2 : 1     | [TensorFlow Turkey](https://kommunity.com/tensorflow-turkey)                         | 499              |
| Go         | 4.3 : 1     | [GoTurkiye](https://kommunity.com/goturkiye)                                         | 819              |
| JavaScript | 4.4 : 1     | [JS İzmir](https://kommunity.com/js-izmir)                                           | 390              |
| Swift      | 4.5 : 1     | [Swift Buddies](https://kommunity.com/swiftbuddies)                                  | 559              |
| Go         | 4.7 : 1     | [Ankara Gophers](https://kommunity.com/ankara-gophers)                               | 459              |
| Spring     | 4.8 : 1     | [Spring Türkiye](https://kommunity.com/spring-turkiye)                               | 499              |
| React      | 5.4 : 1     | [React Turkiye](https://kommunity.com/reacttr)                                       | 459              |
| Ruby       | 5.4 : 1     | [Ruby Turkiye](https://kommunity.com/ruby-turkiye)                                   | 539              |
| PHP        | 5.5 : 1     | [Istanbul PHP User Group](https://kommunity.com/istanbulphp)                         | 479              |
| .Net       | 5.7 : 1     | [DotNet Istanbul](https://kommunity.com/dotnet-istanbul)                             | 479              |

### Tech focused communities

This table shows the measurements for non-language specific communities. As this is out-of-scope, the table is shared just for showing the situation in wider landscape.

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

## Comparison

| Subject    | Masculinity | Maturity (yrs) | Masculinity/Maturity |
| ---------- | ----------- | -------------- | -------------------- |
| Java       | 3,2         | 30             | 0,10                 |
| JavaScript | 3,6         | 30             | 0,12                 |
| PHP        | 5,5         | 30             | 0,18                 |
| Ruby       | 5,4         | 30             | 0,18                 |
| Spring     | 4,8         | 23             | 0,20                 |
| .Net       | 5,7         | 23             | 0,24                 |
| Go         | 4,4         | 16             | 0,27                 |
| TensorFlow | 3,2         | 10             | 0,32                 |
| Swift      | 4,5         | 11             | 0,40                 |
| Flutter    | 3,2         | 7              | 0,45                 |
| React      | 5,4         | 12             | 0,45                 |

