# heap
go 1.18+ 简单易用的最大最小堆

example
---

```golang
c := 5
heap := MaxHeap(c, func(a, b int) bool {
    return a < b
})
heap.Init([]int{2333, 223, 123, 11, 21, 14})
for c > 0 {
    fmt.Println(heap.Pop())
    c--
}

```
#### output
```
2333
223
123
11
21
14
```
