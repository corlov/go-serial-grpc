*) сгенерировать новые файлы для go
        export PATH=$PATH:/home/kostya/go/bin
        protoc --go_out=stream --go-grpc_out=stream stream.proto 
        cp -r stream $HOME/go/src/
        ls -las $HOME/go/src/stream/

*) скопировать сгенеренные прото-файлы и сам каталог stream в каталог $HOME/go/src/stream/
Д.б. так, иначе не увидятся прото-файлы
    ls -las $HOME/go/src/stream/
    total 32
    4 drwxrwxr-x 2 kostya kostya  4096 сен 27 16:02 .
    4 drwxrwxr-x 7 kostya kostya  4096 сен 27 16:02 ..
    12 -rw-rw-r-- 1 kostya kostya 12208 сен 27 16:02 stream_grpc.pb.go
    12 -rw-rw-r-- 1 kostya kostya 11922 сен 27 16:02 stream.pb.go

*) в импорт написать pb "stream"

*) для работы с сериал портом понадобится библиотека tarm
    go get github.com/tarm/serial

*) Эмулятор сериал порта
        socat -d -d -d pty,raw,echo=0 pty,raw,echo=0

// socat -d -d -d pty,raw,echo=0 pty,raw,echo=0
// cat < /dev/pts/5
// echo -e \\x0D\\x78 > /dev/pts/5
// echo -e \\x8F\\x66 > /dev/pts/5
// echo -e \\x8F\\xA0 > /dev/pts/5

========================================================================================================





*) install the following gRPC gen plugins
(https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable)

    go get -u google.golang.org/protobuf/cmd/protoc-gen-go
    go install google.golang.org/protobuf/cmd/protoc-gen-go

    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

*) Посмотреть куда установилось все:

    sudo find / -iname protoc-gen-go-grpc

*) установить переменную PATH, в моем случае в значения 

    export PATH=$PATH:/home/kostya/go/bin

*) Сгенерить сообщения и сервисы командой:
    protoc --go_out=./stream/ --go-grpc_out=./stream/ stream.proto
    
    # protoc -I ./stream/ --go_out=./stream/ --go-grpc_out=./stream/ ./stream/stream.proto
    