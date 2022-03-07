from py3rijndael import RijndaelCbc, ZeroPadding
from Crypto.Protocol import KDF
import zlib
import base64
import sys
import os
from hashlib import sha256
import struct

def getCipher():
    password = b"TKX73OHHK1qMonoICbpVT0hIDGe7SkW0"
    salt = b"71Ba2p0ULBGaE6oJ7TjCqwsls1jBKmRL"
    keyiv = KDF.PBKDF2(password,salt,64,10)
    key = keyiv[:32]
    iv = keyiv[32:]
    return RijndaelCbc(key,iv,ZeroPadding(32),32)

def inflate(data):
    decompress = zlib.decompressobj(-15)
    inflated = decompress.decompress(data)
    inflated += decompress.flush()
    return inflated

def deobfuscateFile(fileName):
    inFile = open(fileName,'r')
    buffer = inFile.read()
    buffer = buffer[3:-2]
    while (len(buffer) % 4) != 0:
        buffer += '='
    #print(buffer[0])
    encBytes = base64.b64decode(buffer,validate=False)
    #print(len(encBytes))
    encBytes = bytearray(encBytes)
    while(len(encBytes) % 32) != 0:
        encBytes.append(0)
    #print("encBytes")
    decBytes = getCipher().decrypt(bytes(encBytes))
    r = ""
    for b in decBytes:
        r = r + str(b) + ","
    #print(r)
    print(inflate(decBytes))

if __name__ == '__main__':
    globals()[sys.argv[1]](sys.argv[2])
