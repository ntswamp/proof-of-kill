# proof-of-kill


### testing program：

This demo only supports mac/linux currently, install golang properly on the mentioned OS's for running this program:
https://golang.org/doc/install


**Build**


```shell
 go build -o app main.go
```
<br>


**Prepare Multiple Terminals**

You want to simulate multiple nodes of a P2P network by openning multiple terminal windows.
If you are using VScode, this goal can be easily achieved tapping `Split Terminal` button in the upper right-hand corner of default terminal window.
>If nodes can't find each other, turn off firewall and try again.

![Screenshot 2021-09-16 113610](https://user-images.githubusercontent.com/50705651/133540241-1bf10cb4-11fd-4457-aa42-92e427ada100.jpg)

<br>

**Play with Configuration File**
  
  The key field is `listen_port`, an unique ports stands for an unique node in a simulated P2P network.</br>
  You can leave other parts default. but don't set `mine_difficulty_value` lower than 24, otherwise you won't see the mining process show up in the log when it done too fast.
```shell
 vim config.yaml
```
```yaml
blockchain:
  # difficulty
  mine_difficulty_value: 24
  # mining reward
  token_reward_num: 10
  # start mining when this number is reached
  trade_pool_length: 2
  # log directory
  log_path: "./"
  # directory for mnemonicwords
  chinese_mnemonic_path: "./chinese_mnemonic_world.txt"
network:
  # local monitoring IP
  listen_host: "127.0.0.1"
  # local monitoring port
  listen_port: "9000"
  # unique identifier of node group(nodes can only discover each other in the same group)
  rendezvous_string: "pok"
  # nodes only send data to the nodes with the same protocol id.
  protocol_id: "/chain/1.1.0"

```

<br>

**Launch the Node, Create Wallets, Generate a Genesis Block**

Launch Node 1
```shell
 ./app
```

Create 3 wallet addresses
```
-> newwal
Mnemonic Word： ["吴昆","黔鳄","脾虚","牙挺","小膜","野驴","拆股"]
Private Key  ： 44NwhHw15MSebrVyNmg6m5jm9hGKxmXgjVeZbb7p5z7S
Address      ： 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd
-> newwal
Mnemonic Word： ["愈合","提额","盗汗","头型","专约","拒付","四创"]
Private Key  ： 3FZkWLFHNTGFd8MR2QikaN88nP6dmJeDRJkaM4XastN9
Address      ： 1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF
-> newwal
Mnemonic Word： ["胸水","榆蘑","范明","顺诚","沉香","无畏","肾盏"]
Private Key  ： 4k6iSU5sANDqXZQmQLkt7dbzpJAMJzDrxPGxRVHuKqqh
Address      ： 1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5
```

Generate the Genesis Block(fund 1st address 100 coins)
```
-> genesis -a 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd -v 100
Made Genesis Block.
```

Check out Node 1's Log by Command:(see mining process in detail.)
```shell
 tail -f log9000.txt 
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118144251486.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1OTExMTg0,size_16,color_FFFFFF,t_70)

<br>

**Synchronize**

Launch Node 2, Node 3 with #listen_port# field in the config.yaml set up to 7001,7002.</br>
Look closer to Node 1's log you will notice that other nodes are detected.
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118145703154.png)Node 2, Node 3 will synchronize local chain with Node 1 automatically once they get fully launched.
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118145752942.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1OTExMTg0,size_16,color_FFFFFF,t_70)

<br>

**Making Transfer**

Before transfer, you want to set the address to receive mining reward up, for each node.(if you haven't set it, nodes won't receive any rewards)</br>
Node 1:
```
-> setmineaddr -a 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd
Receiving Mining Reward On Address [1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd].
```
Node 2:
```
-> setmineaddr -a 1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF
Receiving Mining Reward On Address [1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF].
```
Node 3:
```
-> setmineaddr -a 1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5
Receiving Mining Reward On Address [1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5].
```

Transfer 30 coins to each node from Node 1:
```
-> send -from ["1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd","1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd"] -to ["1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF","1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5"] -amount [30,30]
transaction has been broadcast.
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/2019111815314125.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1OTExMTg0,size_16,color_FFFFFF,t_70)

<br>

**Checking Balance**

Node 1 got 100 coins at the beginning, but after transferrd 60 to Node 2 and Node 3, only 40 coins left for Node 1 now.<br>
Node 2 received 30 from Node 1.<br>
Node 3 received 30 from Node 1, and mined the block. therefore Node 3 received 10 coins as reward. It holds 40 coins.<br>

![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118153547470.png)type `bal` checking balance at 3 addresses
```
Node 1:
-> bal -a 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd
Address: 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd
Balance：40

Node 2:
-> bal -a 1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF
Address: 1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF
Balance：30

Node 3:
-> bal -a 1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5
Address: 1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5
Balance：40
```

<br>
