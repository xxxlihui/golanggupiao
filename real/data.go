package real

var pool = make(map[string][]string)

type Pool struct {
	//股票池的名称
	pools []string
	//股票池的内容
	pool map[string][]string
}

func (s *Pool) ForEach(iter func(name string, codes []string)) {
	for _, v := range s.pools {
		iter(v, s.pool[v])
	}
}

//添加代码到股票池
func (s *Pool) AddPoll(name string, codes []string) {
	if !s.ContainsPool(name) {
		s.pools = append(s.pools, name)
		s.pool[name] = codes
	} else {
		//合并
		for _, v := range codes {
			s.add(name, v)
		}
	}
}
func (s *Pool) Add(name string, code string) {
	if !s.ContainsPool(name) {
		s.pools = append(s.pools, name)
		s.pool[name] = []string{code}
	} else {
		//合并
		s.add(name, code)
	}
}
func (s *Pool) add(name string, code string) {
	ss := s.pool[name]
	for _, v := range ss {
		if v == code {
			return
		}
	}
	ss = append(ss, code)
	s.pool[name] = ss
}

/**
获取某一个股票池的票
*/
func (s *Pool) Get(name string) []string {
	if !s.ContainsPool(name) {
		return nil
	}
	return pool[name][:]
}

//判断是否包含有某一个池
func (s *Pool) ContainsPool(name string) bool {
	for _, v := range s.pools {
		if v == name {
			return true
		}
	}
	return false
}
