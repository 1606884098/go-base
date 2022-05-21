// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/conformal/btcchain"
	"github.com/conformal/btcdb"
	"github.com/conformal/btcscript"
	"github.com/conformal/btcwire"
	"github.com/conformal/seelog"
	"os"
	"time"
)

const (
	// lockTimeThreshold is the number below which a lock time is
	// interpreted to be a block number.  Since an average of one block
	// is generated per 10 minutes, this allows blocks for about 9,512
	// years.  However, if the field is interpreted as a timestamp, given
	// the lock time is a uint32, the max is sometime around 2106.
	lockTimeThreshold uint32 = 5e8 // Tue Nov 5 00:53:20 1985 UTC
)

var (
	log = seelog.Disabled
)

// logClosure is used to provide a closure over expensive logging operations
// so don't have to be performed when the logging level doesn't warrant it.
type logClosure func() string

// String invokes the underlying function and returns the result.
func (c logClosure) String() string {
	return c()
}

// newLogClosure returns a new closure over a function that returns a string
// which itself provides a Stringer interface so that it can be used with the
// logging system.
func newLogClosure(c func() string) logClosure {
	return logClosure(c)
}

// newLogger creates a new seelog logger using the provided logging level and
// log message prefix.
func newLogger(level string, prefix string) seelog.LoggerInterface {
	//<seelog type="adaptive" mininterval="2000000" maxinterval="100000000"
	//	critmsgcount="500" minlevel="%s">

	fmtstring := `
	<seelog type="sync" minlevel="%s">
		<outputs formatid="all">
			<console/>
		</outputs>
		<formats>
			<format id="all" format="%%Time %%Date [%%LEV] %s: %%Msg%%n" />
		</formats>
	</seelog>`
	config := fmt.Sprintf(fmtstring, level, prefix)

	logger, err := seelog.LoggerFromConfigAsString(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create logger: %v", err)
		os.Exit(1)
	}

	return logger
}

// useLogger sets the btcd logger to the passed logger.
func useLogger(logger seelog.LoggerInterface) {
	log = logger
}

// setLogLevel sets the log level for the logging system.  It initializes a
// logger for each subsystem at the provided level.
func setLogLevel(logLevel string) []seelog.LoggerInterface {
	var loggers []seelog.LoggerInterface

	// Define sub-systems.
	subSystems := []struct {
		level     string
		prefix    string
		useLogger func(seelog.LoggerInterface)
	}{
		{logLevel, "BTCD", useLogger},
		{logLevel, "BCDB", btcdb.UseLogger},
		{logLevel, "CHAN", btcchain.UseLogger},
		{logLevel, "SCRP", btcscript.UseLogger},
	}

	// Configure all sub-systems with new loggers while keeping track of
	// the created loggers to return so they can be flushed.
	for _, s := range subSystems {
		newLog := newLogger(s.level, s.prefix)
		loggers = append(loggers, newLog)
		s.useLogger(newLog)
	}

	return loggers
}

// directionString is a helper function that returns a string that represents
// the direction of a connection (inbound or outbound).
func directionString(inbound bool) string {
	if inbound {
		return "inbound"
	}
	return "outbound"
}

// formatLockTime returns a transaction lock time as a human-readable string.
func formatLockTime(lockTime uint32) string {
	// The lock time field of a transaction is either a block height at
	// which the transaction is finalized or a timestamp depending on if the
	// value is before the lockTimeThreshold.  When it is under the
	// threshold it is a block height.
	if lockTime < lockTimeThreshold {
		return fmt.Sprintf("height %d", lockTime)
	}

	return time.Unix(int64(lockTime), 0).String()
}

// invSummary returns an inventory messege as a human-readable string.
func invSummary(invList []*btcwire.InvVect) string {
	// No inventory.
	invLen := len(invList)
	if invLen == 0 {
		return "empty"
	}

	// One inventory item.
	if invLen == 1 {
		iv := invList[0]
		switch iv.Type {
		case btcwire.InvTypeError:
			return fmt.Sprintf("error %s", iv.Hash)
		case btcwire.InvTypeBlock:
			return fmt.Sprintf("block %s", iv.Hash)
		case btcwire.InvTypeTx:
			return fmt.Sprintf("tx %s", iv.Hash)
		}

		return fmt.Sprintf("unknown (%d) %s", uint32(iv.Type), iv.Hash)
	}

	// More than one inv item.
	return fmt.Sprintf("size %d", invLen)
}

// locatorSummary returns a block locator as a human-readable string.
func locatorSummary(locator []*btcwire.ShaHash, stopHash *btcwire.ShaHash) string {
	if len(locator) > 0 {
		return fmt.Sprintf("locator %s, stop %s", locator[0], stopHash)
	}

	return fmt.Sprintf("no locator, stop %s", stopHash)

}

// messageSummary returns a human-readable string which summarizes a message.
// Not all messages have or need a summary.  This is used for debug logging.
func messageSummary(msg btcwire.Message) string {
	switch msg := msg.(type) {
	case *btcwire.MsgVersion:
		return fmt.Sprintf("agent %s, pver %d, block %d",
			msg.UserAgent, msg.ProtocolVersion, msg.LastBlock)

	case *btcwire.MsgVerAck:
		// No summary.

	case *btcwire.MsgGetAddr:
		// No summary.

	case *btcwire.MsgAddr:
		return fmt.Sprintf("%d addr", len(msg.AddrList))

	case *btcwire.MsgPing:
		// No summary - perhaps add nonce.

	case *btcwire.MsgPong:
		// No summary - perhaps add nonce.

	case *btcwire.MsgAlert:
		// No summary.

	case *btcwire.MsgMemPool:
		// No summary.

	case *btcwire.MsgTx:
		hash, _ := msg.TxSha()
		return fmt.Sprintf("hash %s, %d inputs, %d outputs, lock %s",
			hash, len(msg.TxIn), len(msg.TxOut),
			formatLockTime(msg.LockTime))

	case *btcwire.MsgBlock:
		header := &msg.Header
		hash, _ := msg.BlockSha()
		return fmt.Sprintf("hash %s, ver %d, %d tx, %s", hash,
			header.Version, header.TxnCount, header.Timestamp)

	case *btcwire.MsgInv:
		return invSummary(msg.InvList)

	case *btcwire.MsgNotFound:
		return invSummary(msg.InvList)

	case *btcwire.MsgGetData:
		return invSummary(msg.InvList)

	case *btcwire.MsgGetBlocks:
		return locatorSummary(msg.BlockLocatorHashes, &msg.HashStop)

	case *btcwire.MsgGetHeaders:
		return locatorSummary(msg.BlockLocatorHashes, &msg.HashStop)

	case *btcwire.MsgHeaders:
		return fmt.Sprintf("num %d", len(msg.Headers))
	}

	// No summary for other messages.
	return ""
}
