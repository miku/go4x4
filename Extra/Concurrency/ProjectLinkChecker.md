# Parallel Link Checker

Given a list of URLs we want to check their status in parallel.

Overall design:

* use a channel to distribute tasks among workers
* each worker sends results to another channel
* one goroutine collects the results and writes them to stdout (e.g. as JSON)


