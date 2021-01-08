# Simple TCP Streamer

The repo provides a TCP server useful for paired programming interviews. The goal is to build a TCP client to handle the protocol defined below.

# Instructions

Build a client application that can consume a list of random numbers between 0-99 over a TCP connection and sort the list of numbers. The client application should be able to request a list of random numbers between 1 and 10 using this command:

```
14GET RANDNUM 10
```

The protocol format is text based and is structured as such:

|left padded string size|command string|

which the string size is always two characters with zero padding. The command string provides the request for the server to process. The server will start on port 64362 and can easily be tested like so:

```shell
echo -n "14GET RANDNUM 10" | nc servername 64362
```
