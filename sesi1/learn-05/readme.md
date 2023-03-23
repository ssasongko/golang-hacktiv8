# Channels
-  c <-- chan string is a receive-only channel
-  c chan<- string is a send-only channel
- channel direction adalah opsional

# Buffered vs Unbuffered
-  Channel bersifat unbuffered atau tidak di buffer di memori
-  Unbuffered channel yang dimana proses penerimaan dan pengiriman data bersifat synchronous.
-  Buffered channel yang dimana kita dapat menentukkan kapasitas dari buffer nya bersifat asynchronous
-  