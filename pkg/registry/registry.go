package registry

import (
	"fmt"
	"github.com/srikartati/go-ipfixlib/pkg/entities"
	"strings"
)

const reversePen = uint32(29305)

type Registry interface {
	LoadRegistry()
	GetInfoElement(name string)  (*entities.InfoElement, error)
	GetReverseInfoElement(name string) (*entities.InfoElement, error)
}

type ianaRegistry struct {
	registry map[string]entities.InfoElement
}

type antreaRegistry struct {
	registry map[string]entities.InfoElement
}


func NewIanaRegistry() *ianaRegistry {
	reg := make(map[string]entities.InfoElement)
	return &ianaRegistry{
		registry: reg,
	}
}

func NewAntreaRegistry() *antreaRegistry {
	reg := make(map[string]entities.InfoElement)
	return &antreaRegistry{
		registry: reg,
	}
}

func (reg *ianaRegistry) registerInfoElement(ie entities.InfoElement) error {
	if _, exist := reg.registry[ie.Name]; exist {
		return fmt.Errorf("IANA Registry: Information element %s has already been registered", ie.Name)
	}
	reg.registry[ie.Name] = ie
	return nil
}

func (reg *ianaRegistry) GetInfoElement(name string) (*entities.InfoElement, error) {
	var exist bool
	var ie entities.InfoElement
	if ie, exist = reg.registry[name]; !exist {
		err := fmt.Errorf("IANA Registry: There is no information element with name %s", name)
		return &ie, err
	}
	return &ie, nil
}

func (reg *ianaRegistry) GetReverseInfoElement(name string) (*entities.InfoElement, error) {
	var exist bool
	var ie entities.InfoElement
	if ie, exist = reg.registry[name]; !exist {
		err := fmt.Errorf("IANA Registry: There is no information element with name %s", name)
		return &ie, err
	}
	if !isReversible(ie.Name) {
		err := fmt.Errorf("IANA Registry: The information element %s is not reverse element", name)
		return &ie, err
	}
	reverseName := "reverse" + strings.Title(ie.Name)
	return entities.NewInfoElement(reverseName, ie.ElementId, ie.DataType, reversePen, ie.Len), nil
}

func (reg *antreaRegistry) registerInfoElement(ie entities.InfoElement) error {
	if _, exist := reg.registry[ie.Name]; exist {
		return fmt.Errorf("Antrea Registry: Information element %s has already been registered", ie.Name)
	}
	reg.registry[ie.Name] = ie
	return nil
}

func (reg *antreaRegistry) GetInfoElement(name string) (*entities.InfoElement, error) {
	var exist bool
	var ie entities.InfoElement
	if ie, exist = reg.registry[name]; !exist {
		err := fmt.Errorf("Antrea Registry: There is no information element with name %s", name)
		return &ie, err
	}
	return &ie, nil
}

func (reg *antreaRegistry) GetReverseInfoElement(name string) (*entities.InfoElement, error) {
	var exist bool
	var ie entities.InfoElement
	if ie, exist = reg.registry[name]; !exist {
		err := fmt.Errorf("Antrea Registry: There is no information element with name %s", name)
		return &ie, err
	}
	reverseName := "reverse" + strings.Title(ie.Name)
	return entities.NewInfoElement(reverseName, ie.ElementId, ie.DataType, reversePen, ie.Len), nil
}

var nonReversibleIEList = [...]string{
	"biflowDirection",
	"collectorIPv4Address",
	"collectorIPv6Address",
	"collectorTransportPort",
	"commonPropertiesId",
	"exportedMessageTotalCount",
	"exportedOctetTotalCount",
	"exportedFlowRecordTotalCount",
	"exporterIPv4Address",
	"exporterIPv6Address",
	"exporterTransportPort",
	"exportInterface",
	"exportProtocolVersion",
	"exportTransportProtocol",
	"flowId",
	"flowKeyIndicator",
	"ignoredPacketTotalCount",
	"ignoredOctetTotalCount",
	"notSentFlowTotalCount",
	"notSentPacketTotalCount",
	"notSentOctetTotalCount",
	"observationDomainId",
	"observedFlowTotalCount",
	"paddingOctets",
	"templateId",
}

func isReversible(name string) bool {
	for _, ieName := range nonReversibleIEList {
		if ieName == name {
			return false
		}
	}
	return true
}