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

type Users struct {
	Registertime int64 `json:"registertime"`

	Userid string `json:"userid"`

	Regionid string `json:"regionid"`

	Gender string `json:"gender"`
}

const UsersAvroCRC64Fingerprint = "z\x17(\xeac\xe6į"

func NewUsers() Users {
	r := Users{}
	return r
}

func DeserializeUsers(r io.Reader) (Users, error) {
	t := NewUsers()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeUsersFromSchema(r io.Reader, schema string) (Users, error) {
	t := NewUsers()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeUsers(r Users, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.Registertime, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Userid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Regionid, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Gender, w)
	if err != nil {
		return err
	}
	return err
}

func (r Users) Serialize(w io.Writer) error {
	return writeUsers(r, w)
}

func (r Users) Schema() string {
	return "{\"fields\":[{\"name\":\"registertime\",\"type\":{\"arg.properties\":{\"range\":{\"max\":1519273364600,\"min\":1487715775521}},\"type\":\"long\"}},{\"name\":\"userid\",\"type\":{\"arg.properties\":{\"regex\":\"User_[1-9]\"},\"type\":\"string\"}},{\"name\":\"regionid\",\"type\":{\"arg.properties\":{\"regex\":\"Region_[1-9]\"},\"type\":\"string\"}},{\"name\":\"gender\",\"type\":{\"arg.properties\":{\"options\":[\"MALE\",\"FEMALE\",\"OTHER\"]},\"type\":\"string\"}}],\"name\":\"ksql.users\",\"type\":\"record\"}"
}

func (r Users) SchemaName() string {
	return "ksql.users"
}

func (_ Users) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Users) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Users) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Users) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Users) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Users) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Users) SetString(v string)   { panic("Unsupported operation") }
func (_ Users) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Users) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.Registertime}

		return w

	case 1:
		w := types.String{Target: &r.Userid}

		return w

	case 2:
		w := types.String{Target: &r.Regionid}

		return w

	case 3:
		w := types.String{Target: &r.Gender}

		return w

	}
	panic("Unknown field index")
}

func (r *Users) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Users) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Users) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Users) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Users) HintSize(int)                     { panic("Unsupported operation") }
func (_ Users) Finalize()                        {}

func (_ Users) AvroCRC64Fingerprint() []byte {
	return []byte(UsersAvroCRC64Fingerprint)
}

func (r Users) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["registertime"], err = json.Marshal(r.Registertime)
	if err != nil {
		return nil, err
	}
	output["userid"], err = json.Marshal(r.Userid)
	if err != nil {
		return nil, err
	}
	output["regionid"], err = json.Marshal(r.Regionid)
	if err != nil {
		return nil, err
	}
	output["gender"], err = json.Marshal(r.Gender)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Users) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["registertime"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Registertime); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for registertime")
	}
	val = func() json.RawMessage {
		if v, ok := fields["userid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Userid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for userid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["regionid"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Regionid); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for regionid")
	}
	val = func() json.RawMessage {
		if v, ok := fields["gender"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Gender); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for gender")
	}
	return nil
}
