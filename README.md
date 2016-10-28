# power
Off the shelf, 100% TLS/SSL, Production Grade Cloud Based Iot Automation Solution


![](http://stream1.gifsoup.com/view2/4045951/homer-light-switch-o.gif)


This project aims to provide a DIY Lights Automation system with components you can find online with a really easy installation process. 
It can provide up to 64 light channels you can control from anywhere in the world. Your data is 100% encrypted. All devices report back to our cloud running 
on https://iot.cescoferraro.xyz where you receive REAL-TIME updates of the current state of all channels. You can name them, schedule half of them to 
shutdown at a specific schedule. Turn it all on in 10 minutes. You house, your rules.


### Hardware Requirements
  - 1 x [Ngrok Account](https://dashboard.ngrok.com/user/signup)
  - 1 x [Raspberry Pi2](https://www.adafruit.com/Raspberrypi?gclid=CjwKEAjw1qHABRDU9qaXs4rtiS0SJADNzJisDXC-HNL_Oqc7qxhBiP5F4IaOWsEsM_2xnzNg4GiesRoCG_zw_wcB)
  - 1 x [PiLeven Arduinod](http://www.freetronics.com.au/products/pileven-arduino-compatible-expansion-for-raspberry-pi#.WAk66nUrL0o)
  - Up to 8 x [Freetronics 8-Channel Relay Driver Shield](http://www.freetronics.com.au/collections/shields/products/relay8-8-channel-relay-driver-shield#.WAk7FHUrL0o)

### Software Requirements
  - Linux/OSX - Windows is supposibilly compatible too, I just dont care.
  - [Ansible](http://docs.ansible.com/)  
  - [HYpriot Flash tool](https://github.com/hypriot/flash)  


### Cost estimate

    Each shield has 8 channels. You can stack up to 8 chield for each rpi/arduino. Each channels needs a relay.

|   	                |  Avg. Price	|     |    Channel      |   8	    |   16   	|   24	    |   32	    |   40   	|   48  	|   56  	|   64  	|
|---	                |---		    |---  |---         	    |---	    |---	    |---	    |---	    |---	    |---	    |---	    |---     	|
|  Raspberry Pi2 	    |    $35.00     |     |  Hardware Price |  $87.53	|   $115	|   $109.93	|  $132.33 	|   $154.73	|  $177.13 	|   $199.53	|   $221.93	|
|  PiLeven Arduino 	    |    $30.13	    |     |  Price/Channel  |  $10.94	|   $7.18	|   $4.54	|  $4.13 	|   $3.86	|  $3.69 	|   $3.56	|   $3.46	|	
|  8-Channel Shield 	|    $22.40	    |     |  Relays         |  $10.94	|   $7.18	|   $4.54	|  $4.13 	|   $3.86	|  $3.69 	|   $3.56	|   $3.46	|
		

### Installation Dependencies

    curl -O https://raw.githubusercontent.com/hypriot/flash/master/$(uname -s)/flash
    chmod +x flash
    sudo mv flash /usr/local/bin/flash

### Installation 
Insert a sd-card to your computer and type this command, It will prompt you for the sd-card path.

    sudo -E flash -n ANYNAME -s WIFI-NETWORK -p WIFI-PASSWORD https://downloads.hypriot.com/hypriotos-rpi-v1.0.0.img.zip

Then put the sd-card into the Rpi and turn it on. Wait a couple of secconds and log into your RPi. Password if hyprio, you should change this later.
   
    ansible-playbook -i ansible/hosts ansible/ansible.yaml --extra-vars "target=ANYNAME.local"
 
    
