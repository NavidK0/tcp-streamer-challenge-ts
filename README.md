# Simple TCP Streamer

This is a simple TCP server, useful for pair programming interviewers. It works on fixed size length of 00-99. The format is all text based. 

# Instructions

Build a client application that can consume a list of random numbers between 0-99 over a TCP connection and sort the list of numbers. The client application should be able to request a list of random numbers between 1 and 10 using this command:

```
14GET RANDNUM 10
```

This is a text based protocol with the length of each request appended to the front. The server will start on localhost port 64362 and can easily be tested like so:

```shell
echo -n "14GET RANDNUM 10" | nc localhost 64362
```