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
- 常看当前用户和用户组
```bash
id
```
- 创建并挂载一个hierarchy
> 在这个文件夹下面创建新的文件夹，会被kernel标记为该`cgroup`的子`cgroup`
```bash
mkdir cgroup-test
mount -t cgroup -o none,name=cgroup-test cgroup-test ./cgroup-test
```
- 将其他进程移动到其他的`cgroup`中
> 只要将该进程的ID放到其`cgroup`的`tasks`里面即可
```bash
echo "进程ID" >> cgroup/tasks 
```