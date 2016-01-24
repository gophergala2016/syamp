![syamp](https://github.com/gophergala2016/syamp/blob/master/reVres/canvas/kay.png)
## syam{p}

syamp pronounced simp brings the full power of the Linux Shell directly to your web browser. Run commands,
moniter and kill programms etc in real time.

## Personal Benefit
I built this app so that i can control and moniter programs on my Raspberypi B+ 2 running snappy ubuntu
core 15.04.
Currently, i use my laptop for this(i ssh into my pi). I hate that i have to move 
to my computer room to power of my pi when am outside hunging out with friends, having my lunch or just feeling
lazy to get out of bed when iam sleepy.

Now i can use my phone!

Note:
-----
You need to run syamp on the system you want to control.
syam{p} only supports linux at the moment.

## Installation
go get github.com/gophergala2016/syamp

## Running
go run syamp.go
<br>
go run syamp.go "192.168.1.2:2016"

## Security
syam{p} runs an https server. All your commands are encypted with a 2048bit key in transit.
![login](https://github.com/gophergala2016/syamp/blob/master/img/keys.png)

<h3>Login Page</h3>
<code>UserID: gopher</code>
<br>
<code>Password: root</code>
![login](https://github.com/gophergala2016/syamp/blob/master/img/login.png)
<hr>
![home](https://github.com/gophergala2016/syamp/blob/master/img/home.png)
<hr>
![home](https://github.com/gophergala2016/syamp/blob/master/img/home2.png)

### User Interface
The User interface is simple, morden and easy to work with.

## License
MIT see license file

Developer
---------
kampamba chanda (a.k.a kampsy)
