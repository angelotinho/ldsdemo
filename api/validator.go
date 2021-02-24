package api

import (
	"net/url"
)

func (n *Numbers) validate() url.Values {
	errs := url.Values{}
	if n.First == nil || n.Second == nil {
		errs.Add("Missing Numbers", "One of the numbers to the add function is missing. Add first, second")
	}
	return errs
}

func (n *Days) validate() url.Values {
	errs := url.Values{}
	if n.NumberOfDays == nil {
		errs.Add("DaysMissing", "Please Provide Number of Days")
	}
	return errs
}
