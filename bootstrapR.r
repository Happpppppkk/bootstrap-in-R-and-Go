library(boot)
library(pryr)

Rprof("bootstrapR.out")

data <- read.csv("data1.csv", header = FALSE)$V1
# Bootstrap function
bootstrap_mean <- function(data, indices) {
  return(mean(data[indices]))
}

set.seed(465)

samplesize <- 100000
start.time <- proc.time()

mem_used()

mem_change({results <- boot(data, statistic = bootstrap_mean, R = samplesize)

# Computing percentile confidence interval
conf_interval <- boot.ci(results, type = "perc", conf = 0.95)})
#Performing bootstrap

elapsed.time <- proc.time() - start.time


Rprof(NULL)


print(results)
print(conf_interval)
print(sprintf("Execution Time: %s", elapsed.time))
print(summaryRprof("bootstrapR.out"))
