package main

import (
	"fmt"
	"os"

	block "github.com/ntswamp/proof-of-kill/blc"

	"github.com/ntswamp/proof-of-kill/cli"
	"github.com/ntswamp/proof-of-kill/database"
	"github.com/ntswamp/proof-of-kill/network"

	log "github.com/corgi-kx/logcustom"
	"github.com/spf13/viper"
)

//initialize the system with values read from config.yaml
func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	logPath := viper.GetString("blockchain.log_path")
	listenHost := viper.GetString("network.listen_host")
	listenPort := viper.GetString("network.listen_port")
	rendezvousString := viper.GetString("network.rendezvous_string")
	protocolID := viper.GetString("network.protocol_id")
	tokenRewardNum := viper.GetInt("blockchain.token_reward_num")
	tradePoolLength := viper.GetInt("blockchain.trade_pool_length")
	mineDifficultyValue := viper.GetUint64("blockchain.mine_difficulty_value")
	verifyBit := viper.GetUint64("blockchain.verify_bit")
	chineseMnwordPath := viper.GetString("blockchain.chinese_mnemonic_path")

	network.TradePoolLength = tradePoolLength
	network.ListenHost = listenHost
	network.RendezvousString = rendezvousString
	network.ProtocolID = protocolID
	network.ListenPort = listenPort
	database.ListenPort = listenPort
	block.LISTEN_PORT = listenPort
	block.REWARD_AMOUNT = tokenRewardNum
	block.ROUND_BIT = mineDifficultyValue
	block.VERIFY_BIT = verifyBit
	block.MNWORD_PATH = chineseMnwordPath

	//set up logs
	file, err := os.OpenFile(fmt.Sprintf("%slog%s.txt", logPath, listenPort), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Error(err)
	}
	log.SetOutputAll(file)
}

func main() {
	//run the cli
	c := cli.New()
	c.Run()
}
