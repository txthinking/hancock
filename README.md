# hancock

[🗣 News](https://t.me/txthinking_news)
[💬 Join](https://join.txthinking.com)
[🩸 Youtube](https://www.youtube.com/txthinking) 
[❤️ Sponsor](https://github.com/sponsors/txthinking)

<p align="center"><img src="hancock.jpeg" width="200"></p>

Manage multiple remote servers and execute commands remotely
> 管理多个远程服务器并远程执行命令

❤️ A project by [txthinking.com](https://www.txthinking.com)

### Install via [nami](https://github.com/txthinking/nami)

```
nami install hancock
```

### Usage

```
Note:

    When adding an instance, the user must be allowed to execute sudo without a password!!!
    nami and joker are automatically installed when you run command for the first time on instance.

Add instance

    $ hancock add --name mylinux --server 1.2.3.4:22 --user root --password mypassword
    $ hancock add --name mylinux --server 1.2.3.4:22 --user root --key ./path/to/mykey.pem
    $ hancock add -n mylinux -s 1.2.3.4:22 -u root -p mypassword
    $ hancock add -n mylinux -s 1.2.3.4:22 -u root -k ./path/to/mykey.pem

List instances

    $ hancock list

Remove instance

    $ hancock mylinux remove


Run nami on an instance

    $ hancock mylinux nami install brook
    $ hancock mylinux nami list

Run joker on an instance

    $ hancock mylinux joker brook server --listen :9999 --password hello
    $ hancock mylinux joker last
    $ hancock mylinux joker list
    $ hancock mylinux joker stop 1234
    $ hancock mylinux joker log 1234

Run command and wait output on an instance

    $ hancock mylinux echo hello
    $ hancock mylinux sleep 3 '&&' echo hello

Start command and do not wait output on an instance

    $ hancock mylinux start echo hello
    $ hancock mylinux start sleep 3 '&&' echo hello

Upload your own command

    $ hancock mylinux upload ./path/to/command
```

## License

Licensed under The GPLv3 License
