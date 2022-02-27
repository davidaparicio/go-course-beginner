# Go course (beginner)

## To run the slides in your browser

The course material consists of slides containing executable examples.
You can run them in your browser after following a few simple steps.

1. Make sure Go is installed on your machine:

    ```shell
    go version
    ```

2. Install the `present` tool:

    ```shell
    go install golang.org/x/tools/cmd/present
    ```

3. Clone this repository and `cd` to the clone:

    ```shell
    git clone https://github.com/jub0bs/go-course-beginner && cd go-course-beginner
    ```

4. Run the present tool:

    ```shell
    present
    ```

## Troubleshooting

If the `present` binary cannot be found, run the following command:

```
go env GOBIN
```

Make sure the path in the output is on your `PATH`.
If the output is empty, run the following command:

```shell
echo "$(go env GOPATH)/bin"
```

Make sure the path in the output is on your `PATH`.


