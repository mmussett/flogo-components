package redis

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"fmt"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestGetKey1(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
  tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivPassword, "")
  tc.SetInput(ivMethod,methodGet)
  tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
  if err != nil {
  	t.Error("unable to get key value",err)
  	t.Fail()
  }

  fmt.Println(tc.GetOutput(ovResult))


}

func TestGetKey2(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivPassword, "")
	tc.SetInput(ivMethod,methodGet)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to get key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestExists(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivPassword, "")
	tc.SetInput(ivMethod,methodExists)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to exist key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}


func TestSetStringVal(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	//tc.SetInput(ivExpiration, 10000)
	tc.SetInput(ivMethod,methodSet)
	tc.SetInput(ivKeyValue,"hello")
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestSetIntVal(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivMethod,methodSet)
	tc.SetInput(ivKeyValue,"10")
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestIncr(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivMethod,methodIncr)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestDecr(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivMethod,methodDecr)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestIncrBy(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivKeyValue, "2")
	tc.SetInput(ivMethod,methodIncrBy)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestDecrBy(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	tc.SetInput(ivKeyValue, "2")
	tc.SetInput(ivMethod,methodDecrBy)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestExpire(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivExpiration, 100)
	tc.SetInput(ivMethod,methodExpire)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set expire on key",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestPersistKey1(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivMethod,methodPersist)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set persist on key",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}


func TestPExpire(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivExpiration, 36000)
	tc.SetInput(ivMethod,methodPExpire)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set expire on key",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestDelete(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivMethod,methodDelete)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to delete key",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestPing(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivMethod,methodPing)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to ping redis",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}



func TestTTL(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivMethod,methodTTL)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to get key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestPTTL(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	tc.SetInput(ivMethod,methodPTTL)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to get key value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestRename(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key2")
	//tc.SetInput(ivExpiration, 10000)
	tc.SetInput(ivMethod,methodRename)
	tc.SetInput(ivKeyValue,"key1")
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set rename key",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}

func TestAppend(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivKeyName,"key1")
	//tc.SetInput(ivExpiration, 10000)
	tc.SetInput(ivMethod,methodAppend)
	tc.SetInput(ivKeyValue,", world")
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to set append value",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))


}


func TestFlushAll(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivMethod,methodFlushAll)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to flush all",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}

func TestFlushDB(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())


	//setup attrs
	tc.SetInput(ivAddress, "localhost:12000")
	tc.SetInput(ivMethod,methodFlushDB)
	tc.SetInput(ivDatabase,0)

	_,err := act.Eval(tc)
	if err != nil {
		t.Error("unable to flush DB",err)
		t.Fail()
	}

	fmt.Println(tc.GetOutput(ovResult))

}