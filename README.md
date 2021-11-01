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

        $ sudo hostnamectl set-hostname greenhouse-proto
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

# Authenticate and Authorize an External User

Log into rasbperry and copy the content of the following cert file: /etc/kubernetes/pki/ca.crt and save it temporarily locally in a text file.
Then encode the text in the file locally as a base64 string:

    echo '<text from file>' | grep base64

Then use the base64 encoded string for a new cluster entry in the local ./kube/config.

    - cluster:
        certificate-authority-data:<base64 encodet string>
        server: https://greenhouse-proto:6443
      name: greenhouse-proto

Copy the following files via scp

    scp ubuntu@greenhouse-proto:/home/greenhouse/.certs/greenhouse.crt ~/<your local folder>/greenhouse.crt

    scp ubuntu@greenhouse-proto:/home/greenhouse/.certs/greenhouse.key ~/<your local folder>/greenhouse.key

Add the user and credentials to the kubeconfig file

    kubectl config set-credentials greenhouse-proto \
    --client-certificate=<your local folder>/greenhouse.crt \
    --client-key=<your local folder>/greenhouse.key

Add a new context for the new user in kubeconfig

    kubectl config set-context greenhouse-proto \
    --cluster=greenhouse-proto \
    --user=greenhouse-proto

Log in to the Rasbperry and assign admin privileges to the greenhous user:

    kubectl create clusterrolebinding greenhouse-proto --clusterrole=cluster-admin --user=greenhouse
