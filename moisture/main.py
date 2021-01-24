import machine
from machine import Pin
from machine import ADC
import urandom
import ujson
import utime
import ntptime
try:
    import usocket as socket
except:
    import socket

moisture = ADC(0)

AirValue = 562
WaterValue = 134

ntptime.settime()
machine.RTC().datetime()


soilmoisture = {
    "ID":     str(urandom.getrandbits(30)),
    "Name":   "Soil Moisture",
    "Value":  "",
    "Unit":   "%",
    "Time":   ""
}

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('', 5200))
s.listen(5)


def myMap(x, in_min, in_max, out_min, out_max):
    return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min


def getMoisture():
    readtime = 946684800 + utime.time()
    moisture_value = moisture.read()
    print(moisture_value)
    print('%.2f' % myMap(moisture_value, AirValue, WaterValue, 0, 100), "%")
    soilmoisture["Time"] = str(readtime)
    soilmoisture["Value"] = str('%.2f' % myMap(
        moisture_value, AirValue, WaterValue, 0, 100))
    return ujson.dumps(soilmoisture)


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
