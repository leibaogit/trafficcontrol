package v13

/*

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"testing"

	"github.com/apache/incubator-trafficcontrol/lib/go-log"
	"github.com/apache/incubator-trafficcontrol/lib/go-tc"
)

func TestDivisions(t *testing.T) {

	CreateTestDivisions(t)
	UpdateTestDivisions(t)
	GetTestDivisions(t)
	DeleteTestDivisions(t)

}

func CreateTestDivisions(t *testing.T) {
	for _, division := range testData.Divisions {
		resp, _, err := TOSession.CreateDivision(division)
		log.Debugln("Response: ", resp)
		if err != nil {
			t.Errorf("could not CREATE division: %v\n", err)
		}
	}
}

func UpdateTestDivisions(t *testing.T) {

	firstDivision := testData.Divisions[0]
	// Retrieve the Division by division so we can get the id for the Update
	resp, _, err := TOSession.GetDivisionByName(firstDivision.Name)
	if err != nil {
		t.Errorf("cannot GET Division by division: %v - %v\n", firstDivision.Name, err)
	}
	remoteDivision := resp[0]
	expectedDivision := "division-test"
	remoteDivision.Name = expectedDivision
	var alert tc.Alerts
	alert, _, err = TOSession.UpdateDivisionByID(remoteDivision.ID, remoteDivision)
	if err != nil {
		t.Errorf("cannot UPDATE Division by id: %v - %v\n", err, alert)
	}

	// Retrieve the Division to check division got updated
	resp, _, err = TOSession.GetDivisionByID(remoteDivision.ID)
	if err != nil {
		t.Errorf("cannot GET Division by division: %v - %v\n", firstDivision.Name, err)
	}
	respDivision := resp[0]
	if respDivision.Name != expectedDivision {
		t.Errorf("results do not match actual: %s, expected: %s\n", respDivision.Name, expectedDivision)
	}

}

func GetTestDivisions(t *testing.T) {
	for _, division := range testData.Divisions {
		resp, _, err := TOSession.GetDivisionByName(division.Name)
		if err != nil {
			t.Errorf("cannot GET Division by division: %v - %v\n", err, resp)
		}
	}
}

func DeleteTestDivisions(t *testing.T) {

	division := testData.Divisions[1]
	// Retrieve the Division by name so we can get the id
	resp, _, err := TOSession.GetDivisionByName(division.Name)
	if err != nil {
		t.Errorf("cannot GET Division by name: %v - %v\n", division.Name, err)
	}
	respDivision := resp[0]

	delResp, _, err := TOSession.DeleteDivisionByID(respDivision.ID)
	if err != nil {
		t.Errorf("cannot DELETE Division by division: %v - %v\n", err, delResp)
	}

	// Retrieve the Division to see if it got deleted
	divisionResp, _, err := TOSession.GetDivisionByName(division.Name)
	if err != nil {
		t.Errorf("error deleting Division division: %s\n", err.Error())
	}
	if len(divisionResp) > 0 {
		t.Errorf("expected Division : %s to be deleted\n", division.Name)
	}
}
