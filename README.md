# Setup the Raspberry

## Install Raspberry OS Lite (without desktop)

https://www.raspberrypi.com/software/

## Setup Wi-Fi directly from your SD card

Enable ssh to allow remote login

    $ touch /Volumes/boot/ssh

Add your WiFi network info

    $ touch /Volumes/boo/wpa_supplicant.conf

Paste the following lines into this file

    country=DE # your country code
    ctrl_interface=DIR=/var/run/wpa_supplicant GROUP=netdev
    update_config=1

    network={
        ssid="NETWORK-NAME"
        psk="NETWORK-PASSWORD"
    }

For more information just visit this page:
https://desertbot.io/blog/headless-raspberry-pi-4-ssh-wifi-setup

## Setup authorized keys and hostname

1.  Log into the Raspberry, create the file ~/.ssh/authorized_keys(if not already present) and add the contents of your public key file (e.g. /.ssh/id_rsa ).
2.  You can change the hosfile with the follofing command

        $ sudo raspi-config

    Go to System Options > Hostname

## Install k3s

Add "cgroup_memory=1 cgroup_enable=memory" to your linux cmdline (/boot/cmdline.txt on a Raspberry Pi)
Reboot your Pi

    $ sudo reboot

# Authenticate and Authorize an External User

    https://rancher.com/docs/k3s/latest/en/cluster-access/#accessing-the-cluster-from-outside-with-kubectl
