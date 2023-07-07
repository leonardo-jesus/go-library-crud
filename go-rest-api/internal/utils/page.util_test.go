package utils

import (
	"reflect"
	"testing"
)

func TestNewPageParamUtil(t *testing.T) {
	tests := []struct {
		description string
		expected    PageParamUtilInterface
	}{
		{"when create new page param util, should return PageParamUtilInterface", NewPageParamUtil()},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := NewPageParamUtil(); !reflect.DeepEqual(got, test.expected) {
				t.Errorf("NewPageParamUtil() FAILED. Expected %v, got %v", test.expected, got)
			}
		})
	}
}

func Test_pageParamUtil_ConvertPageStringToInt(t *testing.T) {
	type params struct {
		page string
	}
	tests := []struct {
		description string
		params      params
		expected    int
	}{
		{"when page is '1', should return 1", params{page: "1"}, 1},
		{"when page is '0', should return 1", params{page: "0"}, 1},
		{"when page is '-1', should return 1", params{page: "-1"}, 1},
		{"when page is '', should return 1", params{page: ""}, 1},
		{"when there is no page param, should return 1", params{}, 1},
		{"when page is '2', should return 2", params{page: "2"}, 2},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			r := &pageParamUtil{}
			if got := r.ConvertPageStringToInt(test.params.page); got != test.expected {
				t.Errorf("pageParamUtil.GetPageFromQueryString() FAILED. Expected %v, got %v", test.expected, got)
			}
		})
	}
}
