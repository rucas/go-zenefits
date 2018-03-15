# Contributing

We love accepting your contributions. Just follow these simple guidelines to get your PR merged ASAP.

## Installing

```bash
$ dep ensure
```

## Tests

In order to run tests you need to set these environment variables:

- `ZENEFITS_API_KEY` 
- `ZENEFITS_PERSON_ID` 
- `ZENEFITS_COMPANY_ID`

```bash
$ go test -v ./zenefits

# Run only one test
$ go test -v ./zenefits -run TestPeopleService_Get
```

## Code Organization

Everthing is under the `zenefits` package, with API methods broken up into
service objects.

## Submitting a Patch

1. Create an [issue]()
2. Fork the repo
3. Run tests (go test, go vet)
4. Push commits to your fork and submit a PR
