package main

import "flag"

const PROTOCOL_HEADER = "\xF8\x55\xCE"
const CMD_SET_TARE = "\xA3"
const CMD_SET_TARE_LEN = "\x05\x00"
const CMD_SET_ZERO = "\x72"
const CMD_SET_ZERO_LEN = "\x01\x00"
const CMD_GET_MASSA = "\x23"
const CMD_GET_MASSA_LEN = "\x01\x00"

const CMD_ERROR = 0x28
const CMD_ACK_SET = 0x27
const CMD_NACK_TARE = 0x15
const CMD_ACK_SET_TARE = 0x12
const CMD_NACK = 0xF0
const CMD_ACK_MASSA = 0x24

const ERR_CODE_OVERWEIGHT = 0x08
const ERR_CODE_INCORRECT_MODE = 0x09
const ERR_CODE_UNACCESS = 0x17
const ERR_CODE_PLATFORM = 0x18
const ERR_CODE_MODULE_FAULT = 0x19


// выдаем эти значения по запросам клиентов, если были запросы веса, состояния, то сюда записываем свежие показания с устройства
var scalesWeigth = 0


var (
	listenPort = flag.Int("listenPort", 50055, "The server port")

	// socat:
	//serialPortAddress = flag.String("serialPortAddress", "/dev/pts/5", "The scales emulator address")

	// scales device:
	serialPortAddress = flag.String("serialPortAddress", "/dev/ttyACM0", "The scales address")	

	serialBaudRate = flag.Int("serialBaudRate", 115200, "serialBaudRate")
)