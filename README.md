# ccwc-go
ccwc is a command line utility written in Go using the Go stdlib that:
- returns the bytes, lines, words, or characters in a file or stdin 
- Accepts input from `stdin` via the pipe operator (`|`) or a file path
- provides flexible output based on simple flags
- *go must be installed on the host to use this utility
- DEFAULT return is: lines words bytes

# Installation
## Step 0: Clone the repo
```sh
`git clone https://github.com/ogorman89/ccwc-go.git`
cd ccwc-go
```

## Step 1: Build the binary for your environment
```sh
go build -o ccwc main.go`
```

## Step 2: Move the binary into your GOPATH (for global installation)
```sh
sudo mv ccwc /Users/usr/go/bin`
```
*your bin may be in a different location, if you need to find your gopath type `$GOPATH` in your terminal

## Using `ccwc`
You can now use `ccwc` via the command line!

### run on a specific file
```sh
ccwc test.txt
```

### run on stdin via pipe operator
```sh
echo "hello world!" | ccwc
```
or
```sh
cat test.txt | ccwc
```

#### Optional flags
| Flag | Description |
|----|-------------|
| -c |	# bytes |
| -l |	# lines |
| -w |	# words |
| -m |	# characters |
| no | defaults to l w c |

Examples:

```sh
cat test.txt | ccwc -l
```

```sh
ccwc -c test.txt
```

