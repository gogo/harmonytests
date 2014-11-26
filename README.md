harmonytests
============

compatibility tests between gogoprotobuf and goprotobuf

[![Build Status](https://drone.io/github.com/gogo/harmonytests/status.png)](https://drone.io/github.com/gogo/harmonytests/latest)

This is done by generating structs using both from the same specification and marshalling and unmarshalling messages between them.

The idea is that this tests gogoprotobuf's marshalling and unmarshalling.
See https://code.google.com/p/gogoprotobuf

This project also tests JSON and Prototext.
