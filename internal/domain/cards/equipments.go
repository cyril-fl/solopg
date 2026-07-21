package cards

type Equipment struct {
	Helmet   *EquipmentGear
	Chestplate  *EquipmentGear
	Gauntlets  *EquipmentGear
	Greaves   *EquipmentGear
	Boots   *EquipmentGear
	RightHand *EquipmentGear
	LeftHand  *EquipmentGear
}

type EquipmentGear struct {
	Item
	DestinedSlot EquipmentSlot
}

type EquipmentSlot string

const (	
	Helmet EquipmentSlot = "helmet"
    Chestplate  EquipmentSlot = "chestplate"
    Gauntlets EquipmentSlot = "gauntlets"
    Greaves EquipmentSlot = "greaves"
    Boots  EquipmentSlot = "boots"
    RightHand EquipmentSlot = "right_hand"
	LeftHand  EquipmentSlot = "left_hand"
)