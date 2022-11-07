# Project Reservoir Sampling

Implement reservoir sampling. The idea is the following:

From an potential infinite stream of data, we want to take a random sample.

* we have a "reservoir", a fixed number of elements
* if we have seen fewer elements than the reservoir size, we just add the element
* if we have seen more element than the reservoir size, we replace one random element with probability (reservoir size / elements seen)

That's it.

* implement a small tool, that would allow to stream data from stdin and would output a random sample
* allow the reservoir size to be altered

