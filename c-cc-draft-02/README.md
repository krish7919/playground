
### Relevant docs at:
1. wget https://raw.githubusercontent.com/rfcs/crypto-conditions/draft-02/src/asn1/CryptoConditions.asn
2. wget https://raw.githubusercontent.com/rfcs/crypto-conditions/draft-02/src/spec/crypto-conditions.md

### asn1c Compiler Installation
wget https://github.com/vlm/asn1c/releases/download/v0.9.27/asn1c-0.9.27.tar.gz
tar xzvf asn1c-0.9.27.tar.gz
cd asn1c-0.9.27
./configure && make && make check
make install

### Compile CryptoConditions.asn
asn1c -R  CryptoConditions.asn 

asn1c -fbless-SIZE -fincludes-quoted CryptoConditions.asn

asn1c -pdu=all -fincludes-quoted -fbless-SIZE CryptoConditions.asn

asn1c -R -pdu=all -fbless-SIZE CryptoConditions.asn


### Wrapper?
Statically link this to a wrapper?


vim converter-sample.c
make -f Makefile.am.sample

