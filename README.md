# go_sandbox

sandbox in golang

## line Up

### 01_retry_request

- 外部 API へのリクエスト処理を行う。 (client - provide)
- Provider はレスポンスに 3 秒を要する。Client の Timeout は 1 秒に指定しているので TimeoutError が必ず起きる。
- clinet は 4 回 retry を行い、TimeoutError を返す
