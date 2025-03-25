# ratelimiter
Example Token bucket algorithm in golang

# How does it work?
- Whenever a request comes it sees if there are any tokens in the bucket
- If there aren’t any tokens, we throttle the request by returning 429 (too many requests) to the user
- If the bucket is not empty then the request takes up one of the tokens and proceeds ahead.
- At every refill rate interval new tokens - tokensPerInterval are added so that final number of tokens in the bucket would be min(current_tokens+tokensPerInterval, bucket_size)

# Testing with 2 requests per second, no bursting 
```
for i in {1..6}; do curl -i http://localhost:8080/ping; sleep 2; done 
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:31 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:33 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:35 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:37 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:39 GMT
Content-Length: 81

{"status":"Successful","body":"Hi! You've reached the API. How may I help you?"}
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:05:41 GMT
Content-Length: 81
```

# Testing with more than 2 requests per second, last 2 requests are throttled
```
for i in {1..6}; do curl -I http://localhost:8080/ping; done         
HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 81

HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 81

HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 81

HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 81

HTTP/1.1 429 Too Many Requests
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 78
Content-Type: text/plain; charset=utf-8

HTTP/1.1 429 Too Many Requests
Date: Tue, 25 Mar 2025 12:19:20 GMT
Content-Length: 78
Content-Type: text/plain; charset=utf-8
```



