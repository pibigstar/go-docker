# go-docker
> 用go写一个docker

**注意**
> windows下要修改goland的OS环境为 linux,不然只会引用`exec_windows.go`而不会引用`exec_linxu_go`
> 在Setting->Go->Build Tags & Vendoring -> OS=linux

## 小记

- 查看Linux程序父进程
```bash
pstree -pl | grep main
```
- 查看进程id
```bash
echo $$
```
- 查看进程的uts
```bash
readling /proc/进程id/ns/uts
```
- 修改hostname
```bash
hostname -b 新名称
```