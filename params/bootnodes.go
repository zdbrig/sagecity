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
	"enode://f32f6097c5cc3881b7ba4759dd82b1c5c0edc443a446886b820ccf90f0f1e16471c45b24f20d82cc3779965344c824fe000f0b3da9f983d403b07d1c4ec55df2@45.76.34.21:60888", // Amsterdam
	"enode://35f02e35eaccc114a139f08df1d1a52ea887c32804b4feefc521cd109fa81c9a17c77fadf48eba09f3439144287ffc9e2ac314a13fc389a05bbf72efc545c96a@45.63.57.21:60888",  // Los Angeles
	"enode://441dcda34c25112e17cccac63c713a2140f4a62499ce1771463abadbce385127d37d7a052e15cce912536f017204d33366cd7618b787990c069c1a358e0bebae@104.156.227.21:60888", // New Jersey
	"enode://16d6ac25f86960b2999a658250780a9daf3e492e163a81495111f9ce55116acbca92e7e93e1c3a957149b04556c2cdab30ac5204955d834317c25b302087fed0@45.77.131.29:60888", // Tokoyo
	"enode://5aad6ecfc040415403f2f694d469273005db2ffa5022b07566205a3e42995eb9d5c80aebebb1a7801d091a1e12de73f35b386f7edb433fe25e89a8590fbf4528@107.191.57.205:60888",  // Sydney
	"enode://4ee45e009766c99e7fa91dd21332d8efe4458e38a4b1e65dcb07cbc1e937a4f53b12f39277df836fe4b056263a5105f279b7ac73af5e7ca544519e9c1aaff91a@207.180.198.242:60888",  // Germany



	// Ethereum Foundation Cpp Bootnodes
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{

}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{

}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{

}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}