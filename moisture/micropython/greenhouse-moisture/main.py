import machine
from machine import Pin
from machine import ADC
import time
from time import sleep
try:
    import usocket as socket
except:
    import socket

moisture = ADC(0)

AirValue = 560
WaterValue = 162

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('', 5200))
s.listen(5)


def myMap(x, in_min, in_max, out_min, out_max):
    return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min


def getMoisture():
    moisture_value = moisture.read()
    print(moisture_value)
    print('%.2f' % myMap(moisture_value, AirValue, WaterValue, 0, 100), "%")
    return str('%.2f' % myMap(moisture_value, AirValue, WaterValue, 0, 100)) + " %"


while True:
    conn, addr = s.accept()
    request = conn.recv(1024)
    if 'favicon' in str(request):
        continue
    print('Got a connection from %s' % str(addr))
    request = str(request)
    print('Content = %s' % request)
    response = getMoisture()
    conn.send('HTTP/1.1 200 OK\n')
    conn.send('Content-Type: application/json\n')
    conn.send('Connection: close\n\n')
    conn.sendall(response)
    conn.close()
