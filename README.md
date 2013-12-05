LEDSaid uses the asign library (https://github.com/krussell/asign) to send packets to Alpha LED Signs.

## Simple Hello World Example
	ledsaid -t "{SOT}{STX}{WriteText}{A}{Flash}{Red}Hello {Green}World{ETX}{EOT}"

## Demo Packet
This packet displays most features of the alpha sign protocol:  
https://github.com/tonymagro/ledsaid/blob/master/demo/demo.txt

## Read from a web address
You can load a file served over http using the -w flag

	ledsaid -w "https://raw.github.com/tonymagro/ledsaid/master/demo/demo.txt"
