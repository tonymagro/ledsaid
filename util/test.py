import usb
import logging
from logging import getLogger

log = getLogger(__name__)
log.setLevel(logging.DEBUG)
log.addHandler(logging.StreamHandler())

ALPHA_VENDOR_ID = 0x8765
ALPHA_PRODUCT_ID = 0x1234
ALPHA_INTERFACE_ID = 0x00

RECV_LENGTH    = 0x40 # Max Packet Size of 64
SEND_LENGTH    = 0x40 # Max Packet Size of 64
WRITE_ENDPOINT = 0x02 
READ_ENDPOINT  = 0x82
WRITE_TIMEOUT  = 5000
READ_TIMEOUT   = 5000

class Sign(object):
    def __init__(self, vendor_id, product_id, interface_id):
        self.vendor_id = vendor_id
        self.product_id = product_id
        self.iterface_id = interface_id
        self.device = self._get_device(self.vendor_id, self.product_id)
        self.handle = None
        self.no_write = False
        self._open()
        self.reset()

    def __del__(self):
        self._close()

    def reset(self):
        if self.handle == None:
            raise Exception("Tried to write with no valid device")

        out = "\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0\0"
        self.send_raw(out)

    def _get_device(self, vendor_id, product_id):
        self.device = None
        buses = usb.busses()
        for bus in buses:
            for device in bus.devices:
                if device.idVendor == vendor_id:
                    if device.idProduct == product_id:
                        log.info('Found device')
                        return device

        log.info('Failed to find device')
        return None

    def _open(self):
        self.handle = self.device.open()
        if self.handle == None:
            raise Exception("Could not open device")
        err = self.handle.claimInterface(0x00)
        if err != None:
            raise Exception("Could not claim interface")
        
    def send_raw(self, text):
        log.debug('SEND: '+text.encode('string_escape'))
        if not self.no_write:
            self.handle.bulkWrite(WRITE_ENDPOINT, text, WRITE_TIMEOUT)

    def _close(self):
        if self.handle != None:
            log.debug('Closing Sign')
            self.handle.releaseInterface()
            self.handle = None

sign = Sign(ALPHA_VENDOR_ID, ALPHA_PRODUCT_ID, ALPHA_INTERFACE_ID)
head = '\x00'*5 + '\x01' + 'Z00'
out = '\x02' + "AB" + "HELLO"
sign.send_raw(head + out + '\x04')