# Go course (beginner)

## To run the slides in your browser

The course material consists of slides containing executable examples.
You can run them in your browser after following a few simple steps.

1. Make sure Go is installed on your machine:

    ```
    go version
    ```

2. Install the `present` tool:

    ```
    go get golang.org/x/tools/cmd/present
    ```

3. Clone this repository and `cd` to the clone:
 
    ```
    git clone https://github.com/jub0bs/go-course-beginner && cd go-course-beginner
    ```

4. Run the present tool:

    ```
    present
    ```

### Troubleshooting

At step 4, you may get the following error message:

```
Couldn't find gopresent files: no required module provides package golang.org/x/tools/cmd/present: go.mod file not found in current directory or any parent directory; see 'go help modules'

By default, gopresent locates the slide template files and associated
static content by looking for a "golang.org/x/tools/cmd/present" package
in your Go workspaces (GOPATH).

You may use the -base flag to specify an alternate location.
```

This problem seems caused by a [regression](https://github.com/RedHatOfficial/GoCourse/issues/108).
As a workaround, run the following command instead:

```
present -base "$(go env GOMODCACHE)"/golang.org/x/tools@v0.1.2/cmd/present
```

You may have to adjust the version (the `v0.1.2` part), depending on which version you have installed.
