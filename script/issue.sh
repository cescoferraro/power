#!/usr/bin/env bash



$ cat /etc/inittab
# Spawn a getty on Raspberry Pi serial line
T0:23:respawn:/sbin/getty -L ttyAMA0 115200 vt100



$ cat /etc/inittab
# Spawn a getty on Raspberry Pi serial line



$ cat /boot/cmdline.txt
+dwc_otg.lpm_enable=0 console=tty1 root=/dev/mmcblk0p2 rootfstype=ext4 cgroup_enable=memory swapaccount=1 elevator=deadline fsck.repair=yes rootwait console=ttyAMA0,115200 kgdboc=ttyAMA0,115200

$ cat /boot/cmdline.txt
+dwc_otg.lpm_enable=0 console=tty1 root=/dev/mmcblk0p2 rootfstype=ext4 cgroup_enable=memory swapaccount=1 elevator=deadline fsck.repair=yes rootwait

