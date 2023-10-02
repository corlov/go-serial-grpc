package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"strconv"
	pb "stream"
	"time"

	"github.com/tarm/serial"
	"google.golang.org/grpc"
)

const PROTOCOL_HEADER = "\xF8\x55\xCE"
const CMD_SET_TARE = "\xA3"
const CMD_SET_TARE_LEN = "\x05\x00"

// выдаем эти значения по запросам клиентов, если были запросы веса, состояния, то сюда записываем свежие показания с устройства
var scalesWeigth = 3448
var scalesState = "init message"

var (
	listenPort = flag.Int("listenPort", 50055, "The server port")
	//serialPortAddress = flag.String("serialPortAddress", "/dev/pts/5", "The scales address")
	serialPortAddress = flag.String("serialPortAddress", "/dev/ttyACM0", "The scales address")	
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
	tareVal, err := strconv.Atoi(in.Message)
    if err != nil {
        return &pb.ResponseSetScale{ Error: "Incorrect tare value"}, nil
    }

	serialConfig := &serial.Config{
		Name: *serialPortAddress, 
		Baud: *serialBaudRate,
		ReadTimeout: time.Second*15,
	}	
	serial, err := serial.OpenPort(serialConfig)
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}

	n, err := serial.Write([]uint8(PROTOCOL_HEADER + CMD_SET_TARE_LEN + CMD_SET_TARE))
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}

	// tare:
	var tare[]uint8	
	_, err = serial.Write(int32ToSlice(uint32(tareVal), tare))
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}

	// crc: 
	_, err = serial.Write(crc16(append([]uint8(CMD_SET_TARE), tare...)))
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}

	fmt.Println("Command has sent")
	
	
	header := make([]uint8, 3)
	n, err = serial.Read(header)
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}
	log.Print("header: ", header[:n])

	len := make([]uint8, 2)
	n, err = serial.Read(len)
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}
	log.Print("len: ", len[:n])

	cmd := make([]uint8, 1)
	n, err = serial.Read(cmd)
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}
	log.Print("cmd: ", cmd[:n])


	crc := make([]uint8, 2)
	n, err = serial.Read(crc)
	if err != nil {
		return &pb.ResponseSetScale{ Error: err.Error()}, nil
	}
	log.Print("crc: ", crc[:n])

	serial.Close()


	// FIXME: проверить контрольную сумму

	if cmd[0] == 0x12 {
		return &pb.ResponseSetScale{ Error: ""}, nil	
	} else if cmd[0] == 0x15 {
		return &pb.ResponseSetScale{ Error: "nack"}, nil	
	} else {
		return &pb.ResponseSetScale{ Error: "code: " + strconv.Itoa(int(cmd[0]))}, nil
	}

}


// * В ряде весовых устройств команда не поддерживается (весовое устройство отвечает командой «CMD_NACK»).
func (s *server) SetZero(ctx context.Context, in *pb.Empty) (*pb.ResponseSetScale, error) {
	fmt.Println("SetZero")

	serialConfig := &serial.Config{
		Name: *serialPortAddress, 
		Baud: *serialBaudRate,
		ReadTimeout: time.Second*15,
	}	
	serial, err := serial.OpenPort(serialConfig)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// header:
	n, err := serial.Write([]uint8("\xF8\x55\xCE"))
	if err != nil {
		log.Fatal(err)
	}
	// length:
	_, err = serial.Write([]uint8("\x01\x00"))
	if err != nil {
		log.Fatal(err)
	}
	// command:
	_, err = serial.Write([]uint8("\xA3"))
	if err != nil {
		log.Fatal(err)
	}
	// crc: 
	_, err = serial.Write(crc16([]uint8("\xA3")))
	if err != nil {
		log.Fatal(err)
	}

	
	
	header := make([]uint8, 3)
	n, err = serial.Read(header)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("header: ", header[:n])

	len := make([]uint8, 2)
	n, err = serial.Read(len)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("len: ", len[:n])

	cmd := make([]uint8, 1)
	n, err = serial.Read(cmd)
	if err != nil {		
		log.Fatal(err)
	}
	// FIXME: приходит почему-то значение 240 а д.б. 27 и 28, и вес на весах не уходит в 0
	log.Print("cmd: ", cmd[:n])


	errorCode := make([]uint8, 1)
	if cmd[0] == 0x28 {
		n, err = serial.Read(errorCode)
		if err != nil {		
			// FIXME: тут надо отдавать сообщение клиенту!
			log.Fatal(err)
		}
		log.Print("errorCode: ", errorCode[:n])
	}
	

	crc := make([]uint8, 2)
	n, err = serial.Read(crc)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("crc: ", crc[:n])

	serial.Close()

	resp := make([]uint8, sliceToInt(len, "BE"))
	
	if cmd[0] == 0x28 {
		copy(resp, cmd)
		copy(resp[1:], errorCode)
	} else {
		copy(resp, cmd)
	}


	fmt.Println("calculated crc: ", (crc16(resp)))

	if !reflect.DeepEqual((crc16(resp)), crc) {
		return &pb.ResponseSetScale{ Error: "CRC error"}, nil
	}

	if cmd[0] == 0x28 {
		if errorCode[0] == 0x15 {
			return &pb.ResponseSetScale{ Error: "Setting to >0< is unavailable"}, nil
		}
		return &pb.ResponseSetScale{ Error: "code: " + strconv.Itoa(int(errorCode[0]))}, nil
	}

	return &pb.ResponseSetScale{ Error: ""}, nil		
}



// получить текущий вес с прибора
func (s *server) GetInstantWeight(ctx context.Context, in *pb.Empty) (*pb.ResponseInstantWeight, error) {
	serialConfig := &serial.Config{
		Name: *serialPortAddress, 
		Baud: *serialBaudRate,
		ReadTimeout: time.Second*15,
	}	
	serial, err := serial.OpenPort(serialConfig)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// header:
	n, err := serial.Write([]uint8("\xF8\x55\xCE"))
	if err != nil {
		log.Fatal(err)
	}
	// length:
	_, err = serial.Write([]uint8("\x01\x00"))
	if err != nil {
		log.Fatal(err)
	}
	// command:
	_, err = serial.Write([]uint8("\x23"))
	if err != nil {
		log.Fatal(err)
	}
	// crc: 
	_, err = serial.Write(crc16([]uint8("\x23")))
	if err != nil {
		log.Fatal(err)
	}

	
	header := make([]uint8, 3)
	n, err = serial.Read(header)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("header: ", header[:n])

	len := make([]uint8, 2)
	n, err = serial.Read(len)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("len: ", len[:n])

	cmd := make([]uint8, 1)
	n, err = serial.Read(cmd)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("cmd: ", cmd[:n])

	// FIXME: обработать случай, когда прибор возвращает тут ошибку!

	weight := make([]uint8, 4)
	n, err = serial.Read(weight)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("weight: ", weight[:n])

	division := make([]uint8, 1)
	n, err = serial.Read(division)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("division: ", division[:n])

	stable := make([]uint8, 1)
	n, err = serial.Read(stable)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("stable: ", stable[:n])

	net := make([]uint8, 1)
	n, err = serial.Read(net)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("net: ", net[:n])

	zero := make([]uint8, 1)
	n, err = serial.Read(zero)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("zero: ", zero[:n])

	crc := make([]uint8, 2)
	n, err = serial.Read(crc)
	if err != nil {		
		log.Fatal(err)
	}
	log.Print("crc: ", crc[:n])

	resp := make([]uint8, sliceToInt(len, "BE"))
	copy(resp, cmd)
	copy(resp[1:], weight)
	copy(resp[5:], division)
	copy(resp[6:], stable)
	copy(resp[7:], net)
	copy(resp[8:], zero)
	
	
	if !reflect.DeepEqual(reverse(crc16(resp)), crc) {
		panic("CRC error")
	}
                
	serial.Close()
	
	return &pb.ResponseInstantWeight{ Error: "", Message: strconv.Itoa(int(int32(sliceToInt32(weight, "BE"))))}, nil
}


// состояние весов.(Подключены, не подключены)
func (s *server) GetState(ctx context.Context, in *pb.Empty) (*pb.ResponseScale, error) {	
	serialConfig := &serial.Config{
		Name: *serialPortAddress, 
		Baud: *serialBaudRate,
		ReadTimeout: time.Second*15,
	}	

	serial, err := serial.OpenPort(serialConfig)
	if err != nil {
		fmt.Println(err)
		return &pb.ResponseScale{ Error: "1", Message: string(err.Error()), Type: "", Subtype: ""}, nil
	}
	serial.Close()

	return &pb.ResponseScale{ Error: "", Message: "ok", Type: "", Subtype: ""}, nil
}


func sendCommand(command string) ([]uint8, int, string, error) {
	serialConfig := &serial.Config{
		Name: *serialPortAddress, 
		Baud: *serialBaudRate,
		ReadTimeout: time.Second*15,
	}	
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
		return nil, 0, "read from serial error, " + err.Error(), err
		// log.Fatal(err)
	}
	log.Print("received: ", buf[:n])
	serial.Close()

	return buf[:n], n, "", nil
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



	


