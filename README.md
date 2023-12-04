# bootstrap-in-R-and-Go

# Week 10—Modern Applied Statistics with Go

## Statistics Method Implementation:
In exploring the comparison of R and Golang capabilities, the selected statistical method is bootstrap. R has a fully developed bootstrap resampling package, “boot”. However, in Go, I haven't found a similar package, necessitating manual creation in this project.
https://cran.r-project.org/web/packages/boot/index.html

## Univariate Data and Usage:
The objective of this project is to use the bootstrap resampling method to construct a confidence interval for a provided dataset. This dataset is sourced from UC Irvine's repository under the study “Bias correction of numerical prediction model temperature forecast”. I used one random column as the input data, constituting a univariate dataset with an unknown distribution.

The challenges of the project include the dataset volume and the absence of a Go library for implementing bootstrap. For R, theoretically, performance drastically decreases with large datasets in terms of runtime and memory usage.

## Bootstrap and CI calculation:

In the bootstrap function, both the size of the sample and the seed value of the random number generator are editable. In this project, both R and Go programs use a seed value of 465 to ensure result consistency, and both create 100,000 bootstrapped samples. After performing the bootstrap, the mean of all samples is used to calculate the 95% confidence interval using the percentile method.

## Software profiling:

For R, it comes with Rprofiler function to start the profiling in program and using “pryr” package for memory usage monitor. For Go, the runtime and time package provide memory usage and timing functions. Both programs are implemented with CPU and Memory usage test.

## Result:

Both program produced the same result: 

95% Confidence Interval for the Mean: [60.658973, 63.087930]

Go demonstrated faster processing time, with an Execution Time of 11.160163034s, compared to R's 26.572s. From this result, Go is nearly twice as fast as R. However, memory usage tells a different story: the R program only consumed 5.57 MB of memory, whereas the Go program used 6557217904 bytes (6557.217904 MB).

## Concern for Go and Recommendation:

Go is an excellent language to learn and use, designed for efficiency, speed, and concurrent performance. However, it also depends on the product requirements and design. For a cloud-based application that requires easy-to-use developed packages with comparable memory usage, it is reasonable to consider using the most cost-effective language. The lack of more developed packages is a significant hurdle for Go in gaining widespread popularity. It may take years for Go to develop similar or better statistical packages, but it's clear that Go could be beneficial for building microservices, DevOps, and backend programs, etc.


Reference:

https://archive.ics.uci.edu/dataset/514/bias+correction+of+numerical+prediction+model+temperature+forecast

https://bookdown.org/rdpeng/rprogdatascience/profiling-r-code.html#
