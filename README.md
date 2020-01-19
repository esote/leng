# Package `leng`

Package leng provides some sorting algorithms.

leng.Sort() performs around 4x better than the "sort" standard library package.
It scales well on multiple processors.

leng.SortInPlace() performs a bit worse than the "sort" package. The sort
package allocates 32 B/op, whereas SortInPlace will not allocate any memory.
This function isn't very useful.

"leng" is the Klingon verb "to travel."

# Comparison

![line graph](/plot/time.svg)
