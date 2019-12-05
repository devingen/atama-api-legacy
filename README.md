# Atama API

This project contains Atama functions and algorithm. The functions are used in standalone REST API for development 
and AWS Lambda for deployment.

## Endpoints (Functions)

#### build-score-matrix
Builds the score matrix for all pairs. It expects 2 different lists to build the score matrix. If the lists have 
`N` and `M` items, the result matrix will be `N x M`.

#### build-pairs
Generates a pair list by going through different pair combinations that exists. It uses the `build-score-matrix`
function internally to generate the matrix. The goal is to go through all possible combinations. However it takes 
very long time for long lists. So we limit the iteration and deepness to go through some combinations only.

The number of possible combinations can be calculated as `(maxIterationLimit + 1) ^ maxIterationLevel`. If no limit 
is set, `maxIterationLimit` and `maxIterationLevel` re equal to the item count. If we assume that the lists have same
number of items, `N = M`, the possible iteration count table will be like this;

| Items | Combinations  |
|-------|---------------|
| 2     | 9             |
| 3     | 64            |
| 4     | 625           |
| 5     | 7,776         |
| 6     | 117,649       |
| 7     | 2,097,152     |
| 8     | 43,046,721    |
| 9     | 1,000,000,000 |

## Development

Run this command to start development server: `go run api.go`. It will start the server on port `8080`. 

## Testing

**Running tests:** `go test ./...`

**Running tests with coverage report:** `go test ./... -coverprofile=coverage.out`

**Viewing the coverage report:** `go tool cover -html=coverage.out`

## Deployment

1. Generate executable `GOOS=linux go build aws_lambda/build_pairs.go` 
2. Generate bundle `zip build_pairs.zip build_pairs`
3. Upload the bundle to the Lambda with handler name `build_pairs`