# enterprise-go-examples

This project is just going to be used as a base to finish the kata we will do during the workshop.
More info: https://www.eventbrite.com/myevent?eid=105717748632

Checkout `master` for the first kata, but if you get stuck and you can't continue with the second one we have a branch that you can use: `get-product-already-done`.

## Prerequisites

It's a pretty simple project so you don't need much:

- [Go](https://golang.org/dl/) in one of the versions that support "go mod", we recommend >= 1.14
- [Golang CI lint](https://github.com/golangci/golangci-lint) if you are going to use the linter.

You can check that you are OK by just running [`enval`](https://github.com/Adhara-Tech/enval):

    $ enval
    ✔ go:
        ✔ version(>= 1.4): 1.14.2
    ✔ golangci-lint:
        ✔ version(>= 1.26): 1.26.0

## Quick start

    make test # for run all the tests in the project
    make build # to create a binary
    make dev # to run a dev server (you will need gin-gonic/gin)

## More info

After the workshop is finished we will open the repo [`theenglishcut`](https://github.com/landistas/theenglishcut) to all the attendants.
There you can find more gooodies.