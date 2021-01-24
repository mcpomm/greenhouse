#!/usr/bin/python

# Import required Python libraries
import RPi.GPIO as GPIO

# We will be using the BCM GPIO numbering
GPIO.setmode(GPIO.BCM)

# Select a control GPIO
GPIO_CONTROL = 21

# Set CONTROL to OUTPUT mode
GPIO.setup(GPIO_CONTROL, GPIO.IN)

# Main function


def main():
    # Start by setting the relay to OFF
    # relayState = False
    # GPIO.output(GPIO_CONTROL, relayState)
    # print('Relay is OFF')
    # state = GPIO.input(GPIO_CONTROL)
    try:
        # Repeat till the program is ended by the user
        while True:
            state = GPIO.input(GPIO_CONTROL)
            if (state is True):
                print('State is ON')
            else:
                print('Relay is OFF')
                # Wait while ENTER is pressed
                # This is a small hack, raw_input() can read keyboard input until you press ENTER and save it to a variable, but we dont care about other keys to be pressed, so we don't
                # raw_input()
                # # Toggle the relayState value
                # relayState = not(relayState)
                # # Change the CONTROL output value
                # GPIO.output(GPIO_CONTROL, relayState)
                # # Print state to console
                # if relayState:
                #     print('Relay is ON')
                # else:
                #     print('Relay is OFF')
    # If the program is ended cleanup GPIOs
    except KeyboardInterrupt:
        GPIO.cleanup()


# Run the main function when the script is executed
if __name__ == "__main__":
    main()
