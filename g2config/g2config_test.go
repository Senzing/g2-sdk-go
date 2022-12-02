package g2config

import (
	"context"
	"fmt"
	"log"
	"testing"

	truncator "github.com/aquilax/truncate"
	"github.com/senzing/go-helpers/g2engineconfigurationjson"
	"github.com/senzing/go-logging/logger"
	"github.com/stretchr/testify/assert"
)

const (
	defaultTruncation = 76
)

var (
	g2configSingleton G2config
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) G2config {
	if g2configSingleton == nil {
		g2configSingleton = &G2configImpl{}

		// g2configSingleton.SetLogLevel(ctx, logger.LevelTrace)
		log.SetFlags(0)

		moduleName := "Test module name"
		verboseLogging := 0
		iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
		if jsonErr != nil {
			test.Logf("Cannot construct system configuration. Error: %v", jsonErr)
		}
		initErr := g2configSingleton.Init(ctx, moduleName, iniParams, verboseLogging)
		if initErr != nil {
			test.Logf("Cannot Init. Error: %v", initErr)
		}
	}
	return g2configSingleton
}

func truncate(aString string, length int) string {
	return truncator.Truncate(aString, length, "...", truncator.PositionEnd)
}

func printResult(test *testing.T, title string, result interface{}) {
	if 1 == 0 {
		test.Logf("%s: %v", title, truncate(fmt.Sprintf("%v", result), defaultTruncation))
	}
}

func printActual(test *testing.T, actual interface{}) {
	printResult(test, "Actual", actual)
}

func testError(test *testing.T, ctx context.Context, g2config G2config, err error) {
	if err != nil {
		lastException, _ := g2config.GetLastException(ctx)
		test.Log("Error:", err.Error())
		assert.FailNow(test, lastException)
	}
}

func testErrorNoFail(test *testing.T, ctx context.Context, g2config G2config, err error) {
	if err != nil {
		lastException, _ := g2config.GetLastException(ctx)
		test.Log("Error:", err.Error(), "LastException:", lastException)
	}
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestBuildSimpleSystemConfigurationJson(test *testing.T) {
	actual, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, actual)
	}
	printActual(test, actual)
}

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestG2configImpl_AddDataSource(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	actual, err := g2config.AddDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_ClearLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	err := g2config.ClearLastException(ctx)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_Close(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_Create(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestG2configImpl_DeleteDataSource(test *testing.T) {

	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)

	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "Original", actual)

	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	_, err = g2config.AddDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "     Add", actual)

	err = g2config.DeleteDataSource(ctx, configHandle, inputJson)
	testError(test, ctx, g2config, err)

	actual, err = g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	printResult(test, "  Delete", actual)

	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_GetLastException(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.GetLastException(ctx)
	testErrorNoFail(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestG2configImpl_GetLastExceptionCode(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	actual, err := g2config.GetLastExceptionCode(ctx)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestG2configImpl_Init(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	moduleName := "Test module name"
	verboseLogging := 0
	iniParams, jsonErr := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	testError(test, ctx, g2config, jsonErr)
	err := g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_ListDataSources(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.ListDataSources(ctx, configHandle)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
	err = g2config.Close(ctx, configHandle)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_Load(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	jsonConfig, err := g2config.Save(ctx, configHandle)
	testError(test, ctx, g2config, err)
	err = g2config.Load(ctx, configHandle, jsonConfig)
	testError(test, ctx, g2config, err)
}

func TestG2configImpl_Save(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	configHandle, err := g2config.Create(ctx)
	testError(test, ctx, g2config, err)
	actual, err := g2config.Save(ctx, configHandle)
	testError(test, ctx, g2config, err)
	printActual(test, actual)
}

func TestG2configImpl_Destroy(test *testing.T) {
	ctx := context.TODO()
	g2config := getTestObject(ctx, test)
	err := g2config.Destroy(ctx)
	testError(test, ctx, g2config, err)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleG2configImpl_AddDataSource() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	inputJson := `{"DSRC_CODE": "GO_TEST"}`
	result, err := g2config.AddDataSource(ctx, configHandle, inputJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"DSRC_ID":1001}
}

func ExampleG2configImpl_ClearLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	err := g2config.ClearLastException(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_Close() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	err = g2config.Close(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_Create() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configHandle > 0) // Dummy output.
	// Output: true
}

func ExampleG2configImpl_DeleteDataSource() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	inputJson := `{"DSRC_CODE": "TEST"}`
	err = g2config.DeleteDataSource(ctx, configHandle, inputJson)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_Destroy() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	err := g2config.Destroy(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_GetLastException() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	result, err := g2config.GetLastException(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output:
}

func ExampleG2configImpl_GetLastExceptionCode() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	result, err := g2config.GetLastExceptionCode(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: 0
}

func ExampleG2configImpl_Init() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	moduleName := "Test module name"
	iniParams, err := g2engineconfigurationjson.BuildSimpleSystemConfigurationJson("")
	if err != nil {
		fmt.Println(err)
	}
	verboseLogging := 0
	err = g2config.Init(ctx, moduleName, iniParams, verboseLogging)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_ListDataSources() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	result, err := g2config.ListDataSources(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Output: {"DATA_SOURCES":[{"DSRC_ID":1,"DSRC_CODE":"TEST"},{"DSRC_ID":2,"DSRC_CODE":"SEARCH"}]}
}

func ExampleG2configImpl_Load() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	jsonConfig, err := g2config.Save(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	err = g2config.Load(ctx, configHandle, jsonConfig)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}

func ExampleG2configImpl_Save() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	configHandle, err := g2config.Create(ctx)
	if err != nil {
		fmt.Println(err)
	}
	jsonConfig, err := g2config.Save(ctx, configHandle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(truncate(jsonConfig, defaultTruncation))
	// Output: {"G2_CONFIG":{"CFG_ATTR":[{"ATTR_ID":1001,"ATTR_CODE":"DATA_SOURCE","ATTR...
}

func ExampleG2configImpl_SetLogLevel() {
	// For more information, visit https://github.com/Senzing/g2-sdk-go/blob/main/g2config/g2config_test.go
	g2config := &G2configImpl{}
	ctx := context.TODO()
	err := g2config.SetLogLevel(ctx, logger.LevelInfo)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
}
