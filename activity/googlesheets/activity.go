package googlesheets

import (
	"errors"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/opentracing/opentracing-go"
	ctx "golang.org/x/net/context"

	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const (
	methodGet      = "GET"
	methodSet      = "SET"
	methodDelete   = "DELETE"
	methodPing     = "PING"
	methodExists   = "EXISTS"
	methodExpire   = "EXPIRE"
	methodPersist  = "PERSIST"
	methodPExpire  = "PEXPIRE"
	methodTTL      = "TTL"
	methodPTTL     = "PTTL"
	methodRename   = "RENAME"
	methodAppend   = "APPEND"
	methodFlushAll = "FLUSHALL"
	methodFlushDB  = "FLUSHDB"
	methodDecr     = "DECR"
	methodIncr     = "INCR"
	methodDecrBy   = "DECRBY"
	methodIncrBy   = "INCRBY"

	ivMethod     = "method"
	ivAddress    = "address"
	ivPassword   = "password"
	ivDatabase   = "database"
	ivKeyName    = "keyName"
	ivKeyValue   = "keyValue"
	ivExpiration = "expiration"
	ivTracing    = "tracing"

	ovResult  = "result"
	ovTracing = "tracing"
)

var (
	errorUnableToGetKeyValue   = errors.New("unable to get key value")
	errorUnableToGetTTLValue   = errors.New("unable to get TTL value")
	errorUnableToPing          = errors.New("unable to ping redis")
	errorUnableToDoExist       = errors.New("unable to perform exist method")
	errorUnableToSetExpiration = errors.New("unable to set expiration")
	errorUnableToSetPersist    = errors.New("unable to set persist")
	errorUnableToRenameKey     = errors.New("unable to rename key")
	errorUnableToDeleteKey     = errors.New("unable to delete key")
	errorUnableToFlushAll      = errors.New("unable to flush all")
	errorUnableToFlushDB       = errors.New("unable to flush db")
	errorUnableToDecr          = errors.New("unable to decrement key value")
	errorUnableToIncr          = errors.New("unable to increment key value")
	errorNotAnInteger          = errors.New("value not an integer")
)

var log = logger.GetLogger("activity-tibco-redis")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	var span opentracing.Span
	if tracing := context.GetInput(ivTracing); tracing != nil {
		span = opentracing.SpanFromContext(tracing.(ctx.Context))
	}

	if span != nil {
		span = opentracing.StartSpan(
			context.TaskName(),
			opentracing.ChildOf(span.Context()))
		context.SetOutput(ovTracing, opentracing.ContextWithSpan(ctx.Background(), span))
		defer span.Finish()
	}

	setTag := func(key string, value interface{}) {
		if span != nil {
			span.SetTag(key, value)
		}
	}

	logError := func(format string, a ...interface{}) {
		str := fmt.Sprintf(format, a...)
		setTag("error", str)
		log.Error(str)
	}

	action := context.GetInput(ivMethod).(string)
	address := context.GetInput(ivAddress).(string)
	password := context.GetInput(ivPassword).(string)
	database := context.GetInput(ivDatabase).(int)
	key := context.GetInput(ivKeyName).(string)
	value := context.GetInput(ivKeyValue).(string)
	expiration := context.GetInput(ivExpiration).(int)

	expiration, ok := context.GetInput(ivExpiration).(int)
	if !ok {
		expiration = -1
	}

	setTag("address", address)
	setTag("key", key)
	setTag("value", value)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       database,
	})

	switch action {

	case methodPing:

		val, err := client.Ping().Result()
		client.Close()
		if err != nil {
			logError(errorUnableToPing.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodExists:

		val, err := client.Exists(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToDoExist.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		if val == 0 {
			context.SetOutput(ovResult, "false")
		} else {
			context.SetOutput(ovResult, "true")
		}
		return true, nil

	case methodExpire:

		val, err := client.Expire(key, time.Duration(expiration)*time.Second).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToSetExpiration.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodPersist:

		val, err := client.Persist(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToSetPersist.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodPExpire:

		val, err := client.PExpire(key, time.Duration(expiration)*time.Millisecond).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToSetExpiration.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodGet:

		val, err := client.Get(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToGetKeyValue.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodTTL:

		val, err := client.TTL(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToGetTTLValue.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		if val == -2000000000 {
			context.SetOutput(ovResult, -2)
		} else if val == -1000000000 {
			context.SetOutput(ovResult, -1)
		} else {
			// Set the output value in the context
			context.SetOutput(ovResult, val)
		}

		return true, nil

	case methodPTTL:

		val, err := client.PTTL(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToGetTTLValue.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		if val == -2000000 {
			context.SetOutput(ovResult, -2)
		} else if val == -1000000 {
			context.SetOutput(ovResult, -1)
		} else {
			// Set the output value in the context
			context.SetOutput(ovResult, val/time.Millisecond)
		}

		return true, nil

	case methodSet:

		if expiration >= 0 {
			client.Set(key, value, time.Duration(expiration)*time.Second)
		}

		if expiration == -1 {
			client.Persist(key)
		}

		client.Close()

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil

	case methodAppend:

		val, err := client.Append(key, value).Result()
		client.Close()

		if err != nil {
			logError(errorUnableToRenameKey.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodRename:

		_, err := client.Rename(key, value).Result()
		client.Close()

		if err != nil {
			logError(errorUnableToRenameKey.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, "OK")
		return true, nil

	case methodDelete:

		val, err := client.Del(key).Result()
		client.Close()

		if err != nil {
			logError(errorUnableToDeleteKey.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodFlushAll:

		_, err := client.FlushAll().Result()
		client.Close()

		if err != nil {
			logError(errorUnableToFlushAll.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		context.SetOutput(ovResult, "OK")
		return true, nil

	case methodFlushDB:

		_, err := client.FlushDB().Result()
		client.Close()

		if err != nil {
			logError(errorUnableToFlushDB.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		context.SetOutput(ovResult, "OK")
		return true, nil

	case methodDecr:

		val, err := client.Decr(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToDecr.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodIncr:

		val, err := client.Incr(key).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToIncr.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodDecrBy:

		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logError(errorNotAnInteger.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		val, err := client.DecrBy(key, n).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToDecr.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	case methodIncrBy:

		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			logError(errorNotAnInteger.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		val, err := client.IncrBy(key, n).Result()
		client.Close()
		if err != nil {
			logError(errorUnableToIncr.Error())
			// Set the output value in the context
			context.SetOutput(ovResult, err.Error())
			return true, err
		}

		// Set the output value in the context
		context.SetOutput(ovResult, val)
		return true, nil

	}

	// Set the output value in the context
	context.SetOutput(ovResult, "NOK")
	return true, nil
}
