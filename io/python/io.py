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

def deobfuscateFile(file_name, omit_first_bytes):
    in_file = open(file_name,'r')
    buffer = in_file.read()
    if omit_first_bytes == "1":
        buffer = buffer[3:]
    while (len(buffer) % 4) != 0:
        buffer += '='
    enc_bytes = base64.b64decode(buffer, validate=False)
    enc_bytes = bytearray(enc_bytes)
    while(len(enc_bytes) % 32) != 0:
        enc_bytes.append(0)
    enc_bytes = getCipher().decrypt(bytes(enc_bytes))
    print(inflate(enc_bytes))

def deflate(data):
    compress = zlib.compressobj(-1, zlib.DEFLATED, -15)
    deflated = compress.compress(bytes(data, "utf-8"))
    deflated += compress.flush()
    return deflated

def obfuscateFile(out_file, data_file):
    df = open(data_file, 'r')
    data = df.read()
    print(data)
    deflated = deflate(data)
    enc_bytes = getCipher().encrypt(deflated)
    encoded = base64.b64encode(enc_bytes)
    f = open(out_file, 'wb')
    f.write(encoded)
    f.write(b"\r\n")


if __name__ == '__main__':
    globals()[sys.argv[1]](sys.argv[2], sys.argv[3])
