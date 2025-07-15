## What is gRPC
- gRPC는 **원격 프로시저 호출(Remote Procedure Calls)**을 다루기 위한 강력한 프레임워크다.

## Let's go back a bit — What is RPC?
- gRPC에 들어가기 전에, 먼저 <u>RPC(Remote Procedure Call)**</u>가 무엇인지 이해할 필요가 있다.
- RPC는 일반적인 HTTP 호출 대신 함수 호출을 사용하는 Client-Server 통신 방식이다.
- 호출할 함수와 데이터 타입을 정의하기 위해 **IDL(Interface Definition Language)**을 사용한다.
- gRPC는 완전히 새로운 기술이 아니라, 기존의 RPC 개념을 개선한 버전이다. 
- 이러한 개선 덕분에 gRPC는 매우 빠르게 인기 있는 기술이 되었다.

## When did gRPC become open source?
- 2015년, Google은 자사의 프로젝트를 오픈소스로 공개했으며, 그것이 바로 오늘날의 gRPC다.

## What makes gRPC so popular?
- 추상화가 쉰다. (함수 호출처럼 보인다)
- 다양한 프로그래밍 언어를 지원한다.
- 성능이 뛰어나다.
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


### Types of grpc 

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