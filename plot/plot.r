library(ggplot2)
library(svglite)
data <- read.csv("time.csv")
data <- as.data.frame(data)
time = ggplot(data = data) +
	ggtitle("leng vs sort (lower is better)") +
	geom_line(size = 2.0, aes(x = size, y = lengSort, color = "leng.Sort")) +
	geom_line(size = 2.0, aes(x = size, y = lengSortInPlace, color = "leng.SortInPlace")) +
	geom_line(size = 2.0, aes(x = size, y = stl, color = "sort.Ints")) +
	labs(y = "ns / op", x = "Input size", color = "Algorithm") +
	theme_minimal(base_size=18) +
	theme(plot.title = element_text(size=20, face="bold", hjust = 0.5))
ggsave(file="time.svg", plot=time, width=10, height=6)
