# Welcome to my little Fileless malware project!

## About the stager
In order to launch our malware on the system without introducing a file, I decided to use the memfd_create method to create a file in memory, which shares the same properties as a normal file.

## About the agent
The future agent will be able to use a P2P network. I am not sure yet if this will be permanent or only used as a fallback when the C2 dies.
