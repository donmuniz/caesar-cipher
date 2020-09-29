# Cryptography 101

## 🛰️ Check it Live!

### Decode

```
curl --request POST \
  --url http://infinite-bayou-68203.herokuapp.com/decode \
  --header 'content-type: application/json' \
  --data '{
	"msg": "UHCD KNOAOA PA MLG UYJX TOWBGR RF JWIIMSOBW VRFRII MA WHV WUQOITLY. ZJ UHG LYTV GVV APXMT LFZQZ CNKAF RN ZV, RQLW VAY RVRUKVPVDNJ IVWWFAX!"
}'
```

### Encode
```
curl --request POST \
  --url http://infinite-bayou-68203.herokuapp.com/encode \
  --header 'content-type: application/json' \
  --data '{
	"msg": "Look Ma, im online!"
}'
```

## Intro
We at Backend Architecture just discovered an alien civilization, 
and are orchestrating with them a world domination masterplan. 

To avoid being caught by the authorities, we developed a cypher 
that will encrypt / decrypt the communications with them.

## How does the cypher work
The cypher has a key, which we shared with the alien civilization. 
They key is used to encode the different characters of the message we want to send.

This is a version of the [CAESAR](https://en.wikipedia.org/wiki/Caesar_cipher) cypher, with the difference
 that each character is encoded using a different value in the cypher key. 

For the first  character of the message, the cypher will use the first character of the key. 
Then the second, and so on.  
  
## The Problem
Given the cypher key (`BACKENDARCHITECTURE`) and an encrypted message, 
please provide the unencrypted version of the message (ie: build a function to decrypt messages)

Encrypted Message:
`HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML`

#### Bonus
Now that you know how to decrypt lets build an encrypt function that does the opposite thing. 

Now send us an encrypted message!


### Notes
Everything is in uppercase (A to Z).  
Take a look at what a rune is, this could simplify your life for this exercise (or not :troll:)

```go
imarune := 'A'
imarunetoo := "string"[0]
iamstillarune := uint8(66) // Before it said int32(66) and this could be missleading
````





