## 실시간 주문 알림(PUSH) Server


## Description
- 주문 알림과 동시에 비동기로 주문 정보를 사용자에게 알리는 PUSH 서버
- RabbitMQ을 이용해 비동기적 수신 후 처리
- Firebase & FCM 을 이용한 PUSH 발송 
- 단일 PUSH, 멀티 PUSH, Topic PUSH 지원
- IOS, AOS 지원
- Docker 을 이용한 자동 배포




<a href="https://github.com/bluejin1/"><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/bluejin1/mobile-content-trading/tree/main&count_bg=%2379C83D&title_bg=%23142FBC&icon=postwoman.svg&icon_color=%23E7E7E7&title=Back-End&edge_flat=false"/></a>

<img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=Go&logoColor=white"/>
<img src="https://img.shields.io/badge/Firebase-FFCA28?style=flat-square&logo=firebase&logoColor=black"/>
<img src="https://img.shields.io/badge/Rabbitmq-FF6600?style=flat-square&logo=rabbitmq&logoColor=white" />
<img src="https://img.shields.io/badge/Amazon AWS-232F3E?style=flat-square&logo=amazonaws&logoColor=white"/>
<img src="https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=Docker&logoColor=white"/>

<img src="https://img.shields.io/badge/Android-3DDC84?style=flat-square&logo=android&logoColor=white"/>
<img src="https://img.shields.io/badge/iOS-000000?style=flat-square&logo=android&logoColor=white"/>

## Getting started

```bash
# Local 환경 개발(window)

1. go mode init
2. go get

````

```shell
# Build
docker build -t dev/go-fcm-sender:v1 . 
```


## Project Structure
````
root
├── cmd
├── configs
├── helper
├── internal
├── models
└── pkg

````
