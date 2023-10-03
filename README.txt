*) suppose $PROJ_ROOT is project root directory

*) make dir PROJ_ROOT/stream if not exists and put stream.proto file to it

*) go back to PROJ_ROOT

*) generate proto-file, run command protoc -I stream/ stream/stream.proto --go_out=plugins=grpc:stream

*) create server directory, to server/main.go import secion write:
    import (
        ...

        pb "../stream"

        ...
    )
    and write your server code

*) go run main.go - to start server

*) client
    create PROJ_ROOT/client
    write code client

    go run . stream
    go run . weight

*) if you don't have the scales device then you can emulate the device using socat utility

    run emulator:
        socat -d -d -d pty,raw,echo=0 pty,raw,eco=0

    read from:
        cat < /dev/pts/5

    write to:
        echo -e \\x0D\\x78 > /dev/pts/5
        echo -e \\x8F\\x66 > /dev/pts/5
        echo -e \\x8F\\xA0 > /dev/pts/5
