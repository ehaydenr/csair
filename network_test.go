package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestLongestFlight(t *testing.T) {
	//{[SYD LAX] 12051}
	answer := FlightPath{
		[]string{
			"SYD",
			"LAX",
		},
		12051,
	}

	proposal := network.computeLongestFlight()

	if !reflect.DeepEqual(answer, proposal) {
		t.Fail()
	}
}

func TestShortestFlight(t *testing.T) {
	//{[WAS NYC] 334}
	answer := FlightPath{
		[]string{
			"WAS",
			"NYC",
		},
		334,
	}

	proposal := network.computeShortestFlight()

	if !reflect.DeepEqual(answer, proposal) {
		t.Fail()
	}
}

func TestAverageDistance(t *testing.T) {
	//2300
	answer := 2300
	proposal := network.computeAverageDistance()

	if answer != proposal {
		t.Fail()
	}
}

func TestBiggestCity(t *testing.T) {
	//TYO
	answer := "TYO"
	proposal := network.computeBiggestCity()

	if answer != proposal {
		t.Fail()
	}
}

func TestSmallestCity(t *testing.T) {
	//ESS
	answer := "ESS"
	proposal := network.computeSmallestCity()

	if answer != proposal {
		t.Fail()
	}
}

func TestAveragePopulation(t *testing.T) {
	//11796144
	answer := 11796144
	proposal := network.computeAveragePopulation()

	if answer != proposal {
		t.Fail()
	}
}

func TestContinentList(t *testing.T) {
	//Australia Europe Asia South America Africa North America
	answer := []string{
		"Australia", "Europe", "Asia", "South America", "Africa",
		"North America",
	}

	proposal := network.computeListOfContinents()

	sort.Strings(answer)
	sort.Strings(proposal)

	if !reflect.DeepEqual(answer, proposal) {
		t.Fail()
	}
}

func TestHubList(t *testing.T) {
	//HKG IST
	answer := []string{
		"HKG",
		"IST",
	}

	proposal := network.computeHubCities()

	sort.Strings(answer)
	sort.Strings(proposal)

	if !reflect.DeepEqual(answer, proposal) {
		t.Fail()
	}
}

func TestMapUrl(t *testing.T) {
	answer := "http://www.gcmap.com/mapui?P=SCL-LIM,LIM-MEX,LIM-BOG,MEX-LAX,MEX-CHI,MEX-MIA,MEX-BOG,BOG-MIA,BOG-SAO,BOG-BUE,BUE-SAO,SAO-MAD,SAO-LOS,LOS-KRT,LOS-FIH,FIH-KRT,FIH-JNB,JNB-KRT,KRT-CAI,CAI-ALG,CAI-IST,CAI-BGW,CAI-RUH,ALG-MAD,ALG-PAR,ALG-IST,MAD-NYC,MAD-LON,MAD-PAR,LON-NYC,LON-ESS,LON-PAR,PAR-ESS,PAR-MIL,MIL-ESS,MIL-IST,ESS-LED,LED-MOW,LED-IST,MOW-THR,MOW-IST,IST-BGW,BGW-THR,BGW-KHI,BGW-RUH,THR-DEL,THR-KHI,THR-RUH,RUH-KHI,KHI-DEL,KHI-BOM,DEL-CCU,DEL-MAA,DEL-BOM,BOM-MAA,MAA-CCU,MAA-BKK,MAA-JKT,CCU-HKG,CCU-BKK,BKK-HKG,BKK-SGN,BKK-JKT,HKG-SHA,HKG-TPE,HKG-MNL,HKG-SGN,SHA-PEK,SHA-ICN,SHA-TYO,SHA-TPE,PEK-ICN,ICN-TYO,TYO-SFO,TYO-OSA,OSA-TPE,TPE-MNL,MNL-SFO,MNL-SYD,MNL-SGN,SGN-JKT,JKT-SYD,SYD-LAX,LAX-SFO,LAX-CHI,SFO-CHI,CHI-YYZ,CHI-ATL,ATL-WAS,ATL-MIA,MIA-WAS,WAS-YYZ,WAS-NYC,NYC-YYZ"

	proposal := network.computeMapUrl()
	if answer != proposal {
		t.Fail()
	}
}

func TestShortestRoute(t *testing.T) {
	// [ATL CHI SFO TYO]
	answer := []string{
		"ATL", "CHI", "SFO", "TYO",
	}
	proposal := network.computeShortestRoute("ATL", "TYO")

	if !reflect.DeepEqual(answer, proposal) {
		t.Fail()
	}
}
