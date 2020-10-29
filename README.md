# Metrics

Log your metrics.

## Usage

Setup your logger using the https://github.com/tokenized/logger

Then at the top of functions that you want to measure.

```
defer metrics.Elapsed(ctx, time.Now(), "your-func-name")
```


## License

(c) Tokenized Ltd. 2020 All rights reserved
