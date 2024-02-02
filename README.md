# serverclip

Tool for copying contents of a file over various ssh connections. Need to get a log file over to your laptop? This can help.


## Usage
- Download `ServerClip.exe` or `serverclip` (depending on OS)
- Run `ServerClip.exe` from a terminal
  - You can choose the HTTP port via `-p` flag.
- Allow the server in your firewall (if it pops up)
- SSH with a port forward to your server(s)
  - See the output from the `ServerClip.exe` command, it will show what ssh command to run.
- Once you're ready to upload a file, run the curl command on the VM/host.
  - `curl -XPOST 127.0.0.1:5000 --data-binary @./<file>`
    - The port may be different.

### Example:
- Laptop
```
PS C:\Users\ssebs\serverclip> .\ServerClip.exe -p 5000
ServerClip - Listening on http://10.0.0.2:5000   (press CTRL+C to quit)
SSH to your host using:
$ ssh -R 5000:127.0.0.1:5000 <user>@<hostname>
To send a file to your clipboard, run this command:
$ curl -XPOST 127.0.0.1:5000 --data-binary @/path/to/your_file
Waiting for file...

Copied file to clipboard!
Waiting for file...(press CTRL+C to quit)
```

- VM
```
[ssebs@vm ~]$ echo "tmp" > tmp.txt
[ssebs@vm ~]$ curl -XPOST 127.0.0.1:5000 --data-binary @./tmp.txt 
Uploaded, check your clipboard.
[ssebs@vm ~]$ 
```
- Laptop
  - Now if you paste, you'll paste the contents of that file!


## Building
- Install golang
- `git clone github.com/ssebs/ServerClip`
- `cd ServerClip`
- Windows:
  - `go build -o ServerClip.exe .\cmd\main.go`
- Mac/Linux:
  - `go build -o serverclip ./cmd/main.go`


## How it works
- Server listening for data on laptop
- When sshing in, add a port forward
  - Supports multiple hops
    - laptop => bastion host => VM
    - The port you choose must be open on all hosts!
- Client can send data via HTTP to server on laptop
  - CURL command
    - printed out for convenience by server on laptop

## TODO
- [x] Fix missing newlines / special chars
- [ ] Build for more platforms
- [ ] Better usage


### laptop => VM
```
[laptop $] ssh -R 5000:127.0.0.1:5000 user@VM
...
[VM $] curl -XPOST 127.0.0.1:5000 --data-binary @/path/to/your_file
```

### laptop => bastion => VM
```
[laptop $] ssh -R 5000:127.0.0.1:5000 user@bastion
[bastion $] ssh -R 5000:127.0.0.1:5000 user@VM
[VM $] curl -XPOST 127.0.0.1:5000 --data-binary @/path/to/your_file
```

[LICENSE Apache 2.0](./LICENSE)
