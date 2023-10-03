package main

import "flag"

const PROTOCOL_HEADER = "\xF8\x55\xCE"
const CMD_SET_TARE = "\xA3"
const CMD_SET_TARE_LEN = "\x05\x00"
const CMD_SET_ZERO = "\x72"
const CMD_SET_ZERO_LEN = "\x01\x00"

const CMD_ERROR = 0x28
const CMD_ACK_SET = 0x27
const CMD_NACK_TARE = 0x15
const CMD_ACK_SET_TARE = 0x12
const CMD_NACK = 0xF0

// выдаем эти значения по запросам клиентов, если были запросы веса, состояния, то сюда записываем свежие показания с устройства
var scalesWeigth = 3448
var scalesState = "init message"

var (
	listenPort = flag.Int("listenPort", 50055, "The server port")
	//serialPortAddress = flag.String("serialPortAddress", "/dev/pts/5", "The scales address")
	serialPortAddress = flag.String("serialPortAddress", "/dev/ttyACM0", "The scales address")	
	serialBaudRate = flag.Int("serialBaudRate", 115200, "serialBaudRate")
)