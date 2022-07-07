# nmap-to-ports
Easy as that - nmap to ports notation.

# Usage
```
nmap -T4 -iL ranges.txt -vv | nmap-to-ports | httpx  
nmap -T4 -iL ranges.txt -vv | nmap-to-ports | httpx | nuclei 
```

By using and grepping for "discovered open port", this tool just converts it to host:port notation so that you can easily pass it to httpx.

## Installation

If you have a properly configured GOPATH and $GOPATH/bin is in your PATH, then run this command for a one-liner install, thank you golang!

```
go install  github.com/pry0cc/nmap-to-ports@latest
```


