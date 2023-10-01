package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	pb "stream"

	"github.com/tarm/serial"

	"google.golang.org/grpc"
)

// выдаем эти значения по запросам клиентов, если были запросы веса, состояния, то сюда записываем свежие показания с устройства
var scalesWeigth = 3448
var scalesState = "init message"

var (
	listenPort = flag.Int("listenPort", 50055, "The server port")
	serialPortAddress = flag.String("serialPortAddress", "/dev/pts/5", "The scales address")
	serialBaudRate = flag.Int("serialBaudRate", 115200, "serialBaudRate")
)



type server struct {
	pb.UnimplementedApiCallerScaleServer
}


// Двунаправленный потоковый RPC (Bidirectional streaming RPC), 
func (s *server) ScalesMessageOutChannel(srv pb.ApiCallerScale_ScalesMessageOutChannelServer) (error) {
	log.Println("start new server")
	
	ctx := srv.Context()

	for {
		// exit if context is done
		// or continue
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}


		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// return will close stream from server side
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}

		if req.Message == "weigth" {
			resp := pb.ResponseScale{ Error: "", Message: strconv.Itoa(int(scalesWeigth)), Type: "", Subtype: ""}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}			
		}

		resp := pb.ResponseScale{ Error: "", Message: "", Type: "", Subtype: ""}
		switch req.Message {
			case "weigth":
				fmt.Println("get weigth")
				resp = pb.ResponseScale{ Error: "", Message: strconv.Itoa(int(scalesWeigth)), Type: "", Subtype: ""}
			case "state":
				fmt.Println("get state")
				resp = pb.ResponseScale{ Error: "", Message: scalesState, Type: "", Subtype: ""}
		}

		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}			
	}
	
}



func (s *server) SetTare(ctx context.Context, in *pb.Empty) (*pb.ResponseSetScale, error) {
	_, _, errText, _ := sendCommand("\x0D")	
	if errText != "" {
		return &pb.ResponseSetScale{ Error: errText}, nil
	}	
	return &pb.ResponseSetScale{ Error: ""}, nil
}



func (s *server) SetTareValue(ctx context.Context, in *pb.RequestTareValue) (*pb.ResponseSetScale, error) {
	// TODO: данная команда доступна только по протоколу 100, по протоколу 2 недоступна
	return &pb.ResponseSetScale{ Error: ""}, nil
}



func (s *server) SetZero(ctx context.Context, in *pb.Empty) (*pb.ResponseSetScale, error) {
	_, _, errText, _ := sendCommand("\x0E")	
	if errText != "" {
		return &pb.ResponseSetScale{ Error: errText}, nil
	}	
	return &pb.ResponseSetScale{ Error: ""}, nil
}



// получить текущий вес с прибора
func (s *server) GetInstantWeight(ctx context.Context, in *pb.Empty) (*pb.ResponseInstantWeight, error) {
	buf, n, errText, _ := sendCommand("\x45")	
	if errText != "" {
		return &pb.ResponseInstantWeight{ Error: errText, Message: ""}, nil
	}	

	weight := int16(sliceToInt(buf[:n]) & 0x7FFF)
    minusFlag := (sliceToInt(buf[:n]) & 0x8000) != 0
    if minusFlag {
      weight *= -1
    }

	scalesWeigth = int(weight)
	
	return &pb.ResponseInstantWeight{ Error: "", Message: strconv.Itoa(int(weight))}, nil
}



func (s *server) GetState(ctx context.Context, in *pb.Empty) (*pb.ResponseScale, error) {
	log.Printf("GetState")

	buf, n, errText, _ := sendCommand("\x44")	
	fmt.Println(errText)
	if errText != "" {
		return &pb.ResponseScale{ Error: errText, Message: "", Type: "", Subtype: ""}, nil
	}
	responseData := sliceToInt(buf[:n])
	/*
		D7 – индикатор процесса взвешивания: 0 – не завершен, 1– завершен;
		D6 – индикатор : 0 – не высвечен, 1 – высвечен;
		D5 – индикатор : 0 – не высвечен, 1 – высвечен;
		D15-D8, D4-D0 – неопределенное состояние 
	*/
	D7 := ((responseData >> 7) & 1) != 0
    D6 := ((responseData >> 6) & 1) != 0
    D5 := ((responseData >> 5) & 1) != 0

	stable := "0"
	if D7 { stable = "1" }

	zero := "0"
	if D6 { zero = "1" }

	net := "0"
	if D5 { net = "1" }

	respMessageJson := "{ \"stable\": \"" + stable + "\", \"zero\": \"" + zero + "\", \"net\": \"" + net + "\"} "
	scalesState = respMessageJson

	return &pb.ResponseScale{ Error: "", Message: respMessageJson, Type: "scales", Subtype: "state"}, nil
}



func sendCommand(command string) ([]uint8, int, string, error) {
	serialConfig := &serial.Config{Name: *serialPortAddress, Baud: *serialBaudRate}	
	serial, err := serial.OpenPort(serialConfig)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		return nil, 0, "open serial error, " + err.Error(), err

	}

	n, err := serial.Write([]uint8(command))
	if err != nil {
		return nil, 0, "write to serial error, " + err.Error(), err
		// log.Fatal(err)
	}

	// получаем значение (2 байта)
	buf := make([]uint8, 2)
	n, err = serial.Read(buf)
	if err != nil {		
		return nil, 0, "read request timeout, " + err.Error(), err
		// log.Fatal(err)
	}
	log.Print("received: ", buf[:n])

	return buf[:n], n, "", nil
}



func sliceToInt(s []uint8) uint16 {
    res := uint16(0)
	shift := 8
    for i := 0; i < len(s); i++ {
		val := uint16(s[i])
		res += uint16(val << uint8(shift*(len(s) - 1 - i)))
		// fmt.Println("s[i]:", s[i], " i:", i, " shift: ", shift*(len(s) - 1 - i), " res:", uint16(val << uint16(shift*(len(s) - 1 - i))))
    }
	fmt.Println(res)
    return res
}



func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApiCallerScaleServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

