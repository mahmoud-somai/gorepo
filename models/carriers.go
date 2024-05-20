// carrier.go

package models

type Carrier struct {
	ID                 uint `gorm:"primaryKey"`
	Tax_Rules_Group_ID uint
	Name               string
	Max_Width          int
	Max_Height         int
	Max_Depth          int
	Max_Weight         float64
	Active             uint
	Is_Free            uint
	URL                string
	Shipping_Handling  uint
	Shipping_External  uint
	Range_Behavior     uint
	Shipping_Method    int
	Grade              int
	Position           int
	Deleted            uint
}

func (Carrier) TableName() string {
	return "carriers"
}
