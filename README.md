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

> Generate port # that's free...

### laptop => VM
```
laptop $ ssh -L 5000:127.0.0.1:5000 user@VM
```

### laptop => bastion => VM
```
laptop $ ssh -L 5000:127.0.0.1:5001 user@bastion
bastion $ ssh -L 5001:127.0.0.1:5002 user@VM
```

