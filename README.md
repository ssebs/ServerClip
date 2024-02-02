# serverclip

Tool for copying contents of a file over various ssh connections. Need to get a log file over to your laptop? This can help.


## How it works
- Server listening for data on laptop
- When sshing in, add a port forward
  - How to handle double port fwd?
    - laptop => bastion host => VM
- Client can send data via HTTP to server on laptop
  - CURL command
    - printed out for convenience by server on laptop

> TODO: Generate port # that's free instead of just CLI

### laptop => VM
```
laptop $ ssh -L 5000:127.0.0.1:5000 user@VM
```

### laptop => bastion => VM
```
laptop $ ssh -L 5000:127.0.0.1:5001 user@bastion
bastion $ ssh -L 5001:127.0.0.1:5002 user@VM
```

## Usage
- Download `ServerClip.exe` or `serverclip` (depending on OS)
- Run `ServerClip.exe`
  - You can choose the HTTP port via `-p` flag.
- Allow the server in your firewall (if it pops up)
- SSH with a port forward to your server(s)
- Once you're ready to upload a file, run the curl command
  - `curl -XPOST <laptop>:5000/upload -d @./<file>`
    - laptop needs to be the IP of your laptop

### Example:
- Laptop
```
PS C:\Users\ssebs\serverclip> .\ServerClip.exe -p 5000
ServerClip - 10.0.0.2:5000
Upload a file via this command:
$ curl -XPOST 10.0.0.2:5000/upload -d @/path/to/your_file

CTRL+C to exit.
Copied file to clipboard
[GIN] 2024/02/02 - 10:43:03 | 201 |       982.9µs |  10.0.0.3 | POST     "/upload"
PS C:\Users\ssebs\serverclip> 
```

- VM
```
[ssebs@bastion ~]$ echo "tmp" > tmp.txt
[ssebs@bastion ~]$ curl -XPOST 10.0.0.2:5000/upload -d @./tmp.txt 
Uploaded, check your clipboard.
[ssebs@bastion ~]$ 
```
- Laptop
  - Now if you paste, you'll paste the contents of that file


## Building
- Install golang
- `git clone github.com/ssebs/ServerClip`
- `cd ServerClip`
- Windows:
  - `go build -o ServerClip.exe .\cmd\main.go`
- Mac/Linux:
  - `go build -o serverclip ./cmd/main.go`
