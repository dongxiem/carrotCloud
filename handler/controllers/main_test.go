package controllers

import (
	"carrotCloud/pkg/serializer"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v8"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	asserts := assert.New(t)

	type Test struct {
		UserName string `validate:"required,min=10,max=20"`
	}
	// 第一个测试案例, UserName: ""
	testCase1 := Test{}
	validate := validator.New(&validator.Config{TagName: "validate"})
	errs := validate.Struct(testCase1)
	res := ErrorResponse(errs)
	// 断言两者是否返回的错误一致
	asserts.Equal(serializer.ParamErr(ParamErrorMsg("UserName", "required"), errs), res)

	// 第二个测试案例，UserName："a"
	testCase2 := Test{
		UserName: "a",
	}
	validate2 := validator.New(&validator.Config{TagName: "validate"})
	errs = validate2.Struct(testCase2)
	res = ErrorResponse(errs)
	asserts.Equal(serializer.ParamErr(ParamErrorMsg("UserName", "min"), errs), res)

	// 第三个测试案例，JSON类型不匹配
	type JsonTest struct {
		UserName string `json:"UserName"`
	}
	testCase3 := JsonTest{}
	errs = json.Unmarshal([]byte("{\"UserName\":1}"), &testCase3)
	res = ErrorResponse(errs)
	asserts.Equal(serializer.ParamErr("JSON类型不匹配", errs), res)

	// 第四个测试案例，参数错误
	errs = errors.New("Unknown error")
	res = ErrorResponse(errs)
	asserts.Equal(serializer.ParamErr("参数错误", errs), res)

}

// 测试 ParamErrorMs
func TestParamErrorMsg(t *testing.T) {
	asserts := assert.New(t)
	testCase := []struct {
		field  string
		tag    string
		expect string
	}{
		{
			"UserName",
			"required",
			"用户名不能为空",
		},
		{
			"Password",
			"min",
			"密码太短",
		},
		{
			"Password",
			"max",
			"密码太长",
		},
		{
			"something unexpected",
			"max",
			"",
		},
		{
			"",
			"",
			"",
		},
	}
	for _, value := range testCase {
		asserts.Equal(value.expect, ParamErrorMsg(value.field, value.tag), "case %v", value)
	}
}
