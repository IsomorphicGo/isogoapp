package main

import "honnef.co/go/js/dom"

func main() {

	println("The sample Isomorphic Go skeleton web app successfully printed to your web browser console using GopherJS!")
	d := dom.GetWindow().Document()
	h := d.GetElementByID("welcomeMessage")
	h.SetInnerHTML("<p>The sample Isomorphic Go skeleton web app is working successfully.</p><p>GopherJS was used to print this welcome message.</p><p>There's another message waiting in your web browser console.</p>")
}
