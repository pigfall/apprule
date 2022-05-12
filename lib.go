package apprule


func HowManyLicenseToBuy(column []string, rows [][]string)(int,error){
	hasChanced := &LicensesHasChanced{}
	noChanced := &LicensesNoChanced{}

	for _, row := range rows {
		device := &Device{
			Meta:column,
			Values:row,
		}
		license,slotIndex := hasChanced.FindMatchedLicense(device)
		if license != nil{ 
			license.Update(device,slotIndex)
			if license.NoLeftChance(){ // the chance of license has been used
				noChanced.Add(license)
				hasChanced.Remove(license)
			}
			continue
		}

		// need to buy a new license
		l := NewLicense()
		l.MustUpdateByFindSlot(device)
		if l.NoLeftChance() {
				noChanced.Add(l)
		}else{
			hasChanced.Add(l)
		}
	}

	return hasChanced.Count() + noChanced.Count(),nil
}

type LicensesHasChanced struct{
	Licenses []*License
}

func (this *LicensesHasChanced) Count()int{
	panic("TODO")
}

func (this *LicensesHasChanced) FindMatchedLicense(device *Device)(l *License,slotIndex int){
	for _,l := range this.Licenses{
		slotIndex := l.findMatchedSlot(device)
		if slotIndex < 0{
			continue
		}
		return l ,slotIndex
	}

	// not found
	return nil,0
}

func (this *LicensesHasChanced) Remove(l *License){
	panic("TODO")
}

func (this *LicensesHasChanced) Add(l *License){
	panic("TODO")

}

type LicensesNoChanced struct{}

func (this *LicensesNoChanced) Add(l *License){
	panic("TODO")
}

func (this LicensesNoChanced) Count()int{
	panic("TODO")
}
