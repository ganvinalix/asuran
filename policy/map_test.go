package policy

import "testing"

func TestMapPolicy(t *testing.T) {
	cmd := "map http://url.g.cn/"
	p, err := Factory(cmd)
	if err != nil {
		t.Errorf(`Factory("%s") err: %v`, cmd, err)
	} else if p.Command() != cmd {
		t.Errorf(`Factory("%s").Command()=> "%s" != "%s"`, cmd, p.Command(), cmd)
	} else if _, ok := p.(*MapPolicy); !ok {
		t.Errorf(`Factory("%s") type wrong: %v`, cmd, p)
	} else {
		m := p.(*MapPolicy)
		url := "http://url.g.cn/"
		if m.URL("") != url {
			t.Errorf(`Factory("%s").URL("")=> "%s" != "%s"`, cmd, m.URL(""), url)
		}
	}

	cmd = "map replace /g.cn/google.cn/"
	p, err = Factory(cmd)
	if err != nil {
		t.Errorf(`Factory("%s") err: %v`, cmd, err)
	} else if p.Command() != cmd {
		t.Errorf(`Factory("%s").Command()=> "%s" != "%s"`, cmd, p.Command(), cmd)
	} else if _, ok := p.(*MapPolicy); !ok {
		t.Errorf(`Factory("%s") type wrong: %v`, cmd, p)
	} else {
		m := p.(*MapPolicy)
		source := "http://url.g.cn/"
		url := "http://url.google.cn/"
		if m.URL(source) != url {
			t.Errorf(`Factory("%s").URL("%s")=> "%s" != "%s"`, cmd, source, m.URL(source), url)
		}
	}

	// err test
	cmd = "map replace g.cn/google.cn/"
	p, err = Factory(cmd)
	if err == nil {
		t.Errorf(`Factory("%s") without err: %v`, cmd, p)
	}
}

func TestRedirectPolicy(t *testing.T) {
	cmd := "redirect http://url.g.cn/"
	p, err := Factory(cmd)
	if err != nil {
		t.Errorf(`Factory("%s") err: %v`, cmd, err)
	} else if p.Command() != cmd {
		t.Errorf(`Factory("%s").Command()=> "%s" != "%s"`, cmd, p.Command(), cmd)
	} else if _, ok := p.(*RedirectPolicy); !ok {
		t.Errorf(`Factory("%s") type wrong: %v`, cmd, p)
	} else {
		m := p.(*RedirectPolicy)
		url := "http://url.g.cn/"
		if m.URL("") != url {
			t.Errorf(`Factory("%s").URL("")=> "%s" != "%s"`, cmd, m.URL(""), url)
		}
	}

	cmd = "redirect replace /g.cn/google.cn/"
	p, err = Factory(cmd)
	if err != nil {
		t.Errorf(`Factory("%s") err: %v`, cmd, err)
	} else if p.Command() != cmd {
		t.Errorf(`Factory("%s").Command()=> "%s" != "%s"`, cmd, p.Command(), cmd)
	} else if _, ok := p.(*RedirectPolicy); !ok {
		t.Errorf(`Factory("%s") type wrong: %v`, cmd, p)
	} else {
		m := p.(*RedirectPolicy)
		source := "http://url.g.cn/"
		url := "http://url.google.cn/"
		if m.URL(source) != url {
			t.Errorf(`Factory("%s").URL("%s")=> "%s" != "%s"`, cmd, source, m.URL(source), url)
		}
	}

	// err test
	cmd = "redirect replace g.cn/google.cn/"
	p, err = Factory(cmd)
	if err == nil {
		t.Errorf(`Factory("%s") without err: %v`, cmd, p)
	}
}
