# Greenhouse moisture module

## Installing the USB to serial adapter driver

Download for MacOSX: https://sparks.gogo.co.nz/ch340.html
You need to reboot your computer after you have installes the driver.
You should now be able to connect to the following port: `/dev/cu.SLAB_USBtoUART`

## Deploying MicroPython firmware to the ESP8266 board

```bash
$ pip install esptool
```

Erase the existing flash with:

```bash
$ esptool.py --port /dev/cu.SLAB_USBtoUART
```

Deploy the MicroPython image to your board:

```bash
$ esptool.py --port /dev/cu.SLAB_USBtoUART --baud 115200 --after no_reset write_flash --flash_mode dio 0x00000 /Users/mirko/Downloads/esp8266-20200911-v1.13.bin
```

Notice: Be sure that the ESP8266 board is not connectd to the arduino IDE, picocom or similar. This is to prevent error messages like `can not acces because the board is busy or already in use`

## calibrating the Soil Moisture Sensor V1.2

In order to get a reasonable percentage value out of the moisture value, you need to know the value when the sensore is completely under water and completeley dry.

1. Take a glas of water and insert the sensore but be sure that the upper, electronics part is not in the water.
2. Record the value as `WaterValue` in you main.py. You can deploy the main.py. This will aslo print the corresponding value.
3. Record the Value when the sensore is completely dry. this will be the `AirValue` in you main.py

For calculating the percentage you can use the following snippet.

```python
import machine
from machine import Pin
from machine import ADC
import time
from time import sleep

moisture = ADC(0)

AirValue = 554
WaterValue = 149


def myMap(x, in_min, in_max, out_min, out_max):
    return (x-in_min)*(out_max-out_min)/(in_max-in_min) + out_min


while True:
    moisture_value = moisture.read()
    print(moisture_value)
    print('%.2f' % myMap(moisture_value, AirValue, WaterValue, 0, 100), "%")
    sleep(3)
```

## Let ESP8266 connect to you router

The best solution would be, to put a connection function direct into your boot.py.

```python
# This file is executed on every boot (including wake-boot from deepsleep)
# import esp
# esp.osdebug(None)
import uos
import machine
import network
# uos.dupterm(None, 1) # disable REPL on UART(0)
import gc
# import webrepl
# webrepl.start()


def do_connect():
    sta_if = network.WLAN(network.STA_IF)
    if not sta_if.isconnected():
        print('connecting to network...')
        sta_if.active(True)
        sta_if.config(dhcp_hostname="greenhouse-satellite-01")
        sta_if.connect('xxx', 'xxx')
        while not sta_if.isconnected():
            pass
    print('network config:', sta_if.ifconfig())
    print('hostname: ', sta_if.config('dhcp_hostname'))


do_connect()
gc.collect()
```

Now your ESP8266 can automatically connect to you Router.
I also changed the the hostname in this script.

See also the official docs for Network basics:
http://docs.micropython.org/en/latest/esp8266/tutorial/network_basics.html?highlight=sta_if#network-basics

## Installing the Adafruit MicroPython Tool (ampy)

ampy allows you to interact(e.g upload or download) with a CircuitPython or MicroPython board over a serial connection.

```bash
$ pip install adafruit-ampy
```

See also: https://learn.adafruit.com/micropython-basics-load-files-and-run-code/install-ampy

## Upload the code to the ESP8266 board

```bash
$ ampy --port /dev/cu.SLAB_USBtoUART put main.py
```
