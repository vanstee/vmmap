# vmmap from golang

Reading from `/proc/self/maps` on linux works just fine. But when trying to
shell out to `vmmap` on the current process, because of how `vmmap` freezes the
process, we never actually return.

So, this code reimplements vmmap using golang's cgo interface. The output was
stripped to display something similar to `cat /proc/self/maps`.

```
$ go run ./vmmap.go
0000000004000000-0000000004155000 00000000 /private/var/folders/cy/82bvkdg10rv9jrm3fvpc22j1ls09gf/T/go-build249262690/b001/exe/vmmap                                                                               
000000000f987000-000000000f9d2000 00000000 /usr/lib/dyld
00007ffffff74000-00007ffffff75000 00000000 /usr/lib/dyld
```
