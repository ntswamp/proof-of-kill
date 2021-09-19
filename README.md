# proof-of-kill

Individual-friendly and low-power-consumption are rarely achieved at the same time in blockchain systems nowadays.
In this implementation I try to resolve this issues by creating a new type of consensus model named "PoK(Proof of Kill)", which uses a game-like frontend(to be called "Agent") instead of CPU power to hold a race between nodes with safety guaranteed.

![pok](https://user-images.githubusercontent.com/50705651/133920498-e6d197ab-75b8-456d-9a94-3da4d95714ad.jpg)



### testing program：

This demo only supports mac/linux currently, install golang properly on the mentioned OS's for running this program:
https://golang.org/doc/install



**Build**


```shell
 go build -o app
```
<br>



**Prepare Multiple Terminals**

You want to simulate multiple nodes of a P2P network by opening multiple terminal windows.
If you are using VScode, this goal can be easily achieved by tapping `Split Terminal` button in the upper right-hand corner of the default terminal window. I suggest an independent one though.
>If nodes can't find each other, turn off the firewall and retry.

![Screenshot 2021-09-16 113610](https://user-images.githubusercontent.com/50705651/133540241-1bf10cb4-11fd-4457-aa42-92e427ada100.jpg)

<br>



**Play with Configuration File**
  
  The key field in the configuration file is `listen_port`, a unique ports stands for a unique node in our simulated P2P network.</br>
  You can leave other parts default. but don't set `mine_difficulty_value` lower than 8, otherwise you won't see any mining details show up in the log when it was done too fast.
```shell
 vim config.yaml
```
```yaml
blockchain:
  # difficulty. don't make this lower than 10.
  mine_difficulty_value: 10
  # verifying bits upon PoK verification. higher bits bring higher security. but never make this value bigger than mine_difficulty_value
  verify_bit: 5
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
  listen_port: "6666"
  # unique identifier of node group(nodes can only discover each other in the same group)
  rendezvous_string: "pok"
  # nodes only send data to the nodes with the same protocol id.
  protocol_id: "/chain/1.1.0"

```

<br>



**Launch the Node, Hire an Agent**

Launch Node 1
```shell
 ./app
```

If you haven't hired any agents yet, bc you just downloaded the program for example, the program will prompt you to have one.
Just follow the instructions.

![agent](https://user-images.githubusercontent.com/50705651/133919000-5366243f-1c9e-45df-90be-9ee35bc2c584.jpg)



**Create Wallets, Generate the Genesis Block**


Once you've got an agent, you unlock more commands.
Create 3 wallet addresses using command `newwal`.
```
-> newwal
Mnemonic Word： ["吴昆","黔鳄","板桥","键盘","小膜","潘总","拆股"]
Private Key  ： 44NwhHw15MSebrVyNmg6m5jm9hGKxmXgjVeZbb7p5z7S
Address      ： 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd
-> newwal
Mnemonic Word： ["愈合","明月","计算","头型","专约","拒付","四创"]
Private Key  ： 3FZkWLFHNTGFd8MR2QikaN88nP6dmJeDRJkaM4XastN9
Address      ： 1KNgFa165mjG2dZLcZ2ifhKtaZLu3SR5iF
-> newwal
Mnemonic Word： ["几何","榆蘑","范明","顺诚","沉香","无畏","光露"]
Private Key  ： 4k6iSU5sANDqXZQmQLkt7dbzpJAMJzDrxPGxRVHuKqqh
Address      ： 1EVrFBakJnhaWAvHQNhCJKLzensYqtJxR5
```

Generate the Genesis Block(fund 1st address 100 coins)
```
-> gen -a 1Dx8UpokXuv7Bvqa5ocgXKv8PKRLnvjdsd -v 100
Made Genesis Block.
```

Open another terminal, Check out Node 1's log by command:(see mining process in detail.)
```shell
 tail -f log6666.txt 
```
![tail](https://user-images.githubusercontent.com/50705651/133919007-1c064936-454f-4920-8aaa-6d3d54e4ff6f.jpg)


<br>



**Synchronize**

Launch Node 2, Node 3 with #listen_port# field in the config.yaml set up to 6667,6668.</br>
Look closer to Node 1's log you will notice that other nodes are detected.
![image](https://user-images.githubusercontent.com/50705651/133558870-490772fe-b1a9-4440-8369-07ad64a3d4d3.png)Node 2, Node 3 will synchronize local chain with Node 1 automatically once they get fully launched.
![updateothers](https://user-images.githubusercontent.com/50705651/133918709-cbe2991c-1902-40a8-865d-7d860d61e089.jpg)


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
![transaction broadcast](https://user-images.githubusercontent.com/50705651/133918734-326db616-4b2b-40aa-b0ee-82ae70452127.jpg)


Once your transaction gets broadcast to the network, nodes which received it will start mining immediately, including yourself.
But if any other nodes successfully mined the block with the same height as yours, your mining will be aborted.
![lost](https://user-images.githubusercontent.com/50705651/133919394-b55cc05d-e476-48d0-a08c-977e6a55139f.jpg)

<br>



**Checking Balance**

Node 1 got 100 coins at the beginning, but after transferring 60 to Node 2 and 3, only 40 coins are left for Node 1 now.<br>
Node 2 received 30 from Node 1.<br>
Node 3 received 30 from Node 1, and mined the block. therefore Node 3 received 10 coins as reward. It holds 40 coins.<br>

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

You can do more transfer then checking balance on both of 3 nodes to make sure the chain is working properly.

![comparing balance](https://user-images.githubusercontent.com/50705651/133918765-a99cc1cc-95a2-402e-b19a-26eb9562d65e.jpg)

<br>


**Print Blockchain**

Simply type cmd below to discover details about the current PoK chain.
```
-> chain
```
![chian](https://user-images.githubusercontent.com/50705651/133918788-0ed2c84d-4e8b-4794-bb8d-f11730cb509e.jpg)
