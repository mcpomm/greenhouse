# Setup the Raspberry

## Install Ubuntu Server

https://ubuntu.com/tutorials/how-to-install-ubuntu-on-your-raspberry-pi#1-overview

## Setup Wi-Fi directly from your SD card

Edit the file: /Volumes/system-boot/network-config

For ssh you can check the file: /Volumes/system-boot/user-data

For more information just visit this page:
https://desertbot.io/blog/headless-raspberry-pi-4-ssh-wifi-setup

## Raspberry Pi Linux kernel extra modules
Install linux-modules-extra-raspi - Raspberry Pi Linux kernel extra modules in order to have VXLAN available.
See also:
https://www.mail-archive.com/search?l=ubuntu-bugs@lists.ubuntu.com&q=subject:%22%5C%5BBug+1947628%5C%5D+Re%5C%3A+VXLAN+support+is+not+present+in+kernel+%5C-+Ubuntu+21.10+on+Raspberry+Pi+4+%5C%2864bit%5C%29%22&o=newest&f=1

        # Update the package index:
        $ sudo apt-get update
        # Install linux-modules-extra-raspi deb package:
        $ sudo apt-get install linux-modules-extra-raspi

## Setup authorized keys and hostname

1.  Log into the Raspberry, create the file ~/.ssh/authorized_keys(if not already present) and add the contents of your public key file (e.g. ~/.ssh/id_rsa_pub ).
2.  You can change the hosfile with the follofing command

        $ sudo nano /etc/hostname
        $ sudo nano /etc/hosts
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

    echo '<text from file>' | base64

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
