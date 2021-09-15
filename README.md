# proof-of-kill


### testing program：

**1.Launch the Node, Create Wallets, Generate a Genesis Block**

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

**5.Synchronize**

Launch Node 2, Node 3 with #listen_port# field in the config.yaml set up to 7001,7002.</br>
Look closer to Node 1's log you will notice that other nodes are detected.
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118145703154.png)Node 2, Node 3 will synchronize local chain with Node 1 automatically once they get fully launched.
![在这里插入图片描述](https://img-blog.csdnimg.cn/20191118145752942.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3FxXzM1OTExMTg0,size_16,color_FFFFFF,t_70)

<br>

**6.Making Transfer**

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

**7.Checking Balance**

Node 1 got 100 coins at the beginning, but after transferrd 60 to Node 2 and Node 3, only 40 coins left for Node 1 now.
Node 2 received 30 from Node 1.
Node 3 received 30 from Node 1, and mined the block. therefore Node 3 received 10 coins as reward. It holds 40 coins.

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