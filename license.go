package apprule


// the license of app
type License struct{
	Slots []*Slot
	OccupiedSlotCount int
}

func NewLicense() *License{
	return &License{}
}

func (this *License) Update(device *Device,slotIndex int){
	slot:= this.Slots[slotIndex]
	slot.Occupy =&SlotOccupy{
		Device:device,
	}
	this.OccupiedSlotCount ++
}

// panic if not found matched slot
func (this *License) MustUpdateByFindSlot(device *Device){
	panic("TODO")
}

func (this *License) findMatchedSlot(devcie *Device)(slotIndex int){
	panic("TODO")
}

func (this *License) NoLeftChance() bool{
	return this.OccupiedSlotCount == len(this.Slots)
}


type Slot struct {
	Occupy *SlotOccupy
	RuleExpr Expr
}

type SlotOccupy struct{
	Device *Device
}
