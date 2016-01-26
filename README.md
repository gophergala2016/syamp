![syamp](https://github.com/gophergala2016/syamp/blob/master/reVres/canvas/kay.png)
## syam{p}
<a href="https://youtu.be/QiSiYVDjEw4">Play Screencast</a>

syam{p} |pronounced as [simp]| brings the full power of the Linux Shell directly to your web browser. Run commands,
monitor and kill programs etc. - in real time.

Personal Benefit
--------------------
I built this app so that I could control and monitor programs on my Raspberypi B+ 2 running snappy ubuntu
core 15.04.
Currently, I use my laptop for this (I ssh into my pi). I hate that I have to move 
to my computer room to power off my pi when am outside hanging out with friends, having my lunch or just feeling
lazy to get out of bed when I am sleepy.

Now I can use my phone!

Note:
-----
You need to run syamp on the system you want to monitor.
syam{p} only supports Linux at the moment.

Installation
-------------
go get github.com/gophergala2016/syamp

Running
--------
go run syamp.go
<br>
go run syamp.go "192.168.1.2:2016"

Security
--------
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

User Interface
---------------
The user interface is simple, modern and easy to work with.

Dependencies
-------------
contype
<br>
https://github.com/kampsy/Go

License
-------
MIT - see license file.

Developer
---------
kampamba chanda (a.k.a kampsy).
