// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://ed534f765247cec5befd8c2f0c1dec2b46886c207fa213e54df4fee7c6add139383e67a510f301a7ad417a0ac5a124f9dc5ec6778950a984d96da62d8aa926fa@45.76.34.21:60777", // Amsterdam
	"enode://700a15fafff93d54340eaeeed2de319d07e2b7fd59fceb3c88da08e19eac29e17c082e49a3b5511d30025f7ef1995e34fd9659a85ffc834571f163466146e8cb@45.63.57.21:60777",  // Los Angeles
	"enode://e1a7fbcbc5939d872ccef39ab3db655ffe412fcce0de58e818bcbaf5e398a3cacf8e88a60c1a3d3d657c7d2ec9735050c759747bb7ddf66300b45729256709e7@104.156.227.21:60777", // New Jersey
	"enode://815367c396c886c167055f004f12f4c0e5df3073e39a8bcd3795a290691cd3cb2582011ec1e923f95be7ae8e2fea6c50cf9a181dfbe36907a7ada7947839970d@45.77.131.29:60777", // Tokoyo
	"enode://43c572300f10dcabb3248bde99fff76b2a70785674e5a8bdfe47f9a8b2b960a717b0cc8bd15a6e1e041110711e49e2d16f3b2aa83fdb82388909feb7e3040c31@107.191.57.205:60777",  // Sydney
	"enode://45c376c5d0946af07625640258bc062b6298afea8b9531b8ae3fe4477d4a4ff3933efdd0b82341ed89d84575d7bda6091c89704f5c6646331d2b6fd605c1b77c@207.180.198.242:60777",  // Germany



	// Ethereum Foundation Cpp Bootnodes
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://ed534f765247cec5befd8c2f0c1dec2b46886c207fa213e54df4fee7c6add139383e67a510f301a7ad417a0ac5a124f9dc5ec6778950a984d96da62d8aa926fa@45.76.34.21:60777", // Amsterdam
	"enode://700a15fafff93d54340eaeeed2de319d07e2b7fd59fceb3c88da08e19eac29e17c082e49a3b5511d30025f7ef1995e34fd9659a85ffc834571f163466146e8cb@45.63.57.21:60777",  // Los Angeles
	"enode://e1a7fbcbc5939d872ccef39ab3db655ffe412fcce0de58e818bcbaf5e398a3cacf8e88a60c1a3d3d657c7d2ec9735050c759747bb7ddf66300b45729256709e7@104.156.227.21:60777", // New Jersey
	"enode://815367c396c886c167055f004f12f4c0e5df3073e39a8bcd3795a290691cd3cb2582011ec1e923f95be7ae8e2fea6c50cf9a181dfbe36907a7ada7947839970d@45.77.131.29:60777", // Tokoyo
	"enode://43c572300f10dcabb3248bde99fff76b2a70785674e5a8bdfe47f9a8b2b960a717b0cc8bd15a6e1e041110711e49e2d16f3b2aa83fdb82388909feb7e3040c31@107.191.57.205:60777",  // Sydney
	"enode://45c376c5d0946af07625640258bc062b6298afea8b9531b8ae3fe4477d4a4ff3933efdd0b82341ed89d84575d7bda6091c89704f5c6646331d2b6fd605c1b77c@207.180.198.242:60777",  // Germany


}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://ed534f765247cec5befd8c2f0c1dec2b46886c207fa213e54df4fee7c6add139383e67a510f301a7ad417a0ac5a124f9dc5ec6778950a984d96da62d8aa926fa@45.76.34.21:60777", // Amsterdam
	"enode://700a15fafff93d54340eaeeed2de319d07e2b7fd59fceb3c88da08e19eac29e17c082e49a3b5511d30025f7ef1995e34fd9659a85ffc834571f163466146e8cb@45.63.57.21:60777",  // Los Angeles
	"enode://e1a7fbcbc5939d872ccef39ab3db655ffe412fcce0de58e818bcbaf5e398a3cacf8e88a60c1a3d3d657c7d2ec9735050c759747bb7ddf66300b45729256709e7@104.156.227.21:60777", // New Jersey
	"enode://815367c396c886c167055f004f12f4c0e5df3073e39a8bcd3795a290691cd3cb2582011ec1e923f95be7ae8e2fea6c50cf9a181dfbe36907a7ada7947839970d@45.77.131.29:60777", // Tokoyo
	"enode://43c572300f10dcabb3248bde99fff76b2a70785674e5a8bdfe47f9a8b2b960a717b0cc8bd15a6e1e041110711e49e2d16f3b2aa83fdb82388909feb7e3040c31@107.191.57.205:60777",  // Sydney
	"enode://45c376c5d0946af07625640258bc062b6298afea8b9531b8ae3fe4477d4a4ff3933efdd0b82341ed89d84575d7bda6091c89704f5c6646331d2b6fd605c1b77c@207.180.198.242:60777",  // Germany


}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://ed534f765247cec5befd8c2f0c1dec2b46886c207fa213e54df4fee7c6add139383e67a510f301a7ad417a0ac5a124f9dc5ec6778950a984d96da62d8aa926fa@45.76.34.21:60777", // Amsterdam
	"enode://700a15fafff93d54340eaeeed2de319d07e2b7fd59fceb3c88da08e19eac29e17c082e49a3b5511d30025f7ef1995e34fd9659a85ffc834571f163466146e8cb@45.63.57.21:60777",  // Los Angeles
	"enode://e1a7fbcbc5939d872ccef39ab3db655ffe412fcce0de58e818bcbaf5e398a3cacf8e88a60c1a3d3d657c7d2ec9735050c759747bb7ddf66300b45729256709e7@104.156.227.21:60777", // New Jersey
	"enode://815367c396c886c167055f004f12f4c0e5df3073e39a8bcd3795a290691cd3cb2582011ec1e923f95be7ae8e2fea6c50cf9a181dfbe36907a7ada7947839970d@45.77.131.29:60777", // Tokoyo
	"enode://43c572300f10dcabb3248bde99fff76b2a70785674e5a8bdfe47f9a8b2b960a717b0cc8bd15a6e1e041110711e49e2d16f3b2aa83fdb82388909feb7e3040c31@107.191.57.205:60777",  // Sydney
	"enode://45c376c5d0946af07625640258bc062b6298afea8b9531b8ae3fe4477d4a4ff3933efdd0b82341ed89d84575d7bda6091c89704f5c6646331d2b6fd605c1b77c@207.180.198.242:60777",  // Germany



}
