package main

import "flag"

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