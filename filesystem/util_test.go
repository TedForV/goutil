package filesystem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPathExisted(t *testing.T) {
	result, err := IsPathExisted("D:/Project_Golang/src/rpc_notifier/model/sms.go")
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, true, result, "")

	result, err = IsPathExisted("./util.go")
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, true, result, "")

	result, err = IsPathExisted("./util1.go")
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, false, result, "")
}
