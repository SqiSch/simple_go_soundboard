# simple_go_soundboard
A very simple and dirty soundboard written by copy pastes in golang for linux with an own API and javascript content. I use it to annoy my collegues if there in my office. I run this webapp on a raspy and a connected Speaker.


#### Dependencies:

##### Golang dependencies

go-json-rest API:
```
go get github.com/ant0ine/go-json-rest/rest
```

##### System dependencies:
**sox** - The Swiss Army knife of sound processing tools


install it under debian/ubuntu with:

```
apt-get install sox
```

or in archlinux:

```
pacman -S sox

```


#### Starting:

just type in the src folder:

```
go run player.go
```


#### Using:

the soundboard webapp binds to every internal ip so you can reach it in a browser with:

http://localhost:8080

#### API:

List all titles:
http://localhost:8080/api/list/

Playing a file:
http://localhost:8080/api/play/#Titlename/


#### Troubleshooting:

##### Changing the webapp port:

just open the player.go file and change the port 8080 at the end of the file to another port.

f.e.:
```
log.Fatal(http.ListenAndServe(":8081", nil))
```

##### run the webapp on port 80 without root privs:

use iptables to forward the traffic from port 80 to the webapp port:

f.e.:
```
```
