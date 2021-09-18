package block

import (
	"math"
)

//当前节点发现的网络中最新区块高度
var NEWEST_BLOCK_HEIGHT int

//当前本地监听端口
var LISTEN_PORT string

//挖矿奖励代币数量
var REWARD_AMOUNT int

//挖矿难度值
var TARGET_BIT uint

var VERIFY_BIT uint

//中文助记词地址
var MNWORD_PATH string

//奖励地址在数据库中的键
const REWARD_ADDR_KEY = "rewardAddress"

//最新区块Hash在数据库中的键
const LATEST_BLOCK_HASH_KEY = "lastHash"

//钱包地址在数据库中的键
const ADDR_LIST_KEY = "addressList"

//公链版本信息默认为0
const VERSION = byte(0x00)

//两次sha256(公钥hash)后截取的字节数量
const CHECKSUM = 4

//随机数不能超过的最大值
const MAXINT = math.MaxInt64

//random seed used in genesis block
const GENESIS_SEED int64 = 8601066706715
