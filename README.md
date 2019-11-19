```
 _   _                                      
| |_(_)_ __ ___   ___  ___ _   _ _ __   ___ 
| __| | '_ ` _ \ / _ \/ __| | | | '_ \ / __|
| |_| | | | | | |  __/\__ \ |_| | | | | (__ 
 \__|_|_| |_| |_|\___||___/\__, |_| |_|\___|
                           |___/            
```

A simple timesync server/client using Cristian's algorithm, implemented with grpc.

Usage
=====

## server

    timesync server --bind :9999

## client

    timesync client --server X.X.X.X:9999
