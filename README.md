# Simple TCP Streamer

The repo provides a TCP server for the TCP Streamer Challenge (modified for TypeScript). The goal is to build a TCP
client to handle the protocol
defined below.

# Instructions

Build a client application that can consume a list of random numbers between 0-99 over a TCP connection and sort the
list of numbers. The client application should be able to request a specifically sized list of random numbers
command:

```
14GET RANDNUM 10
```

Where `14GET` is the command, `RANDNUM` is the type of data to retrieve, and `10` is the number of random numbers to
retrieve.

The protocol format is text based and is structured as such:

|left padded string size|command string|

which the string size is always two characters with zero padding. The command string provides the request for the server
to process. The server will start on port 64362 and can easily be tested like so:

```shell
echo -n "14GET RANDNUM 10" | nc servername 64362
```
