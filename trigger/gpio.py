from gpiozero import LED
from time import sleep
import logging

logging.basicConfig(
    format="%(asctime)s %(levelname)-8s %(message)s",
    level=logging.INFO,
    datefmt="%Y-%m-%d %H:%M:%S",
)

soil_watering_pin = LED(21)


def initialize_gpio_pins():
    logging.info("Initialize GPIO pins on the board.")
    soil_watering_pin.on()


def soil_watering(trigger_duration):
    logging.info(f"Activate soil watering for {trigger_duration} seconds")
    logging.info("Start soil watering")
    soil_watering_pin.off()
    sleep(trigger_duration)
    soil_watering_pin.on()
    logging.info("Stop soil watering")
