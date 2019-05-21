package epoch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []*Epoch{
	{1234567890, "Fri Feb 13 17:31:30 CST 2009"},
	{2651299932, "Tue Jan  6 02:12:12 CST 2054"},
	{3932702439, "Sun Aug 15 03:20:39 CST 2094"},
	{1641617187, "Fri Jan  7 22:46:27 CST 2022"},
	{2184599893, "Thu Mar 24 11:18:13 CST 2039"},
	{1450478535, "Fri Dec 18 16:42:15 CST 2015"},
	{1836762783, "Wed Mar 15 13:53:03 CST 2028"},
	{2923663878, "Thu Aug 24 10:51:18 CST 2062"},
	{4787313012, "Sun Sep 14 11:10:12 CST 2121"},
	{5685584294, "Tue Mar  3 02:58:14 CST 2150"},
	{3139626995, "Fri Jun 28 00:36:35 CST 2069"},
	{1761830335, "Thu Oct 30 07:18:55 CST 2025"},
	{2080146635, "Sat Dec  1 12:30:35 CST 2035"},
	{9762993053, "Sun May 18 08:30:53 CST 2279"},
	{7431001552, "Mon Jun 24 17:05:52 CST 2205"},
	{6836467482, "Mon Aug 21 12:44:42 CST 2186"},
	{2699454101, "Sat Jul 17 10:21:41 CST 2055"},
	{3327234509, "Sat Jun  8 09:48:29 CST 2075"},
	{9756901026, "Sat Mar  8 20:17:06 CST 2279"},
	{1293752076, "Thu Dec 30 17:34:36 CST 2010"},
}

func TestTimestampConvertion(t *testing.T) {
	for _, tc := range testCases {
		assert.Equal(t, tc, Convert(tc.Timestamp))
	}
}
