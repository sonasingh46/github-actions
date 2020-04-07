package main

import (
	"github.com/sonasingh46/github-actions/pkg/foo"
	"k8s.io/klog"
)
func main()  {
	klog.InitFlags(nil)
	number:=3
	if foo.IsEven(number){
		klog.Infof("%d is even",number)
	}else {
		klog.Infof("%d is odd",number)

	}
}