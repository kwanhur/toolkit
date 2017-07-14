#!/usr/bin/env bash

STORE_DIR=`pwd`
# CA私钥
openssl genrsa -des3 -out ca.key 2048
#制作解密后的CA私钥
openssl rsa -in ca.key -out ca_decrypted.key
#ca.crt CA根证书（公钥）
openssl req -new -x509 -days 7305 -key ca.key -out ca.crt

DOMAIN=$1
#生成blog.creke.net证书私钥
openssl genrsa -des3 -out $DOMAIN.pem 1024
#制作解密后的blog.creke.net证书私钥
openssl rsa -in $DOMAIN.pem -out $DOMAIN.key
#生成签名请求
openssl req -new -key $DOMAIN.pem -out $DOMAIN.csr
#用CA进行签名
openssl ca -policy policy_anything -days 1460 -cert ca.crt -keyfile ca.key -in $DOMAIN.csr -out $DOMAIN.crt

echo "$STORE_DIR/$DOMAIN.crt"
echo "$STORE_DIR/$DOMAIN.key"