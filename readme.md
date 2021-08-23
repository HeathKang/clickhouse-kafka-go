# Overview
go lib to consume kafka then sink to clickhouse.

# usage
1. create topic :
    ```
    kaf topic create operational_data_batch -p 10 -r 1
    ```
2. run kafka:
   ```
    go run cmd/main.go
   ``` 
