# Setup the Raspberry

## Install Ubuntu Server

https://ubuntu.com/tutorials/how-to-install-ubuntu-on-your-raspberry-pi#1-overview

## Headless Raspberry Pi 4 SSH WiFi Setup

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

        $ sudo hostnamectl set-hostname greenhouse
        $ sudo reboot

## Install Kubernetes via Ansible

Prepare for the installation

    $ cd ansible/
    $ vi ansible.cfg

Create the host groups which are to be supplied with Ansible and define the path to your private keyfile

    [defaults]
    inventory=/etc/ansible/hosts
    private_key_file=$HOME/.ssh/id_rsa
    deprecation_warnings=False

Add you SSH private key into the SSH authentication agent.

    $ ssh-add ~/.ssh/id_rsa

Apply the Ansible Playbook in order to install Kubernetes on you Raspberry.

    $ ansible-playbook raspberry.yaml
