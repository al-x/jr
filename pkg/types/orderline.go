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

type Orderline struct {
	Product_id int32 `json:"product_id"`

	Category string `json:"category"`

	Quantity int32 `json:"quantity"`

	Unit_price float64 `json:"unit_price"`

	Net_price float64 `json:"net_price"`
}

const OrderlineAvroCRC64Fingerprint = "\xf6+\xf6qm\x17ʦ"

func NewOrderline() Orderline {
	r := Orderline{}
	return r
}

func DeserializeOrderline(r io.Reader) (Orderline, error) {
	t := NewOrderline()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeOrderlineFromSchema(r io.Reader, schema string) (Orderline, error) {
	t := NewOrderline()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeOrderline(r Orderline, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.Product_id, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Category, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.Quantity, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Unit_price, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.Net_price, w)
	if err != nil {
		return err
	}
	return err
}

func (r Orderline) Serialize(w io.Writer) error {
	return writeOrderline(r, w)
}

func (r Orderline) Schema() string {
	return "{\"arg.properties\":{\"options\":[{\"category\":\"calzone\",\"net_price\":75.12,\"product_id\":66,\"quantity\":4,\"unit_price\":18.78},{\"category\":\"salad\",\"net_price\":19.6,\"product_id\":71,\"quantity\":4,\"unit_price\":4.9},{\"category\":\"pizza\",\"net_price\":58.83,\"product_id\":79,\"quantity\":3,\"unit_price\":19.61},{\"category\":\"calzone\",\"net_price\":26.1,\"product_id\":38,\"quantity\":2,\"unit_price\":13.05},{\"category\":\"calzone\",\"net_price\":17.37,\"product_id\":26,\"quantity\":1,\"unit_price\":17.37},{\"category\":\"calzone\",\"net_price\":73.89,\"product_id\":66,\"quantity\":3,\"unit_price\":24.63},{\"category\":\"dessert\",\"net_price\":63.42,\"product_id\":62,\"quantity\":3,\"unit_price\":21.14},{\"category\":\"salad\",\"net_price\":63.39,\"product_id\":68,\"quantity\":3,\"unit_price\":21.13},{\"category\":\"salad\",\"net_price\":16.92,\"product_id\":72,\"quantity\":2,\"unit_price\":8.46},{\"category\":\"calzone\",\"net_price\":7.2,\"product_id\":55,\"quantity\":2,\"unit_price\":3.6},{\"category\":\"pizza\",\"net_price\":76.04,\"product_id\":77,\"quantity\":4,\"unit_price\":19.01},{\"category\":\"pizza\",\"net_price\":4.47,\"product_id\":60,\"quantity\":3,\"unit_price\":1.49},{\"category\":\"dessert\",\"net_price\":69.99,\"product_id\":8,\"quantity\":3,\"unit_price\":23.33},{\"category\":\"calzone\",\"net_price\":44.96,\"product_id\":91,\"quantity\":4,\"unit_price\":11.24},{\"category\":\"pizza\",\"net_price\":22.9,\"product_id\":78,\"quantity\":2,\"unit_price\":11.45},{\"category\":\"dessert\",\"net_price\":66.06,\"product_id\":27,\"quantity\":3,\"unit_price\":22.02},{\"category\":\"salad\",\"net_price\":9.88,\"product_id\":58,\"quantity\":1,\"unit_price\":9.88},{\"category\":\"dessert\",\"net_price\":98.76,\"product_id\":9,\"quantity\":4,\"unit_price\":24.69},{\"category\":\"salad\",\"net_price\":12.62,\"product_id\":77,\"quantity\":1,\"unit_price\":12.62},{\"category\":\"wings\",\"net_price\":6.63,\"product_id\":27,\"quantity\":3,\"unit_price\":2.21},{\"category\":\"calzone\",\"net_price\":19.04,\"product_id\":4,\"quantity\":1,\"unit_price\":19.04},{\"category\":\"calzone\",\"net_price\":16.38,\"product_id\":48,\"quantity\":1,\"unit_price\":16.38},{\"category\":\"dessert\",\"net_price\":46.2,\"product_id\":26,\"quantity\":5,\"unit_price\":9.24},{\"category\":\"pizza\",\"net_price\":80.45,\"product_id\":75,\"quantity\":5,\"unit_price\":16.09},{\"category\":\"salad\",\"net_price\":32.64,\"product_id\":33,\"quantity\":3,\"unit_price\":10.88},{\"category\":\"calzone\",\"net_price\":113.2,\"product_id\":45,\"quantity\":5,\"unit_price\":22.64},{\"category\":\"dessert\",\"net_price\":13.15,\"product_id\":65,\"quantity\":5,\"unit_price\":2.63},{\"category\":\"wings\",\"net_price\":5.68,\"product_id\":52,\"quantity\":1,\"unit_price\":5.68},{\"category\":\"pizza\",\"net_price\":16.47,\"product_id\":23,\"quantity\":3,\"unit_price\":5.49},{\"category\":\"dessert\",\"net_price\":3.02,\"product_id\":67,\"quantity\":1,\"unit_price\":3.02},{\"category\":\"dessert\",\"net_price\":66.76,\"product_id\":71,\"quantity\":4,\"unit_price\":16.69},{\"category\":\"salad\",\"net_price\":51.21,\"product_id\":69,\"quantity\":3,\"unit_price\":17.07},{\"category\":\"pizza\",\"net_price\":46.2,\"product_id\":41,\"quantity\":4,\"unit_price\":11.55},{\"category\":\"wings\",\"net_price\":13.18,\"product_id\":44,\"quantity\":2,\"unit_price\":6.59},{\"category\":\"wings\",\"net_price\":33.08,\"product_id\":66,\"quantity\":2,\"unit_price\":16.54},{\"category\":\"calzone\",\"net_price\":60.35,\"product_id\":8,\"quantity\":5,\"unit_price\":12.07},{\"category\":\"salad\",\"net_price\":69.5,\"product_id\":43,\"quantity\":5,\"unit_price\":13.9},{\"category\":\"calzone\",\"net_price\":101.5,\"product_id\":13,\"quantity\":5,\"unit_price\":20.3},{\"category\":\"salad\",\"net_price\":83.05,\"product_id\":65,\"quantity\":5,\"unit_price\":16.61},{\"category\":\"dessert\",\"net_price\":4.18,\"product_id\":48,\"quantity\":1,\"unit_price\":4.18},{\"category\":\"calzone\",\"net_price\":17.22,\"product_id\":59,\"quantity\":3,\"unit_price\":5.74},{\"category\":\"calzone\",\"net_price\":15.65,\"product_id\":42,\"quantity\":1,\"unit_price\":15.65},{\"category\":\"calzone\",\"net_price\":13.95,\"product_id\":22,\"quantity\":5,\"unit_price\":2.79},{\"category\":\"pizza\",\"net_price\":30.7,\"product_id\":47,\"quantity\":2,\"unit_price\":15.35},{\"category\":\"salad\",\"net_price\":124.85,\"product_id\":72,\"quantity\":5,\"unit_price\":24.97},{\"category\":\"salad\",\"net_price\":2.85,\"product_id\":60,\"quantity\":1,\"unit_price\":2.85},{\"category\":\"dessert\",\"net_price\":1.24,\"product_id\":58,\"quantity\":1,\"unit_price\":1.24},{\"category\":\"calzone\",\"net_price\":96,\"product_id\":88,\"quantity\":4,\"unit_price\":24},{\"category\":\"wings\",\"net_price\":6.66,\"product_id\":98,\"quantity\":1,\"unit_price\":6.66},{\"category\":\"pizza\",\"net_price\":43.86,\"product_id\":37,\"quantity\":3,\"unit_price\":14.62},{\"category\":\"calzone\",\"net_price\":42.63,\"product_id\":37,\"quantity\":3,\"unit_price\":14.21},{\"category\":\"wings\",\"net_price\":18.88,\"product_id\":54,\"quantity\":1,\"unit_price\":18.88},{\"category\":\"wings\",\"net_price\":28.41,\"product_id\":52,\"quantity\":3,\"unit_price\":9.47},{\"category\":\"wings\",\"net_price\":13.39,\"product_id\":80,\"quantity\":1,\"unit_price\":13.39},{\"category\":\"pizza\",\"net_price\":40.92,\"product_id\":6,\"quantity\":3,\"unit_price\":13.64},{\"category\":\"wings\",\"net_price\":24.9,\"product_id\":95,\"quantity\":1,\"unit_price\":24.9},{\"category\":\"pizza\",\"net_price\":38.15,\"product_id\":60,\"quantity\":5,\"unit_price\":7.63},{\"category\":\"wings\",\"net_price\":81.24,\"product_id\":81,\"quantity\":4,\"unit_price\":20.31},{\"category\":\"dessert\",\"net_price\":6.61,\"product_id\":42,\"quantity\":1,\"unit_price\":6.61},{\"category\":\"salad\",\"net_price\":30.64,\"product_id\":6,\"quantity\":4,\"unit_price\":7.66},{\"category\":\"calzone\",\"net_price\":12.92,\"product_id\":36,\"quantity\":4,\"unit_price\":3.23},{\"category\":\"calzone\",\"net_price\":8.64,\"product_id\":27,\"quantity\":1,\"unit_price\":8.64},{\"category\":\"salad\",\"net_price\":31.65,\"product_id\":97,\"quantity\":5,\"unit_price\":6.33},{\"category\":\"salad\",\"net_price\":59.56,\"product_id\":95,\"quantity\":4,\"unit_price\":14.89},{\"category\":\"pizza\",\"net_price\":8.81,\"product_id\":89,\"quantity\":1,\"unit_price\":8.81},{\"category\":\"dessert\",\"net_price\":48.09,\"product_id\":97,\"quantity\":3,\"unit_price\":16.03},{\"category\":\"salad\",\"net_price\":12.68,\"product_id\":25,\"quantity\":4,\"unit_price\":3.17},{\"category\":\"wings\",\"net_price\":10.94,\"product_id\":42,\"quantity\":2,\"unit_price\":5.47},{\"category\":\"salad\",\"net_price\":36.5,\"product_id\":7,\"quantity\":2,\"unit_price\":18.25},{\"category\":\"wings\",\"net_price\":65.84,\"product_id\":4,\"quantity\":4,\"unit_price\":16.46},{\"category\":\"calzone\",\"net_price\":2.15,\"product_id\":78,\"quantity\":1,\"unit_price\":2.15},{\"category\":\"calzone\",\"net_price\":22.3,\"product_id\":71,\"quantity\":1,\"unit_price\":22.3},{\"category\":\"dessert\",\"net_price\":63.45,\"product_id\":80,\"quantity\":5,\"unit_price\":12.69},{\"category\":\"salad\",\"net_price\":1.66,\"product_id\":83,\"quantity\":1,\"unit_price\":1.66},{\"category\":\"dessert\",\"net_price\":72.72,\"product_id\":44,\"quantity\":4,\"unit_price\":18.18},{\"category\":\"calzone\",\"net_price\":8.04,\"product_id\":43,\"quantity\":4,\"unit_price\":2.01},{\"category\":\"calzone\",\"net_price\":39.18,\"product_id\":91,\"quantity\":3,\"unit_price\":13.06},{\"category\":\"wings\",\"net_price\":44.2,\"product_id\":55,\"quantity\":2,\"unit_price\":22.1},{\"category\":\"wings\",\"net_price\":5.33,\"product_id\":88,\"quantity\":1,\"unit_price\":5.33},{\"category\":\"dessert\",\"net_price\":17.36,\"product_id\":10,\"quantity\":2,\"unit_price\":8.68},{\"category\":\"salad\",\"net_price\":2.78,\"product_id\":76,\"quantity\":2,\"unit_price\":1.39},{\"category\":\"salad\",\"net_price\":88.45,\"product_id\":13,\"quantity\":5,\"unit_price\":17.69},{\"category\":\"salad\",\"net_price\":36.88,\"product_id\":27,\"quantity\":4,\"unit_price\":9.22},{\"category\":\"wings\",\"net_price\":21.48,\"product_id\":94,\"quantity\":3,\"unit_price\":7.16},{\"category\":\"dessert\",\"net_price\":20.2,\"product_id\":89,\"quantity\":1,\"unit_price\":20.2},{\"category\":\"wings\",\"net_price\":14.08,\"product_id\":4,\"quantity\":4,\"unit_price\":3.52},{\"category\":\"dessert\",\"net_price\":26.48,\"product_id\":97,\"quantity\":4,\"unit_price\":6.62},{\"category\":\"dessert\",\"net_price\":40.88,\"product_id\":57,\"quantity\":2,\"unit_price\":20.44},{\"category\":\"dessert\",\"net_price\":42.36,\"product_id\":6,\"quantity\":4,\"unit_price\":10.59},{\"category\":\"salad\",\"net_price\":73.8,\"product_id\":86,\"quantity\":3,\"unit_price\":24.6},{\"category\":\"salad\",\"net_price\":35.49,\"product_id\":85,\"quantity\":3,\"unit_price\":11.83},{\"category\":\"wings\",\"net_price\":45.81,\"product_id\":95,\"quantity\":3,\"unit_price\":15.27},{\"category\":\"calzone\",\"net_price\":33.75,\"product_id\":18,\"quantity\":3,\"unit_price\":11.25},{\"category\":\"wings\",\"net_price\":40.59,\"product_id\":85,\"quantity\":3,\"unit_price\":13.53},{\"category\":\"wings\",\"net_price\":66.32,\"product_id\":13,\"quantity\":4,\"unit_price\":16.58},{\"category\":\"salad\",\"net_price\":45.95,\"product_id\":12,\"quantity\":5,\"unit_price\":9.19},{\"category\":\"pizza\",\"net_price\":13.65,\"product_id\":21,\"quantity\":1,\"unit_price\":13.65},{\"category\":\"dessert\",\"net_price\":36.64,\"product_id\":14,\"quantity\":2,\"unit_price\":18.32},{\"category\":\"calzone\",\"net_price\":119.4,\"product_id\":7,\"quantity\":5,\"unit_price\":23.88},{\"category\":\"pizza\",\"net_price\":22.52,\"product_id\":71,\"quantity\":1,\"unit_price\":22.52}]},\"fields\":[{\"name\":\"product_id\",\"type\":\"int\"},{\"name\":\"category\",\"type\":\"string\"},{\"name\":\"quantity\",\"type\":\"int\"},{\"name\":\"unit_price\",\"type\":\"double\"},{\"name\":\"net_price\",\"type\":\"double\"}],\"name\":\"pizza_orders.orderline\",\"type\":\"record\"}"
}

func (r Orderline) SchemaName() string {
	return "pizza_orders.orderline"
}

func (_ Orderline) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Orderline) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Orderline) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Orderline) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Orderline) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Orderline) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Orderline) SetString(v string)   { panic("Unsupported operation") }
func (_ Orderline) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Orderline) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Int{Target: &r.Product_id}

		return w

	case 1:
		w := types.String{Target: &r.Category}

		return w

	case 2:
		w := types.Int{Target: &r.Quantity}

		return w

	case 3:
		w := types.Double{Target: &r.Unit_price}

		return w

	case 4:
		w := types.Double{Target: &r.Net_price}

		return w

	}
	panic("Unknown field index")
}

func (r *Orderline) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *Orderline) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Orderline) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Orderline) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Orderline) HintSize(int)                     { panic("Unsupported operation") }
func (_ Orderline) Finalize()                        {}

func (_ Orderline) AvroCRC64Fingerprint() []byte {
	return []byte(OrderlineAvroCRC64Fingerprint)
}

func (r Orderline) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["product_id"], err = json.Marshal(r.Product_id)
	if err != nil {
		return nil, err
	}
	output["category"], err = json.Marshal(r.Category)
	if err != nil {
		return nil, err
	}
	output["quantity"], err = json.Marshal(r.Quantity)
	if err != nil {
		return nil, err
	}
	output["unit_price"], err = json.Marshal(r.Unit_price)
	if err != nil {
		return nil, err
	}
	output["net_price"], err = json.Marshal(r.Net_price)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Orderline) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["product_id"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Product_id); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for product_id")
	}
	val = func() json.RawMessage {
		if v, ok := fields["category"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Category); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for category")
	}
	val = func() json.RawMessage {
		if v, ok := fields["quantity"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Quantity); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for quantity")
	}
	val = func() json.RawMessage {
		if v, ok := fields["unit_price"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Unit_price); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for unit_price")
	}
	val = func() json.RawMessage {
		if v, ok := fields["net_price"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Net_price); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for net_price")
	}
	return nil
}
