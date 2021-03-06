from gpiozero.pins.native import NativeFactory
from gpiozero import LED
from time import sleep
import logging

logging.basicConfig(
    format="%(asctime)s %(levelname)-8s %(message)s",
    level=logging.INFO,
    datefmt="%Y-%m-%d %H:%M:%S",
)
factory = NativeFactory()
soil_watering_pin = LED(21, pin_factory=factory)


def initialize_gpio_pins():
    logging.info("Initialize GPIO pins on the board.")
    soil_watering_pin.on()


def soil_watering(trigger_duration):
    _do_trigger(soil_watering_pin, "soil watering", trigger_duration)


def _do_trigger(pin, action, duration):
    logging.info(f"Activate {action} for {duration} seconds")
    logging.info(f"Start {action}")
    pin.off()
    sleep(duration)
    pin.on()
    logging.info(f"Stop {action}")
