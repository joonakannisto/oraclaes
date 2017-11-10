package main

import (
        "net"
        "fmt"
        "crypto/aes"
        "crypto/rand"
        "bytes"

)

func main() {
  bs := 16
  key := make([]byte, bs)
  _,err := rand.Read(key)
  ServerAddr, err := net.ResolveUDPAddr("udp", ":2789")
  if err != nil { panic(err)}
  ServerConn, err := net.ListenUDP("udp", ServerAddr)
  if err != nil { panic(err)}
  defer ServerConn.Close()

  for {
    buf := make([]byte, 16)
		_, addr, err := ServerConn.ReadFromUDP(buf)
    if err != nil {
			fmt.Println("error: ", err)
		}
    buf = ECB(key,buf)
    ServerConn.WriteTo(buf[0:bs], addr)

  }
}
func ECB(key, plaintext []byte) []byte {
  block, err := aes.NewCipher(key)
  if err != nil { panic(err)}
  bs := block.BlockSize();
  if len(plaintext) % bs != 0 {
    plaintext = Padding(plaintext,bs)
  }
  var ciphertext []byte
  cipherblock := make([]byte, bs)
  for len(plaintext) > 0 {
    block.Encrypt(cipherblock,plaintext[:bs])
    ciphertext = append(ciphertext,cipherblock[:]...)
    plaintext = plaintext[bs:]
  }
  return ciphertext
}
func Padding (array []byte, moniker int) []byte {
  if(len(array)%moniker!=0) {
    pad := []byte{byte(0x04)}
    padding:= bytes.Repeat(pad,moniker-len(array)%moniker)
    return append(array,padding[:]...)
  } else {
    return array
  }
}
