import machine
import onewire
import ds18x20
from struct import unpack
import urandom
import ujson
import utime
import ntptime
try:
    import usocket as socket
except:
    import socket

ntptime.settime()
machine.RTC().datetime()

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind(('', 5300))
s.listen(5)

ds_pin = machine.Pin(4)
ds_sensor = ds18x20.DS18X20(onewire.OneWire(ds_pin))

roms = ds_sensor.scan()

soitemperature = {
    "ID":     str(urandom.getrandbits(30)),
    "Name":   "Soil Temperature",
    "Value":  "",
    "Unit":   "°C",
    "Time":   ""
}

# while True:
#     ds_sensor.convert_temp()
#     time.sleep_ms(750)
#     for rom in roms:
#         print(ds_sensor.read_temp(rom))
#     time.sleep(5)


def getTemperature():
    ds_sensor.convert_temp()
    readtime = 946684800 + utime.time()
    soitemperature["Time"] = str(readtime)
    soitemperature["Value"] = str(ds_sensor.read_temp(roms[0]))
    print(ds_sensor.read_temp(roms[0]))
    return ujson.dumps(soitemperature)
    # return str(ds_sensor.read_temp(roms[0]))  # ds_sensor.read_temp(roms[0])


while True:
    conn, addr = s.accept()
    request = conn.recv(1024)
    if 'favicon' in str(request):
        continue
    print('Got a connection from %s' % str(addr))
    request = str(request)
    print('Content = %s' % request)
    response = getTemperature()
    conn.send('HTTP/1.1 200 OK\n')
    conn.send('Content-Type: application/json\n')
    conn.send('Connection: close\n\n')
    conn.sendall(response)
    conn.close()
