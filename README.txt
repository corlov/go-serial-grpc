Привет. 
Посмотрели. Есть несколько просьб / моментов.
1) Сделать код-ревью и прочитать документацию еще раз. (Подсказка инетерсует реализация протокола 100, а не 2)
+ 2) Порт открыт, но не закрыт(Первый раз это выполнится, а второй и последующие - уже нет), 
3) Порт открыт без таймаута в конфигурации порта, но функция выдает ошибку таймаута. 
В реальности - чтение из порта будет бесконечным, пока что-то не прилетит из порта. 
Вывод: ошибки таймаута не произойдет никогда.

4) Если выполнить все инструкции из README,то следующий проект с grpc собрать уже будет очень тяжело. 
Нужно попробовать не модифицировать наш изначальный proto. 
В нем самом есть подсказка как его собрать.

protoc -I stream/ stream/stream.proto --go_out=plugins=grpc:stream


основное:
1) реализация протокола 100 и все протестить с реальным устройством


GetState  - состояние весов.(Подключены, не подключены)
Установит текущий вес тарой или отменит тару

rpc SetTareValue(RequestTareValue) установить значение тары в указанное значение?" -да.



rpc ScalesMessageOutChannel(stream RequestScale) returns (stream ResponseScale) {}

Установит текущий вес тарой или отменит тару
rpc SetTare(Empty) returns (ResponseSetScale) {}

rpc SetTareValue(RequestTareValue) returns (ResponseSetScale) {}

rpc SetZero(Empty) returns (ResponseSetScale) {}


================================================
1) скопировать сгенеренные протофайлы и сам каталог stream в каталог $HOME/go/src

    ls -las $HOME/go/src/stream/
    total 32
    4 drwxrwxr-x 2 kostya kostya  4096 сен 27 16:02 .
    4 drwxrwxr-x 7 kostya kostya  4096 сен 27 16:02 ..
    12 -rw-rw-r-- 1 kostya kostya 12208 сен 27 16:02 stream_grpc.pb.go
    12 -rw-rw-r-- 1 kostya kostya 11922 сен 27 16:02 stream.pb.go


2) в импорт написать pb "stream"

3) Сформировать файл stream.proto

4) сгенерировать новые файлы для ГО
        export PATH=$PATH:/home/kostya/go/bin
        protoc --go_out=stream --go-grpc_out=stream stream.proto 
        cp -r stream $HOME/go/src/
        ls -las $HOME/go/src/stream/

5) перекопировать полученные файлы в $HOME/go/src/helloworld
      вызывать сервер go run main.go  
      вызывать клиент go run main.go

6) для работы с сериал портом понадобится библиотека tarm
    go get github.com/tarm/serial

7)
    Эмулятор сериал порта
        socat -d -d -d pty,raw,echo=0 pty,raw,echo=0

// socat -d -d -d pty,raw,echo=0 pty,raw,echo=0
// cat < /dev/pts/5
// echo -e \\x0D\\x78 > /dev/pts/5
// echo -e \\x8F\\x66 > /dev/pts/5

// echo -e \\x8F\\xA0 > /dev/pts/5

=======================================================================================   


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
    