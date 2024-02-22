just test for downloaded m3u8 list with golang

directory

```

├── 라디오 캠페인
│   ├── 2024-02-22
├── 라디오 중급 중국어
│   ├── 2024-02-22
│   │   ├── 01.9과 오랜만의 모임
│   │   │   └── 9과 오랜만의 모임.mp3
│   │   └── 02.10과 얼른 튀어와
│   │       └── 10과 얼른 튀어와.mp3
├── 초급 중국어
│   ├── 2024-02-21
│   │   ├── 01.9과 정말 골치 아프네
│   │   │   └── 01.9과 정말 골치 아프네.mp3
├── 입이 트이는 영어
│   ├── 2024-02-22
│   │   ├── 01.A Happy English Student 행복한 영어 학습자
│   │   │   └── A Happy English Student 행복한 영어 학습자.mp3
│   │   ├── 02.A Busan Native Plays in the Snow 부산 사람의 눈놀이
│   │   │   └── A Busan Native Plays in the Snow 부산 사람의 눈놀이.mp3
├── ‘야사시이‘ 초급 일본어
│   ├── 2024-02-21
│   │   ├── 01.어떻게 할지 이미 정했나요? どうするか もう 決めましたか
│   │   │   └── 01.어떻게 할지 이미 정했나요? どうするか もう 決めましたか.mp3
├── ‘타노시이‘ 중급 일본어
│   ├── 2024-02-22
│   │   ├── 01.슈퍼의 계산대, 점원이 앉으면 안 되나요? ス―パーのレジ、店員が座ったらダメですか
│   │   │   └── 슈퍼의 계산대, 점원이 앉으면 안 되나요? ス―パーのレジ、店員が座ったらダメですか.mp3
│   │   ├── 02.당일치기 온천이네요 日帰り温泉ですね
│   │   │   └── 당일치기 온천이네요 日帰り温泉ですね.mp3
│   │   ├── 03.슈퍼의 계산대, 점원이 앉으면 안 되나요? ス―パーのレジ、店員が座ったらダメですか
│   │   │   └── 슈퍼의 계산대, 점원이 앉으면 안 되나요? ス―パーのレジ、店員が座ったらダメですか.mp3
│   │   ├── 04.당일치기 온천이네요 日帰り温泉ですね
│   │   │   └── 당일치기 온천이네요 日帰り温泉ですね.mp3
│   │   ├── media_4738966.ts
│   │   ├── media_4738967.ts
│   │   ├── media_4738968.ts
│   │   ├── media_4738969.ts
│   │   ├── media_4738970.ts
│   │   └── media_4738971.ts

```

log

```
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:43:56.840+0900","caller":"dir/util.go:95","msg":"path","path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:43:56.840+0900","caller":"scrap/run.go:71","msg":"subtitle is exist","current":{"Title":"‘타노시이‘ 중급 일본어","SubTitle":"당일치기 온천이네요 日帰り温泉ですね","StartAt":"2024-02-22T23:40:00+09:00","EndAt":"2024-02-22T23:58:00+09:00","Path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}}
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738992.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738993.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738994.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738995.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738996.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738997.ts
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:44:56.841+0900","caller":"dir/util.go:95","msg":"path","path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:44:56.841+0900","caller":"scrap/run.go:71","msg":"subtitle is exist","current":{"Title":"‘타노시이‘ 중급 일본어","SubTitle":"당일치기 온천이네요 日帰り温泉ですね","StartAt":"2024-02-22T23:40:00+09:00","EndAt":"2024-02-22T23:58:00+09:00","Path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}}
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738998.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4738999.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4739000.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4739001.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4739002.ts
yml-ebs-scrap-1  | downloading... outputs/‘타노시이‘ 중급 일본어/2024-02-22/media_4739003.ts
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:45:56.844+0900","caller":"dir/util.go:95","msg":"path","path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}
yml-ebs-scrap-1  | {"level":"info","timestamp":"2024-02-22T23:45:56.845+0900","caller":"scrap/run.go:71","msg":"subtitle is exist","current":{"Title":"‘타노시이‘ 중급 일본어","SubTitle":"당일치기 온천이네요 日帰り温泉ですね","StartAt":"2024-02-22T23:40:00+09:00","EndAt":"2024-02-22T23:58:00+09:00","Path":"outputs/‘타노시이‘ 중급 일본어/2024-02-22"}}
```

## how to use

`docker build -t {tag} -f cmd/scrap/Dockerfile .`

and modify app.yml

`docker-compose -f app.yml up -d`
