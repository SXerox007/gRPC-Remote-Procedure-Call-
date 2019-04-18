#key.pem or your file .der,.pem any
# 1) Generate RSA key:
openssl genrsa -out key.pem 1024 
openssl rsa -in key.pem -text -noout 

# 2) Save public key in pub.pem file:
openssl rsa -in key.pem -pubout -out pub.pem 
openssl rsa -in pub.pem -pubin -text -noout 

# 3) Encrypt some data:
echo test test test > file.txt 
openssl rsautl -encrypt -inkey pub.pem -pubin -in file.txt -out file.bin 

# 4) Decrypt encrypted data:
openssl rsautl -decrypt -inkey key.pem -in file.bin 

# It works like a charm



# Read the certificate file
openssl rsa -in /Users/sumitthakur/Downloads/m2psolutions_pub.der -pubout -out pub.pem

# Convert pem public key to ssh-public
ssh-keygen -f /Users/sumitthakur/go/src/gRPC-Remote-Procedure-Call-/pub.pem -i -mPKCS8
#or
ssh-keygen -f pub1key.pub -i

#
openssl genrsa -out priv.pem 2048