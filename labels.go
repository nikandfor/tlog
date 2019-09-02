package tlog

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"
)

type (
	// Labels is a set of labels with optional values.
	//
	// By design Labels contains state diff not state itself.
	// So if you want to delete some label you should use Del method to add special thumbstone value.
	Labels []string
)

var (
	// AutoLabels is an list of automatically filled labels
	//     _hostname - local hostname
	//     _pid - process pid
	//     _md5 - this binary md5 hash
	AutoLabels = map[string]func() string{
		"_hostname": func() string {
			h, err := os.Hostname()
			if h == "" && err != nil {
				h = err.Error()
			}

			return h
		},
		"_pid": func() string {
			return fmt.Sprintf("%d", os.Getpid())
		},
		"_md5": func() string {
			f, err := os.Open(os.Args[0])
			if err != nil {
				return err.Error()
			}
			defer f.Close()

			h := md5.New()
			_, err = io.Copy(h, f)
			if err != nil {
				return err.Error()
			}

			return fmt.Sprintf("%02x", h.Sum(nil))
		},
	}
)

// ParseLabels parses comma separated list of labels and fills them with values (See FillLabelsWithDefaults).
func ParseLabels(s string) Labels {
	return FillLabelsWithDefaults(strings.Split(s, ",")...)
}

// FillLabelsWithDefaults creates Labels and fills autolabels (See AutoLabels).
func FillLabelsWithDefaults(labels ...string) Labels {
	ll := make(Labels, 0, len(labels))

	for _, lab := range labels {
		if f, ok := AutoLabels[lab]; ok {
			ll = append(ll, lab+"="+f())
		} else {
			ll = append(ll, lab)
		}
	}

	return ll
}

// Set sets k label value to v
func (ls *Labels) Set(k, v string) {
	val := k
	if v != "" {
		val += "=" + v
	}

	for i := 0; i < len(*ls); i++ {
		l := (*ls)[i]
		if l == "="+k {
			(*ls)[i] = val
			return
		} else if l == k || strings.HasPrefix(l, k+"=") {
			(*ls)[i] = val
			return
		}
	}
	*ls = append(*ls, val)
}

// Get gets k label value or "", false
func (ls *Labels) Get(k string) (string, bool) {
	for _, l := range *ls {
		if l == k {
			return "", true
		} else if strings.HasPrefix(l, k+"=") {
			return l[len(k)+1:], true
		}
	}
	return "", false
}

// Del replaces k label with special thumbstone.
// It's needed because Labels event contains state diff not state itself.
func (ls *Labels) Del(k string) {
	for i := 0; i < len(*ls); i++ {
		l := (*ls)[i]
		if l == "="+k {
			return
		} else if l == k || strings.HasPrefix(l, k+"=") {
			(*ls)[i] = "=" + k
		}
	}
}

// Merge merges two Labels sets
func (ls *Labels) Merge(b Labels) {
	for _, add := range b {
		if add == "" {
			continue
		}
		kv := strings.SplitN(add, "=", 2)
		if kv[0] == "" {
			ls.Del(kv[1])
		} else {
			ls.Set(kv[0], kv[1])
		}
	}
}

// Copy copies Labels including deleted thumbstones
func (ls *Labels) Copy() Labels {
	r := make(Labels, len(*ls))
	for i, v := range *ls {
		r[i] = v
	}
	return r
}