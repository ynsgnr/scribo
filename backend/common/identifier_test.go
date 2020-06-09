package common_test

import (
	"testing"

	"github.com/ynsgnr/scribo/backend/common"
)

func TestGenerateIDS(t *testing.T) {
	mail := "mail@mail.com"
	is := "internalsecret"
	es := "externalsecret"
	iID, eID, err := common.CalculateIDs(mail, is, es)
	if err != nil {
		t.Error(err)
		return
	}
	iIDfromExternal, err := common.CalculateInternalID(eID, is)
	if err != nil {
		t.Error(err)
		return
	}
	if iIDfromExternal != iID {
		t.Errorf("Expected %s, Actual %s", iID, iIDfromExternal)
	}
}
