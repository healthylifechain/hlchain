package plugin

import (
	_ "github.com/healthylifechain/hlchain/plugin/consensus/init" //consensus init
	_ "github.com/healthylifechain/hlchain/plugin/crypto/init"    //crypto init
	_ "github.com/healthylifechain/hlchain/plugin/dapp/init"      //dapp init
	_ "github.com/healthylifechain/hlchain/plugin/mempool/init"   //mempool init
	_ "github.com/healthylifechain/hlchain/plugin/store/init"     //store init
)
