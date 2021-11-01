# Greenhouse soil temperature module

## Installing the USB to serial adapter driver

Download for MacOSX: https://sparks.gogo.co.nz/ch340.html
You need to reboot your computer after you have installes the driver.
You should now be able to connect to the following port: `/dev/cu.SLAB_USBtoUART`

## Deploying MicroPython firmware to the NodeMCU ESP8266 board

Please note that I had some problems uploading files with adafruit-ampy when my NodeMCU was flashed with esptool v 3.0.

```bash
$ pip install esptool==2.8
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

import upip

def do_install():
    upip.install('picoweb')


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
do_install()
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

# Use this module with GOLANG

## build

`env GOOS=linux GOARCH=arm GOARM=6 go build`

## upload to greenhouse-test

`rsync -avz temperature pi@greenhouse-test:greenhouse/soil-temperature`

## Docker

`docker build -t soil-temperature . --platform linux/arm/v7`
`docker tag soil-temperature:latest mcpomm/greenhouse-soil-temperature:<tag>`
`docker push mcpomm/greenhouse-soil-temperature:<tag>`

## one wire setup

`$ sudo vi /boot/config.txt`
Add the following line add at the end of the file.

`dtoverlay=w1-gpio,gpiopin=4`

This enables you to handle multiple sensors with a single GPIO pin.(we use GPIO 4)

Then reboot the pi.

Each sensor has a uniq ID. After the rebbot you'll find this id's under:

`$ ls -la /sys/bus/w1/devices/`

```
pi@greenhouse:~ $ ls -la /sys/bus/w1/devices/
total 0
drwxr-xr-x 2 root root 0 Feb 26 23:17 .
drwxr-xr-x 4 root root 0 Feb 26 23:17 ..
lrwxrwxrwx 1 root root 0 Feb 27 15:34 28-01203390917a -> ../../../devices/w1_bus_master1/28-01203390917a
lrwxrwxrwx 1 root root 0 Feb 27 15:34 28-0215c2c3c3ff -> ../../../devices/w1_bus_master1/28-0215c2c3c3ff
lrwxrwxrwx 1 root root 0 Feb 26 23:17 w1_bus_master1 -> ../../../devices/w1_bus_master1
```

You can identify the corresponding ID by connecting the sensors one after another to the GPIO pin 4.
