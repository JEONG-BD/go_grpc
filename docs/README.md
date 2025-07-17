## What is gRPC
- gRPC는 **원격 프로시저 호출(Remote Procedure Calls)**을 다루기 위한 강력한 프레임워크다.

## The Motivation of gRPC 
### 서로 다른 언어 간 통신이 필요하다.
- 백엔드와 프론트엔드는 서로 다른 언어로 작성될 수 있다.
- 마이크로서비스들도 각기 다른 언어로 개발될 수 있다.
- 이들 간에 정보를 주고받기 위해 API 계약(API contract)에 대한 합의가 필요하다.
    - 통신 방식: REST, SOAP, Message Queue 
    - 인증 방식: Basic, OAuth, JWT
    - 데이터 포맷: JSON, XML, Binary
    - 데이터 모델
    - 에러 처리 방식

### 통신은 효율적이어야 한다
- 마이크로서비스 간에는 **엄청난 양의 메시지**가 교환된다.  
- **모바일 네트워크는 느리고 대역폭이 제한**될 수 있다.

### 통신은 단순해야 한다
- 클라이언트와 서버는 **핵심 서비스 로직에 집중 해야 하고 나머지는 프레임워크에 맡기는 것이 좋다.**

## Let's go back a bit — What is RPC?
- gRPC에 들어가기 전에, 먼저 <u>**RPC(Remote Procedure Call)**</u>가 무엇인지 이해할 필요가 있다.
- RPC는 일반적인 HTTP 호출 대신 함수 호출을 사용하는 Client-Server 통신 방식이다.
- 호출할 함수와 데이터 타입을 정의하기 위해 IDL(Interface Definition Language)을 사용한다.
- gRPC는 완전히 새로운 기술이 아니라, 기존의 RPC 개념을 개선한 버전이며 덕분에 gRPC는 매우 빠르게 인기 있는 기술이 되었다.
- 개발자가 원격 상호작용의 세부 사항을 명시적으로 코딩하지 않아도 된다.  
- 클라이언트 코드에서는 마치 서버 코드의 함수를 직접 호출하는 것처럼 보인다.  
- 클라이언트와 서버 코드는 서로 다른 언어로 작성될 수 있다.   

## When did gRPC become open source?
- 2015년, Google은 자사의 프로젝트를 오픈소스로 공개했으며, 그것이 바로 오늘날의 gRPC다.

## What makes gRPC so popular?
- 추상화가 쉽다. (함수 호출처럼 보인다)
- 다양한 프로그래밍 언어를 지원하고 성능이 뛰어나다.
- HTTP 호출은 복잡한 반면, gRPC는 훨씬 간결하고 직관적이다.
- gRPC가 인기 있는 핵심 이유는 마이크로서비스 아키텍처가 대중화되었기 때문이다.
- 마이크로서비스 환경에서는 Golang, Java 등 다양한 언어로 구성된 서비스 간 통신이 필요하다.

## Architecture of gRPC
### Request / Response Multiplexing
- 하나의 연결에서 여러 요청과 응답을 동시에 처리할 수 있다.

### What else does gRPC offer 
- Metadata (헤더 정보)
- Streaming (양방향 또는 단방향 스트리밍 지원)

### Load Balancing?

## How gRPC works ? 
- 클라이언트는 서버와 동일한 메서드를 제공하는 자동 생성된 스텁(stub)을 가진다.

## How stubs are generated? 
- API 계약서(API contract) 설명서를 기반으로 서비스와 메시지(payload)는 Protocol Buffer를 사용해 정의된다.  
- 서버와 클라이언트 스텁은 Protocol Buffer 컴파일러와 각 언어별 gRPC 플러그인에 의해 생성된다  

## Why gRPC use Protocol Buffer 
- 사람이 읽을 수 있는 인터페이스 정의 언어(Interface Definition Language, IDL)  
- 다양한 프로그래밍 언어 간 상호 운용성: 여러 언어용 코드 생성기 제공한다. 
- 바이너리 데이터 표현:  
  - 크기가 작다.
  - 전송이 빠르다.  
  - 직렬화/역직렬화가 효율적이다.  
- 강력한 타입 계약 제공  
- API 진화를 위한 규칙을 제공한다.   
- 이전 버전과 이후 버전 간 호환성 지원(Backward & forward compatibility)  


## Types of grpc 
#### Unary operation
→ 단일 요청-응답 방식 (클라이언트가 한 번 요청하고 서버가 한 번 응답)

#### Server streaming
→ 서버 스트리밍 (클라이언트가 한 번 요청하고, 서버가 여러 번 응답을 보냄)

#### Client streaming
→ 클라이언트 스트리밍 (클라이언트가 여러 번 데이터를 보내고, 서버는 한 번 응답)

#### Bidirectional streaming
→ 양방향 스트리밍 (클라이언트와 서버가 동시에 여러 번 데이터를 주고받음)

### protoc 
```shell
protoc --go-grpc_out=. --go_out=. *.proto
```