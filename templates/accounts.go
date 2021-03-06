/*
 * Flow Go SDK
 *
 * Copyright 2019-2020 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package templates

import (
	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"

	"github.com/portto/blocto-flow-go-sdk"
)

const createAccountTemplate = `
transaction(publicKeys: [[UInt8]], code: [UInt8]) {
  prepare(signer: AuthAccount) {
	let acct = AuthAccount(payer: signer)

	for key in publicKeys {
	  acct.addPublicKey(key)
	}
	
	if code.length > 0 {
	  acct.setCode(code)
	}
  }
}
`

// CreateAccount generates a transactions that creates a new account.
//
// This template accepts a list of public keys and a code argument, both of which are optional.
//
// The final argument is the address of the account that will pay the account creation fee.
// This account is added as a transaction authorizer and therefore must sign the resulting transaction.
func CreateAccount(accountKeys []*flow.AccountKey, code []byte, payer flow.Address) *flow.Transaction {
	publicKeys := make([]cadence.Value, len(accountKeys))

	for i, accountKey := range accountKeys {
		publicKeys[i] = bytesToCadenceArray(accountKey.Encode())
	}

	cadencePublicKeys := cadence.NewArray(publicKeys)
	cadenceCode := bytesToCadenceArray(code)

	return flow.NewTransaction().
		SetScript([]byte(createAccountTemplate)).
		AddAuthorizer(payer).
		AddRawArgument(jsoncdc.MustEncode(cadencePublicKeys)).
		AddRawArgument(jsoncdc.MustEncode(cadenceCode))
}

const createAccountWithoutCodeTemplate = `
transaction(publicKeys: [[UInt8]], code: [UInt8]) {
  prepare(signer: AuthAccount) {
	let acct = AuthAccount(payer: signer)

	for key in publicKeys {
		acct.addPublicKey(key)
	}
  }
}
`

// CreateAccountWithoutCode generates a transactions that creates a new account without code.
//
// This template accepts a list of public keys and a code argument, both of which are optional.
//
// The final argument is the address of the account that will pay the account creation fee.
// This account is added as a transaction authorizer and therefore must sign the resulting transaction.
func CreateAccountWithoutCode(accountKeys []*flow.AccountKey, payer flow.Address) *flow.Transaction {
	publicKeys := make([]cadence.Value, len(accountKeys))

	for i, accountKey := range accountKeys {
		publicKeys[i] = bytesToCadenceArray(accountKey.Encode())
	}

	cadencePublicKeys := cadence.NewArray(publicKeys)
	cadenceCode := bytesToCadenceArray(nil)

	return flow.NewTransaction().
		SetScript([]byte(createAccountWithoutCodeTemplate)).
		AddAuthorizer(payer).
		AddRawArgument(jsoncdc.MustEncode(cadencePublicKeys)).
		AddRawArgument(jsoncdc.MustEncode(cadenceCode))
}

const updateAccountCodeTemplate = `
transaction(code: [UInt8]) {
  prepare(signer: AuthAccount) {
	signer.setCode(code)
  }
}
`

// UpdateAccountCode generates a transaction that updates the code deployed at an account.
func UpdateAccountCode(address flow.Address, code []byte) *flow.Transaction {
	cadenceCode := bytesToCadenceArray(code)

	return flow.NewTransaction().
		SetScript([]byte(updateAccountCodeTemplate)).
		AddRawArgument(jsoncdc.MustEncode(cadenceCode)).
		AddAuthorizer(address)
}

const addAccountKeyTemplate = `
transaction(publicKey: [UInt8]) {
  prepare(signer: AuthAccount) {
	signer.addPublicKey(publicKey)
  }
}
`

// AddAccountKey generates a transaction that adds a public key to an account.
func AddAccountKey(address flow.Address, accountKey *flow.AccountKey) *flow.Transaction {
	cadencePublicKey := bytesToCadenceArray(accountKey.Encode())

	return flow.NewTransaction().
		SetScript([]byte(addAccountKeyTemplate)).
		AddRawArgument(jsoncdc.MustEncode(cadencePublicKey)).
		AddAuthorizer(address)
}

const removeAccountKeyTemplate = `
transaction(keyIndex: Int) {
  prepare(signer: AuthAccount) {
    signer.removePublicKey(keyIndex)
  }	
}
`

// RemoveAccountKey generates a transaction that removes a key from an account.
func RemoveAccountKey(address flow.Address, keyIndex int) *flow.Transaction {
	cadenceKeyIndex := cadence.NewInt(keyIndex)

	return flow.NewTransaction().
		SetScript([]byte(removeAccountKeyTemplate)).
		AddRawArgument(jsoncdc.MustEncode(cadenceKeyIndex)).
		AddAuthorizer(address)
}

const replaceAccountKeysTemplate = `
transaction(publicKeys: [[UInt8]], keyIDs: [Int]) {
  prepare(signer: AuthAccount) {
	for id in keyIDs {
		signer.removePublicKey(id)
	}

	for key in publicKeys {
		signer.addPublicKey(key)
	}
  }
}
`

// ReplaceAccountKeys remove keys by ids and add new keys
func ReplaceAccountKeys(address flow.Address, ids []int, accountKeys []*flow.AccountKey) *flow.Transaction {
	publicKeys := make([]cadence.Value, len(accountKeys))
	for i, accountKey := range accountKeys {
		publicKeys[i] = bytesToCadenceArray(accountKey.Encode())
	}

	removeIDs := make([]cadence.Value, len(ids))
	for i, id := range ids {
		removeIDs[i] = cadence.NewInt(id)
	}

	cadencePublicKeys := cadence.NewArray(publicKeys)
	cadenceIDs := cadence.NewArray(removeIDs)

	return flow.NewTransaction().
		SetScript([]byte(replaceAccountKeysTemplate)).
		AddRawArgument(jsoncdc.MustEncode(cadencePublicKeys)).
		AddRawArgument(jsoncdc.MustEncode(cadenceIDs)).
		AddAuthorizer(address)
}

func bytesToCadenceArray(b []byte) cadence.Array {
	values := make([]cadence.Value, len(b))

	for i, v := range b {
		values[i] = cadence.NewUInt8(v)
	}

	return cadence.NewArray(values)
}
