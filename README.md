# ratelimiter
Example Token bucket algorithm in golang

  # How does it work
  - Whenever a request comes it sees if there are any tokens in the bucket
  - If there aren’t any tokens, we throttle the request by returning 429 (too many requests) to the user
  - If the bucket is not empty then the request takes up one of the tokens and proceeds ahead.
  - At every refill rate interval new tokens - tokensPerInterval are added so that final number of tokens in the bucket would be min(current_tokens+tokensPerInterval, bucket_size)
