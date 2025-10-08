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

## Measurements

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
