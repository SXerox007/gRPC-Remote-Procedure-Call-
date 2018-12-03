// #generate gRPC-Code
//To generate the grpc code 
protoc expert/error-handling/proto/search.proto --go_out=plugins=grpc:.

// #server-run
go run expert/ssl-security/server/app.go
or 
make server

// #client-run
go run expert/ssl-security/client/app.go
or 
make client

// #generate the ssl sercurity files 
# Summary 
# Private files: ca.key, server.key, server.pem, server.crt
# "Share" files: ca.crt (needed by the client), server.csr (needed by the CA)

# Changes these CN's to match your hosts in your environment if needed.
SERVER_CN=localhost

# Step 1: Generate Certificate Authority + Trust Certificate (ca.crt)
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096
openssl req -passin pass:1111 -new -x509 -days 3650 -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"

# Step 2: Generate the Server Private Key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

# Step 3: Get a certificate signing request from the CA (server.csr)
openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"

# Step 4: Sign the certificate with the CA we created (it's called self signing) - server.crt
openssl x509 -req -passin pass:1111 -days 3650 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt 

# Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem


// #Note:-
Note:- 
For changing the proto buffer file name use:
option go_package = "example_newpb";
package example;

// #Some Error which you might face 
1. Make file error:
cat -e -t -v  makefile
https://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date
https://stackoverflow.com/questions/834748/gcc-makefile-error-no-rule-to-make-target


// #Refrence:- 
https://flaviocopes.com/golang-environment-variables/


// #Error handling
https://grpc.io/docs/guides/error.html
http://avi.im/grpc-errors/


// #gRPCvsREST || HTTP1.1 vs HTTP2
https://husobee.github.io/golang/rest/grpc/2016/05/28/golang-rest-v-grpc.html
https://imagekit.io/demo/http2-vs-http1

//#.git file remove from sub directory if (Not for others)
//#TODO: Remove
find . -name ".git"
rm -rf .git
rm -rf .git*
ls -lah
find . -name ".git"


//#ssl-sercurity 
https://grpc.io/docs/guides/auth.html