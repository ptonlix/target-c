/*
 * SPDX-License-Identifier: Apache-2.0
 */

package ledgerapi

import (
	"fmt"
	"reflect"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// StateListInterface functions that a state list
// should have
type StateListInterface interface {
	AddState(StateInterface) error
	GetState(string, StateInterface) error
	UpdateState(StateInterface) error
	GetStateByRange(StateInterface, string, string) ([]StateInterface, error)
	GetStateByPartialCompositeKey(StateInterface, ...string) ([]StateInterface, error)
}

// StateList useful for managing putting data in and out
// of the ledger. Implementation of StateListInterface
type StateList struct {
	Ctx         contractapi.TransactionContextInterface
	Name        string
	Deserialize func([]byte, StateInterface) error
}

// AddState puts state into world state
func (sl *StateList) AddState(state StateInterface) error {
	key, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, state.GetSplitKey())
	data, err := state.Serialize()

	if err != nil {
		return err
	}

	return sl.Ctx.GetStub().PutState(key, data)
}

// GetState returns state from world state. Unmarshalls the JSON
// into passed state. Key is the split key value used in Add/Update
// joined using a colon
func (sl *StateList) GetState(key string, state StateInterface) error {
	ledgerKey, _ := sl.Ctx.GetStub().CreateCompositeKey(sl.Name, SplitKey(key))
	data, err := sl.Ctx.GetStub().GetState(ledgerKey)

	if err != nil {
		return err
	} else if data == nil {
		return fmt.Errorf("No state found for %s", key)
	}

	return sl.Deserialize(data, state)
}

// UpdateState puts state into world state. Same as AddState but
// separate as semantically different
func (sl *StateList) UpdateState(state StateInterface) error {
	return sl.AddState(state)
}

func (sl *StateList) GetStateByRange(state StateInterface, startKey, endKey string) ([]StateInterface, error) {
	resultsIterator, err := sl.Ctx.GetStub().GetStateByRange(startKey, endKey) //返回包含给出颜色的组合键的迭代器

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	t := reflect.TypeOf(state)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
	}

	list := []StateInterface{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		newStruc := reflect.New(t)
		err = sl.Deserialize(queryResponse.Value, newStruc.Interface().(StateInterface))
		if err != nil {
			return nil, err
		}

		list = append(list, newStruc.Interface().(StateInterface))
	}

	return list, nil
}

func (sl *StateList) GetStateByPartialCompositeKey(state StateInterface, key ...string) ([]StateInterface, error) {
	resultsIterator, err := sl.Ctx.GetStub().GetStateByPartialCompositeKey(sl.Name, key) //返回包含给出颜色的组合键的迭代器

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	t := reflect.TypeOf(state)
	if t.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		t = t.Elem()
	}

	list := []StateInterface{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		newStruc := reflect.New(t)
		err = sl.Deserialize(queryResponse.Value, newStruc.Interface().(StateInterface))
		if err != nil {
			return nil, err
		}

		list = append(list, newStruc.Interface().(StateInterface))
	}

	return list, nil

}
