#!/usr/bin/env bash

apt-get update
apt-get install libmpc-dev wiringpi

sed -i '/ttyAMA0/d' /etc/inittab
sed -i 's/[a-z]*=ttyAMA0\.\?[^ ]* *//g' /boot/cmdline.txt
/bin/systemctl disable serial-getty@ttyAMA0.service # silent if doesn't exist
/bin/systemctl disable serial-getty@serial0.service
cat > /etc/udev/rules.d/99-PiLeven-ttyAMA0-config.rules <<EOF
KERNEL=="ttyAMA0", ACTION=="add", SYMLINK+="ttyS99", RUN+="/usr/bin/gpio -g mode 14 alt0", RUN+="/usr/bin/gpio -g mode 15 alt0", GROUP="dialout"
EOF
sudo python -c "$(curl -fsSL https://raw.githubusercontent.com/platformio/platformio/master/scripts/get-platformio.py)"
git clone https://github.com/cescoferraro/power /home/pirate
cd power
platformio run
















sed -i "s/[a-z]*=${SERIAL}[0-9,]* *//g" /boot/cmdline.txt
sed -i 's/.\+ttyAMA0/#\0/' /etc/inittab
/bin/systemctl disable serial-getty@ttyAMA0.service # silent if doesn't exist
/bin/systemctl disable serial-getty@serial0.service
KERNEL=="ttyAMA0", ACTION=="add", SYMLINK+="ttyS99", RUN+="/usr/bin/gpio -g mode 17 alt3", RUN+="/usr/bin/gpio -g mode 14 alt0", RUN+="/usr/bin/gpio -g mode 15 alt0", GROUP="dialout"