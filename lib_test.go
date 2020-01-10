package haszero

import "testing"

func TestHasZero(t *testing.T) {
	type A struct {
		bbb string
		ccc string
	}
	a := A{ccc: "qaa"}
	if err := HasZero(a); err == nil {
		t.Error()
	}
	if err := HasZero(1); err == nil {
		t.Error(1)
	}

	type B struct {
		ddd string
	}

	if err := HasZero(B{ddd: "aaa"}); err != nil {
		t.Error(err)
	}

	type C struct {
		a A
		d string
	}
	c := C{
		a: A{
			bbb: "cc",
		},
		d: "aaa",
	}
	if err := HasZero(c); err == nil {
		t.Error(c)
	}

	cc := C{
		a: A{
			bbb: "cc",
			ccc: "ddd",
		},
		d: "aaa",
	}

	if err := HasZero(cc); err != nil {
		t.Error(err)
	}

	type E struct {
		c C
	}

	e := E{
		c: C{
			a: A{
				bbb: "aaa",
			},
			d: "sasas",
		},
	}
	if err := HasZero(e); err == nil {
		t.Error(e)
	}

	ee := E{
		c: C{
			a: A{
				bbb: "xxxx",
				ccc: "aaa",
			},
			d: "sdsd",
		},
	}
	if err := HasZero(ee); err != nil {
		t.Error(err)
	}
	if err := HasZero(&ee); err == nil {
		t.Error(err)
	}

	type H struct {
		ccc string
	}

	type G struct {
		aaa string
		h   H
	}

	type F struct {
		g *G
	}
	f := F{}
	if err := HasZero(f); err == nil {
		t.Error(f)
	}

	ff := F{
		g: &G{
			aaa: "aaa",
			h: H{
				ccc: "",
			},
		},
	}
	if err := HasZero(ff); err == nil {
		t.Error(ff)
	}

}
