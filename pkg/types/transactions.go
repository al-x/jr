// Code generated by github.com/actgardner/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCES:
 *     NetDevice.avsc
 *     User.avsc
 *     campaign_finance.avsc
 *     clickstream.avsc
 *     clickstream_codes.avsc
 *     clickstream_users.avsc
 *     credit_cards.avsc
 *     device_information.avsc
 *     fleet_mgmt_description.avsc
 *     fleet_mgmt_location.avsc
 *     fleet_mgmt_sensors.avsc
 *     gaming_games.avsc
 *     gaming_player_activity.avsc
 *     gaming_players.avsc
 *     insurance_customer_activity.avsc
 *     insurance_customers.avsc
 *     insurance_offers.avsc
 *     inventory.avsc
 *     map_dumb_schema.avsc
 *     orders.avsc
 *     pageviews.avsc
 *     payroll_bonus.avsc
 *     payroll_employee.avsc
 *     payroll_employee_location.avsc
 *     pizza_orders.avsc
 *     pizza_orders_cancelled.avsc
 *     pizza_orders_completed.avsc
 *     product.avsc
 *     purchase.avsc
 *     ratings.avsc
 *     shoe_clickstream.avsc
 *     shoe_customers.avsc
 *     shoe_orders.avsc
 *     shoes_product.avsc
 *     siem_logs.avsc
 *     stockTrades.avsc
 *     stores.avsc
 *     syslog_logs.avsc
 *     transactions.avsc
 *     users.avsc
 */
package types

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v10/compiler"
	"github.com/actgardner/gogen-avro/v10/vm"
	"github.com/actgardner/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type Transactions struct {
	Transaction_id int64 `json:"transaction_id"`

	Card_id int64 `json:"card_id"`

	User_id string `json:"user_id"`

	Purchase_id int64 `json:"purchase_id"`

	Store_id int32 `json:"store_id"`
}

const TransactionsAvroCRC64Fingerprint = "~\xd9vZ\x8a\xc2\xe7\xcc"

func NewTransactions() Transactions {
	r := Transactions{}
	return r
}

func DeserializeTransactions(r io.Reader) (Transactions, error) {
	t := NewTransactions()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeTransactionsFromSchema(r io.Reader, schema string) (Transactions, error) {
	t := NewTransactions()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeTransactions(r Transactions, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Transaction_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Card_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.User_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteLong(r.Purchase_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Store_id, w)
	if err != nil {
		return err
	}
	return err
}

func (r Transactions) Serialize(w io.Writer) error {
	return writeTransactions(r, w)
}

func (r Transactions) Schema() string {
	return "{\"fields\":[{\"name\":\"transaction_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":1}},\"type\":\"long\"}},{\"name\":\"card_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":24,\"min\":1}},\"type\":\"long\"}},{\"name\":\"user_id\",\"type\":{\"arg.properties\":{\"regex\":\"User_[1-9]{0,1}\"},\"type\":\"string\"}},{\"name\":\"purchase_id\",\"type\":{\"arg.properties\":{\"iteration\":{\"start\":0}},\"type\":\"long\"}},{\"name\":\"store_id\",\"type\":{\"arg.properties\":{\"range\":{\"max\":7,\"min\":1}},\"type\":\"int\"}}],\"name\":\"datagen.example.transactions\",\"type\":\"record\"}"
}

func (r Transactions) SchemaName() string {
	return "datagen.example.transactions"
}

func (_ Transactions) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Transactions) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Transactions) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Transactions) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Transactions) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Transactions) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Transactions) SetString(v string)   { panic("Unsupported operation") }
func (_ Transactions) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Transactions) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Transaction_id}

		return w

	case 1:
		w := types.Long{Target: &r.Card_id}

		return w

	case 2:
		w := types.String{Target: &r.User_id}

		return w

	case 3:
		w := types.Long{Target: &r.Purchase_id}

		return w

	case 4:
		w := types.Int{Target: &r.Store_id}

		return w

	}
	panic("Unknown field index")
}

func (r *Transactions) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Transactions) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Transactions) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Transactions) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Transactions) HintSize(int)                     { panic("Unsupported operation") }
func (_ Transactions) Finalize()                        {}

func (_ Transactions) AvroCRC64Fingerprint() []byte {
	return []byte(TransactionsAvroCRC64Fingerprint)
}

func (r Transactions) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["transaction_id"], err = json.Marshal(r.Transaction_id)
	if err != nil {
		return nil, err
	}
	output["card_id"], err = json.Marshal(r.Card_id)
	if err != nil {
		return nil, err
	}
	output["user_id"], err = json.Marshal(r.User_id)
	if err != nil {
		return nil, err
	}
	output["purchase_id"], err = json.Marshal(r.Purchase_id)
	if err != nil {
		return nil, err
	}
	output["store_id"], err = json.Marshal(r.Store_id)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Transactions) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["transaction_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Transaction_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for transaction_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["card_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Card_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for card_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["user_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.User_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for user_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["purchase_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Purchase_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for purchase_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["store_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Store_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for store_id")
	}
	return nil
}
