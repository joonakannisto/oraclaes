# oraclaes
AES UDP oracle

Send data in over UDP, and you get the first 16 bytes of your packet payload encrypted with the same key each time. (current implementation uses random key that changes if the service is restarted) 

Can be used as a TPMish network service for key derivation as an obstacle to offline bruteforce. 
