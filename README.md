# ccwc-go
ccwc is a clone of the popular `wc` unix command line utility written in Go using the Go stdlib. 

The core features are:
- returns the bytes, lines, words, or characters in a file or stdin 
- Accepts input from `stdin` via the pipe operator (`|`) or a file path
- provides flexible output based on simple flags
- default return (no flags) is: `lines words bytes`
- if no filename is specified the input defaults to stdin
- *FYI: go must be installed on the host to use this utility

## 🚧 Next Steps
- Some code runs when it is not needed, which is inefficient. For example, we count lines when the user only requests the character count to be returned. This should be corrected with conditional logic.
- We could improve readability and maintainability by moving the core functionality outside the main(). Breaking the logic into a small number of functions would also make it easy to implement tests for these functions instead of simply testing the main(). 

# Installation
## Step 0: Clone the repo
```sh
git clone https://github.com/ogorman89/ccwc-go.git
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

# Using `ccwc`
You can now use `ccwc` via the command line!

## run on a specific file
```sh
ccwc test.txt
```
\>\>\> 7145 58164 342190 test.txt

## run on stdin via pipe operator
```sh
echo "test one two three" | ccwc
```
\>\>\> 1 4 19

or

```sh
cat test.txt | ccwc
```
\>\>\> 7145 58164 342190 

## Optional flags
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
\>\>\> 7145

```sh
ccwc -c test.txt
```
\>\>\> 342190 test.txt
