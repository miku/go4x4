# Project Filetransfer

Let's write a tool to transfer a file via FTP. A kind of ftpcat.

We will use an external library for FTP:

* [github.com/jlaffaye/ftp](https://github.com/jlaffaye/ftp)

Write a tool that has usage similar to this:

```
// Usage of /tmp/go-build3945962022/b001/exe/main:
//   -H string
//         ftp hostname (default "ftp.ncbi.nlm.nih.gov:21")
//   -L    use lftp
//   -P string
//         password (default "anonymous")
//   -U string
//         username (default "anonymous")
//   -f string
//         filepath to retrieve (default "/pub/pmc/readme.txt")
```

For example, running with the defaults, it would print the contents of a file to
stdout (abbreviated below):

```
$ go run cmd/filefetch/main.go 
On March 18, 2019, PMC will no longer provide bulk packages of ...
files were superseded in August 2016 by the Commercial Use and ...
articles that may be used for commercial purposes (the Commerci...
purposes. Anyone planning to use OA subset content for non-comm...
access the complete collection. See the Open Access Subset page...
be directed to pubmedcentral@ncbi.nlm.nih.gov.

See http://www.ncbi.nlm.nih.gov/pmc/tools/ftp/ 
```
