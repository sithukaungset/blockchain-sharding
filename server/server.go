/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lynn9388/blockchain-sharding/common"
)

func StartServer() {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)

	//go newAPIService(&common.Server.APIAddr)
	startRPCServer()

	time.Sleep(2 * time.Second)

	//go p2p.NewNodeManager()
	//go p2p.NewPeerManager()

	select {
	case <-sigChan:
		stopRPCServer()
		common.Logger.Info("caught stop signal, quitting...")
	}
}
