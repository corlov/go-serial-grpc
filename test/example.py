#!/usr/bin/python3

import serial
import socket
from serial import *

# как весы определяются системой после подключения.
# По USB  соединению весы работают по эмуляции СОМ порта
DEVICE_ADDR = '/dev/ttyACM0'

DEVICE_IP_ADDR = '192.168.1.248'

TEST_USB = True

SP = serial.Serial()


class Scales:
    weigth = 0.0
    state = ""
    resolution = 0.0


def main():
    global SP

    if TEST_USB:
        try:
            SP = serial.Serial(port=DEVICE_ADDR, baudrate=4800, parity=serial.PARITY_EVEN)
        except OSError as os_err:
            print(str(os_err))
            return
        except Exception as e:
            print("Возможно, вы забыли подключить весы к компьютеру через USB интерфейс!")
            print(str(e))
            return
    
    #print(f"Масса на весах: {get_weigth()} грамм")
    print("\n\n")

    set_tare()
    print("\n\n")

    


def set_zero():
    global SP
    
    # header:
    SP.write(b'\xF8\x55\xCE')
    # length:
    SP.write(b'\x01\x00')
    # # command:
    SP.write(b'\x72')
    # CRC:
    SP.write(crc16_massa_k_edition(b'\x72', 1))

    time.sleep(1)
    
    print("recv answer:")
    #if SP.inWaiting() > 0:        
    header = int.from_bytes(SP.read(3), "little")
    print(header)
    len = int.from_bytes(SP.read(2), "little")
    print(len)
    cmd = int.from_bytes(SP.read(1), "little")      
    print(cmd)
    crc = int.from_bytes(SP.read(2), "little")
    print(crc)


def set_tare():
    global SP
    
    # header:
    SP.write(b'\xF8\x55\xCE')
    # length:
    SP.write(b'\x05\x00')
    # # command:
    SP.write(b'\xA3')

    SP.write(b'\x06\x00\x00\x00')
    # CRC:
    SP.write(crc16_massa_k_edition(b'\xA3\x06\x00\x00\x00', 1))

   
    
    print("recv answer:")
    #if SP.inWaiting() > 0:        
    header = int.from_bytes(SP.read(3), "little")
    print(header)
    len = int.from_bytes(SP.read(2), "little")
    print(len)
    cmd = int.from_bytes(SP.read(1), "little")      
    print(cmd)
    crc = int.from_bytes(SP.read(2), "little")
    print(crc)

    SP.close()
       
def get_weigth():
    global SP    

    # header:
    SP.write(b'\xF8\x55\xCE')
    # length:
    SP.write(b'\x01\x00')
    # # command:
    SP.write(b'\x23')
    # CRC:
    #SP.write(b'\x23\x00') # уже заранее посчитаная сумма для этой команды
    SP.write(crc16_massa_k_edition(b'\x23', 1))

    time.sleep(1)

    if SP.inWaiting() > 0:
        header = int.from_bytes(SP.read(3), "little")
        len = int.from_bytes(SP.read(2), "little")
        cmd = int.from_bytes(SP.read(1), "little")
        w = int.from_bytes(SP.read(4), "little")
        division = int.from_bytes(SP.read(1), "little")
        stable = int.from_bytes(SP.read(1), "little")
        net = int.from_bytes(SP.read(1), "little")
        zero = int.from_bytes(SP.read(1), "little")
        crc = int.from_bytes(SP.read(2), "little")

        # для подсчета КС формируем дата-сообщение:
        resp = bytearray()
        resp += cmd.to_bytes(1, 'little')
        resp += w.to_bytes(4, 'little')
        resp += division.to_bytes(1, 'little')
        resp += stable.to_bytes(1, 'little')
        resp += net.to_bytes(1, 'little')
        resp += zero.to_bytes(1, 'little')

        if crc16_massa_k_edition(resp, len) == crc:
            return w
        else:
            print("Не сходится контрольная сумма!")
            return None
   


# как расчитывают CRC!6 производители весов - ни один из стандартных алгоритомв расчета не подошел
def crc16_massa_k_edition(data: bytearray, data_len):
    crc = 0x0000
    for k in range(0, data_len):
        a = 0
        temp = (crc >> 8) << 8
        for bits in range(0, 8):
            if (temp ^ a) & 0x8000:
                a = (a << 1) ^ 0x1021
            else:
                a <<= 1
            temp <<= 1
        crc = a ^ (crc << 8) ^ (data[k] & 0xFF)

    crc = crc & 0xFFFF
    print(crc)
    return crc



if __name__ == "__main__":
    main()
