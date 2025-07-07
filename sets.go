package dmidecode

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yxxhero/dmidecode/parser/power"

	"github.com/yxxhero/dmidecode/parser/baseboard"
	"github.com/yxxhero/dmidecode/parser/battery"
	"github.com/yxxhero/dmidecode/parser/bios"
	"github.com/yxxhero/dmidecode/parser/chassis"
	"github.com/yxxhero/dmidecode/parser/memory"
	"github.com/yxxhero/dmidecode/parser/oem"
	"github.com/yxxhero/dmidecode/parser/onboard"
	"github.com/yxxhero/dmidecode/parser/port"
	"github.com/yxxhero/dmidecode/parser/processor"
	"github.com/yxxhero/dmidecode/parser/slot"
	"github.com/yxxhero/dmidecode/parser/system"
	"github.com/yxxhero/dmidecode/parser/tpm"
)

// NewInformationSet todo
func NewInformationSet() *InformationSet {
	return &InformationSet{
		bios:            []*bios.Information{},
		system:          []*system.Information{},
		baseboard:       []*baseboard.Information{},
		chassis:         []*chassis.Information{},
		onboard:         []*onboard.Information{},
		onboardExtended: []*onboard.ExtendedInformation{},
		portConnetor:    []*port.Information{},
		processor:       []*processor.Processor{},
		cache:           []*processor.Cache{},
		memoryArray:     []*memory.PhysicalMemoryArray{},
		memoryDevice:    []*memory.MemoryDevice{},
		slot:            []*slot.SystemSlot{},
		battery:         []*battery.Information{},
		tpm:             []*tpm.Information{},
		power:           []*power.Information{},
	}
}

// InformationSet 集合
type InformationSet struct {
	oem             []*oem.OEM
	bios            []*bios.Information
	system          []*system.Information
	baseboard       []*baseboard.Information
	chassis         []*chassis.Information
	onboard         []*onboard.Information
	onboardExtended []*onboard.ExtendedInformation
	portConnetor    []*port.Information
	processor       []*processor.Processor
	cache           []*processor.Cache
	memoryArray     []*memory.PhysicalMemoryArray
	memoryDevice    []*memory.MemoryDevice
	slot            []*slot.SystemSlot
	battery         []*battery.Information
	tpm             []*tpm.Information
	power           []*power.Information
}

// Print 打印所有
func (s *InformationSet) Println() {
	for i := range s.bios {
		fmt.Println(s.bios[i])
	}
	for i := range s.system {
		fmt.Println(s.system[i])
	}
	for i := range s.baseboard {
		fmt.Println(s.baseboard[i])
	}
	for i := range s.chassis {
		fmt.Println(s.chassis[i])
	}
	for i := range s.onboard {
		fmt.Println(s.onboard[i])
	}
	for i := range s.onboardExtended {
		fmt.Println(s.onboardExtended[i])
	}
	for i := range s.portConnetor {
		fmt.Println(s.portConnetor[i])
	}
	for i := range s.processor {
		fmt.Println(s.processor[i])
	}
	for i := range s.cache {
		fmt.Println(s.cache[i])
	}
	for i := range s.memoryArray {
		fmt.Println(s.memoryArray[i])
	}
	for i := range s.memoryDevice {
		fmt.Println(s.memoryDevice[i])
	}
	for i := range s.slot {
		fmt.Println(s.slot[i])
	}
	for i := range s.battery {
		fmt.Println(s.battery[i])
	}
	for i := range s.tpm {
		fmt.Println(s.tpm[i])
	}
	for i := range s.power {
		fmt.Println(s.power[i])
	}
}

func (s *InformationSet) addBios(infos []*bios.Information) {
	s.bios = infos
}

func (s *InformationSet) addSystem(infos []*system.Information) {
	s.system = infos
}

func (s *InformationSet) addBaseBoard(infos []*baseboard.Information) {
	s.baseboard = infos
}

func (s *InformationSet) addChassis(infos []*chassis.Information) {
	s.chassis = infos
}

func (s *InformationSet) addOEM(infos []*oem.OEM) {
	s.oem = infos
}

func (s *InformationSet) addOnboard(infos []*onboard.ExtendedInformation) {
	s.onboardExtended = infos
}

func (s *InformationSet) addPortConnector(infos []*port.Information) {
	s.portConnetor = infos
}

func (s *InformationSet) addProcessor(infos []*processor.Processor) {
	s.processor = infos
}

func (s *InformationSet) addCache(infos []*processor.Cache) {
	s.cache = infos
}

func (s *InformationSet) addMemoryArray(infos []*memory.PhysicalMemoryArray) {
	s.memoryArray = infos
}

func (s *InformationSet) addMemoryDevice(infos []*memory.MemoryDevice) {
	s.memoryDevice = infos
}

func (s *InformationSet) addSlot(infos []*slot.SystemSlot) {
	s.slot = infos
}

func (s *InformationSet) addBattery(infos []*battery.Information) {
	s.battery = infos
}

func (s *InformationSet) addTpm(infos []*tpm.Information) {
	s.tpm = infos
}

func (s *InformationSet) addPower(infos []*power.Information) {
	s.power = infos
}

// NewErrorSet todo
func NewErrorSet() *ErrorSet {
	return &ErrorSet{
		errs: []error{},
	}
}

// ErrorSet errorset
type ErrorSet struct {
	errs []error
}

func (s *ErrorSet) checkOrAdd(err error) {
	if err != nil {
		s.errs = append(s.errs, err)
	}
}

// HasError todo
func (s *ErrorSet) Error() error {
	if len(s.errs) > 0 {
		errMSG := make([]string, 0, len(s.errs))
		for i := range s.errs {
			errMSG = append(errMSG, s.errs[i].Error())
		}
		errStrs := strings.Join(errMSG, ",")
		return errors.New(errStrs)
	}

	return nil
}
