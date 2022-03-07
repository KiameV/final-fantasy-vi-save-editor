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
    end = len(buffer) - 1
    for i in range(end, 0, -1):
        if buffer[i] != '=':
            end = i - 1
            break
    buffer = buffer[3:-2]
    while (len(buffer) % 4) != 0:
        buffer += '='
    encBytes = base64.b64decode(buffer,validate=False)
    encBytes = bytearray(encBytes)
    while(len(encBytes) % 32) != 0:
        encBytes.append(0)
    decBytes = getCipher().decrypt(bytes(encBytes))
    print(buffer[:3], inflate(decBytes))

def deflate(data):
    compress = zlib.compressobj(-1, zlib.DEFLATED, -15)
    deflated = compress.compress(bytes(data, "utf-8"))
    deflated += compress.flush()
    return deflated

def obfuscateFile(outFile, prefix, dataFile):
    df = open(dataFile,'r')
    data = df.read()
    deflated = deflate(data)
    encBytes = getCipher().encrypt(deflated)
    encoded = base64.b64encode(encBytes)
    outFile = open(outFile,'wb')
    result = str.encode(prefix) + encoded
    outFile.write(result)

if __name__ == '__main__':
    if len(sys.argv) == 3:
        globals()[sys.argv[1]](sys.argv[2])
    else:
        globals()[sys.argv[1]](sys.argv[2], sys.argv[3], sys.argv[4])
