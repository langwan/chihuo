echo '清理并生成目录'

OUT=../../tls
DAYS=365
RSALEN=2048
CN=chihuo

rm -rf ${OUT}/*
mkdir ${OUT}  >> /dev/null 2>&1

cd ${OUT}


echo '生成CA的私钥'
openssl genrsa -out ca.key ${RSALEN} >> /dev/null 2>&1

echo '生成CA的签名证书'
openssl req -new \
-x509 \
-key ca.key \
-subj "/CN=${CN}" \
-out ca.crt

echo ''
echo '生成server端私钥'
openssl genrsa -out server.key ${RSALEN} >> /dev/null 2>&1

echo '生成server端自签名'
openssl req -new \
-key server.key \
-subj "/CN=${CN}" \
-out server.csr

echo '签发server端证书'
openssl x509 -req  -sha256 \
-in server.csr \
-CA ca.crt -CAkey ca.key -CAcreateserial \
-out server.crt -text >> /dev/null 2>&1

echo '删除server端自签名证书'
rm server.csr

echo ''
echo '生成client私钥'
openssl genrsa -out client.key ${RSALEN}  >> /dev/null 2>&1

echo '生成client自签名'
openssl req -new \
    -subj "/CN=${CN}" \
    -key client.key \
    -out client.csr
echo '签发client证书'
openssl x509 -req -sha256\
 -CA ca.crt -CAkey ca.key -CAcreateserial\
 -days  ${DAYS}\
 -in client.csr\
 -out client.crt\
 -text >> /dev/null 2>&1
echo '删除client端自签名'
rm client.csr

echo ''
echo '删除临时文件'
rm ca.srl

echo ''
echo '完成'