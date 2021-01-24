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
        sta_if.config(dhcp_hostname="greenhouse-satellite-03")
        sta_if.connect('xxx', 'xxx')
        while not sta_if.isconnected():
            pass
    print('network config:', sta_if.ifconfig())
    print('hostname: ', sta_if.config('dhcp_hostname'))


do_connect()
gc.collect()
