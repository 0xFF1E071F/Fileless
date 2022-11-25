# Fileless Malware

## How does the fileless part work?
In order to make the malware work fileless, I used a python script that utilizes the memfd_create syscall in Linux to create an anonymous file in memory. That file has the exact same properties as a normal file, but lives in RAM. Therefore, we can retrieve the bytes of the malware from a webserver and write it straight to memory.

## How does the malware communicate?
The malware will communicate over WebSockets
