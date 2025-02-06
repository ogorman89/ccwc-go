# ccwc-go
wc like command line tool in go

Go must be installed on the host machine to use this tool.

## To run 

#### Step 0: Clone the repo:
`git clone https://github.com/ogorman89/ccwc-go.git`

#### Step 1: Build the binary for your environment
`cd path/to/your/clone
go build -o ccwc main.go`

#### Step 2: Move the binary into your GOPATH
`sudo mv ccwc /Users/usr/go/bin`
your bin may be in a different location, if you need to find your path type `$GOPATH` in your terminal

#### Step 3: Use ccwc
You can now use ccwc via the command line
`ccwc`
